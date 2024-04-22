package algorithmgo

import (
	"fmt"
	"sort"
)

/** find ith smallest number in an unsorted array*/

func MedianAlgo(ar []int, i int) {
	ithS := findIthSmall(ar, i)
	fmt.Printf("ith small number %v\n", ithS)
}

func findIthSmall(ar []int, i int) int {
	var divAr [][]int
	var flag int = 0
	lenar := len(ar) - len(ar)%5
	for flag < lenar {
		divAr = append(divAr, ar[flag:flag+5])
		flag = flag + 5
	}
	divAr = append(divAr, ar[flag:])
	var median []int
	for _, arr := range divAr {
		sort.Slice(arr, func(i2, j int) bool {
			return i2 > j
		})
		ind := len(arr) / 2
		median = append(median, arr[ind])
	}

	sort.Slice(median, func(i2, j int) bool {
		return i2 > j
	})

	medianInd := len(median) / 2
	var medianAccAr []int
	divi := median[medianInd]
	diviInd := 0
	for _, item := range ar {
		if item > divi {
			medianAccAr = append(medianAccAr, item)
		} else {
			diviInd++
			medianAccAr = append([]int{item}, medianAccAr...)
		}
	}
	if diviInd == i {
		return divi
	}
	if diviInd > i {
		return findIthSmall(medianAccAr[:diviInd], i)
	} else {
		return findIthSmall(medianAccAr[diviInd:], i-diviInd)
	}
}
