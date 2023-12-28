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

/** on RWMutex when we take the Lock then can't change the common variable until it is unlock. If we take RLock then until all the RLock is RUnlock
we can't take the Lock again to make changes */

/** for ref output
enter to set value second
set value before sleep second
enter to set value first
enter to set value third
going to unlock with value second
set value before sleep first
going to unlock with value first
print the value first
print the value first
print the value first
print the value before read unlock first
print the value before read unlock first
print the value before read unlock first
set value before sleep third
going to unlock with value third

It means that even the Lock that set value 'first' is Unlock the Lock with value 'third' wait for all RUnlock. */

func normalLU(v string) {
	fmt.Printf("enter to set value %v\n", v)
	defer wg.Done()
	rmmt.Lock()
	myv = v
	fmt.Printf("set value before sleep %v\n", myv)
	time.Sleep(3 * time.Second)
	fmt.Printf("going to unlock with value %v\n", myv)
	rmmt.Unlock()
}

func normalRWL() {
	defer wg.Done()
	time.Sleep(5 * time.Second)
	rmmt.RLock()
	fmt.Printf("print the value %v \n", myv)
	time.Sleep(6 * time.Second)
	fmt.Printf("print the value before read unlock %v \n", myv)
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
