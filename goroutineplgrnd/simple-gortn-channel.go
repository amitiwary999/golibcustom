package gortntry

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

func cpuIntensiveTask(ch *chan string, name string, waitTime int) {
	time.Sleep(time.Duration(waitTime * int(time.Second)))
	fmt.Printf("task completed %s\n", name)
	*ch <- name
}

func waitFirstTaskComplete(ch *chan string, cancel context.CancelFunc) {
	for {
		value, _ := <-*ch
		if value == "firstTask" {
			cancel()
			return
		}
	}
}

func DoOperationSimple() {
	_, cancel := context.WithCancel(context.Background())
	chnl := make(chan string)
	nameArray := [6]string{"firstTask", "secondTask", "thirdTask", "fourthTask", "fifthTask", "sixthTask"}
	waitTimeArray := [6]int{13, 18, 1, 24, 36, 16}

	for i, name := range nameArray {
		go cpuIntensiveTask(&chnl, name, waitTimeArray[i])
	}
	waitFirstTaskComplete(&chnl, cancel)
	fmt.Println("All goroutine done and program stop")
}

func waitFirstTaskCompleteErrGroup(ch *chan string, cntx context.Context) {
	for {
		value, _ := <-*ch
		if value == "firstTask" {
			cntx.Done()
			return
		}
	}
}

func DoOperationSimpleErrGrp() {
	ctx := context.Background()
	eg, egCtx := errgroup.WithContext(ctx)
	chnl := make(chan string)
	nameArray := [6]string{"firstTask", "secondTask", "thirdTask", "fourthTask", "fifthTask", "sixthTask"}
	waitTimeArray := [6]int{13, 18, 1, 24, 36, 16}

	for i, name := range nameArray {
		index := i
		taskName := name
		eg.Go(func() error {
			cpuIntensiveTask(&chnl, taskName, waitTimeArray[index])
			return nil
		})
	}
	waitFirstTaskCompleteErrGroup(&chnl, egCtx)

}
