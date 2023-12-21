package noduplicate

import (
	"errors"
	"fmt"
	"sync"
)

var errGoexit = errors.New("runtime.Goexit was called")

type panicError struct {
	value interface{}
}

func (p *panicError) Error() string {
	return fmt.Sprintf("%v\n", p.value)
}

func newPanicError(v interface{}) error {
	return &panicError{value: v}
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
			if !normalError {
				if r := recover(); r != nil {
					c.err = newPanicError(r)
				}
			}
		}()
		c.val, c.err = fn()
		normalError = true
	}()

	if !normalError {
		errorRecovered = true
	}
}
