package waitgrpsmpl

import (
	"fmt"
	"sync"
	"time"
)

func longRunTask(wg *sync.WaitGroup, name string, taskRunTime int) {
	defer wg.Done()
	fmt.Printf("running task %s \n", name)
	time.Sleep(time.Duration(taskRunTime * int(time.Second)))
}

func DoOperation() {
	wg := new(sync.WaitGroup)
	wg.Add(3)
	go longRunTask(wg, "first", 5)
	go longRunTask(wg, "second", 12)
	go longRunTask(wg, "third", 4)

	wg.Wait()
}
