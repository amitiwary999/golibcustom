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

func cpuIntensiveTaskWithError(name string, waitTime int) error {
	time.Sleep(time.Duration(waitTime * int(time.Second)))
	fmt.Printf("task completed %s\n", name)
	if name == "firstTask" {
		return fmt.Errorf("return error for task %s\n", name)
	}
	return nil
}

func DoOperationSimpleErrGrp() {
	ctx := context.Background()
	eg, _ := errgroup.WithContext(ctx)
	nameArray := [6]string{"firstTask", "secondTask", "thirdTask", "fourthTask", "fifthTask", "sixthTask"}
	waitTimeArray := [6]int{13, 18, 1, 24, 36, 16}

	for i, name := range nameArray {
		index := i
		taskName := name
		eg.Go(func() error {
			return cpuIntensiveTaskWithError(taskName, waitTimeArray[index])
		})
	}
	if err := eg.Wait(); err != nil {
		fmt.Printf("An error occurred: %v\n", err)
	}
}
