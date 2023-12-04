package gortntry

import (
	"context"
	"fmt"
	"time"
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
