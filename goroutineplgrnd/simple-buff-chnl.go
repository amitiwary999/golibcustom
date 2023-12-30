package gortntry

import (
	"fmt"
	"sync"
	"time"
)

type bufChn struct {
	nameChn chan string
	wg      sync.WaitGroup
}

func (b *bufChn) Wait() {
	b.wg.Wait()
}

func (b *bufChn) done() {
	<-b.nameChn
	b.wg.Done()
}

func (b *bufChn) EntryFunc(str string) {
	b.nameChn <- str
	fmt.Printf("before the goroutine %s\n", str)
	b.wg.Add(1)
	go func() {
		defer b.done()
		time.Sleep(time.Duration(5 * time.Second))
		fmt.Printf("print on go %s\n", str)
	}()
}

func NewBufChn() *bufChn {
	return &bufChn{
		nameChn: make(chan string, 4),
	}
}
