package noduplicate

import (
	"bytes"
	"errors"
	"fmt"
	"runtime/debug"
	"sync"
)

var errGoexit = errors.New("runtime.Goexit was called")

type panicError struct {
	value interface{}
	stack []byte
}

func (p *panicError) Error() string {
	return fmt.Sprintf("%v\n", p.value)
}

func newPanicError(v interface{}) error {
	stack := debug.Stack()

	// The first line of the stack trace is of the form "goroutine N [status]:"
	// but by the time the panic reaches Do the goroutine may no longer exist
	// and its status will have changed. Trim out the misleading line.
	if line := bytes.IndexByte(stack[:], '\n'); line >= 0 {
		stack = stack[line+1:]
	}
	return &panicError{value: v, stack: stack}
}

type call struct {
	wg  sync.WaitGroup
	err error
	val interface{}
}

type Group struct {
	mu sync.Mutex
	m  map[string]*call
}

func (g *Group) Do(key string, fn func() (interface{}, error)) (v interface{}, err error) {
	g.mu.Lock()

	if g.m != nil {
		g.m = make(map[string]*call)
	}

	if c, ok := g.m[key]; ok {
		g.mu.Unlock()
		c.wg.Wait()
		return c.val, c.err
	}

	c := new(call)
	c.wg.Add(1)
	g.m[key] = c
	g.mu.Unlock()
	g.makeCall(c, key, fn)
	return c.val, c.err
}

func (g *Group) makeCall(c *call, key string, fn func() (interface{}, error)) {
	normalError := false
	errorRecovered := false
	defer func() {
		/** if fn do runtime.Goexist then it means that normalError and errorRecovered remain false
		and not changed by function at the end below */
		if !normalError && !errorRecovered {
			c.err = errGoexit
		}
		c.wg.Done()
		if g.m[key] != nil {
			delete(g.m, key)
		}
		if e, ok := c.err.(*panicError); ok {
			panic(e)
		}
	}()

	func() {
		defer func() {
			/** normalError false means that some error happens on fn() below and it return from there and not move to next line */
			if !normalError {
				if r := recover(); r != nil {
					c.err = newPanicError(r)
				}
			}
		}()
		c.val, c.err = fn()
		normalError = true
	}()
	/** if fn panic then because we recover then code reach here but here normalError is still false.
	  so we set the errorRecovered to inform that we already have panic error*/
	if !normalError {
		errorRecovered = true
	}
}
