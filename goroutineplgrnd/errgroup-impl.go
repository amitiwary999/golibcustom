package gortntry

import (
	"context"
	"fmt"
	"sync"
)

type semToken struct{}

type ErGroup struct {
	sem     chan semToken
	wg      sync.WaitGroup
	errOnce sync.Once
	cancel  func(error)
	err     error
}

func NewErGroup(parent context.Context) (*ErGroup, context.Context) {
	ctx, cancelFunc := context.WithCancelCause(parent)
	return &ErGroup{cancel: cancelFunc}, ctx
}

func (eg *ErGroup) done() {
	if eg.sem != nil {
		<-eg.sem
	}
	eg.wg.Done()
}

/**return the first error that we save in err of ErGroup struct once all goroutines is done*/
func (eg *ErGroup) Wait() error {
	eg.wg.Wait()
	return eg.err
}

func (eg *ErGroup) Go(f func() error) {
	if eg.sem != nil {
		/** empty struct helps to save memory. if we use string then there is some exchange of data between the goroutines*/
		eg.sem <- semToken{}
	}
	eg.wg.Add(1)
	go func() {
		defer eg.done()
		if err := f(); err != nil {
			/**register the first error that we found*/
			eg.errOnce.Do(func() {
				eg.err = err
				if eg.cancel != nil {
					eg.cancel(err)
				}
			})
		}
	}()
}

/*
  - this function use when we don't want to wait if number of goroutines active is already equal to limit
    but return false that inform that we already have n number of active goroutine and can't process new
*/
func (eg *ErGroup) NoWaitGo(fn func() error) bool {
	if eg.sem != nil {
		select {
		case eg.sem <- semToken{}:
		default:
			return false
		}
	}
	eg.wg.Add(1)
	go func() {
		defer eg.done()
		if err := fn(); err != nil {
			/**register the first error that we found*/
			eg.errOnce.Do(func() {
				eg.err = err
				if eg.cancel != nil {
					eg.cancel(err)
				}
			})
		}
	}()
	return true
}

func (eg *ErGroup) SetLimit(n int) {
	if n < 0 {
		eg.sem = nil
		return
	}
	if len(eg.sem) > 0 {
		fmt.Printf("can't update the limit while %v goroutines are still active", len(eg.sem))
	}
	eg.sem = make(chan semToken, n)
}
