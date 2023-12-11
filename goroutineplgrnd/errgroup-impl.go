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

func (eg *ErGroup) wait() error {
	eg.wg.Wait()
	return eg.err
}

func (eg *ErGroup) Go(f func() error) {
	if eg.sem != nil {
		eg.sem <- semToken{}
	}
	eg.wg.Add(1)
	go func() {
		defer eg.done()
		if err := f(); err != nil {
			eg.errOnce.Do(func() {
				eg.err = err
				if eg.cancel != nil {
					eg.cancel(err)
				}
			})
		}
	}()
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
