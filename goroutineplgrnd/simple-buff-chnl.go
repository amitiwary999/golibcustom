package gortntry

import (
	"fmt"
	"time"
)

type bufChn struct {
	nameChn chan string
}

func (b *bufChn) done() {
	<-b.nameChn
}

func (b *bufChn) EntryFunc(str string) {
	b.nameChn <- str
	fmt.Printf("before the goroutine %s\n", str)
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
