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

/*
for loop keep run and read the value from channel, if it is firstTask then return
So this method make sure that it return and calling function move to next step like DoOperationSimple function.
If it doesn't return, which is possible if channel doesn't have value firstTask, then the calling function stuck and goroutine
remain asleep
*/
func waitFirstTaskComplete(ch *chan string) {
	for {
		value, _ := <-*ch
		if value == "firstTask" {
			return
		}
	}
}

func DoOperationSimple() {
	chnl := make(chan string)
	nameArray := [6]string{"firstTask", "secondTask", "thirdTask", "fourthTask", "fifthTask", "sixthTask"}
	waitTimeArray := [6]int{13, 18, 1, 24, 36, 16}
	for i, name := range nameArray {
		go cpuIntensiveTask(&chnl, name, waitTimeArray[i])
	}
	/* it keep wait here and once function return then move to next step. if value is firstTask then it return*/
	waitFirstTaskComplete(&chnl)
	fmt.Println("All goroutine done and program stop")
}

func waitFirstTaskCompleteErrGroup(ch *chan string, cntx context.Context) {
	for {
		value, _ := <-*ch
		if value == "firstTask" {
			cntx.Done()
			// return
		}
	}
}

/** I need to work on this, once I understand the error group usage */
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
