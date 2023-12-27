package gortntry

import (
	"fmt"
	"sync"
	"time"
)

var mt sync.Mutex
var rmmt sync.RWMutex
var wg sync.WaitGroup
var myv string

func normalLU(v string) {
	fmt.Printf("enter to set value %v\n", v)
	defer wg.Done()
	mt.Lock()
	myv = v
	fmt.Printf("set value before sleep %v\n", myv)
	time.Sleep(3 * time.Second)
	fmt.Printf("going to unlock with value %v\n", myv)
	mt.Unlock()
}

func normalRWL() {
	defer wg.Done()
	time.Sleep(4 * time.Second)
	rmmt.RLock()
	fmt.Printf("print the value %v \n", myv)
	time.Sleep(1 * time.Second)
	rmmt.RUnlock()
}

func DoCall() {
	wg.Add(1)
	go normalLU("first")
	wg.Add(1)
	go normalLU("third")
	wg.Add(1)
	go normalRWL()
	wg.Add(1)
	go normalRWL()
	wg.Add(1)
	go normalRWL()
	wg.Add(1)
	go normalLU("second")
	wg.Wait()
}
