package algorithmgo

func maxHeap(hArr []int, ind int, size int) {
	left := 2 * ind
	right := 2*ind + 1
	larInd := ind
	if left >= 0 && left < size && hArr[larInd] < hArr[left] {
		larInd = left
	}
	if right >= 0 && right < size && hArr[larInd] < hArr[right] {
		larInd = right
	}
	if larInd != ind {
		hArr[ind], hArr[larInd] = hArr[larInd], hArr[ind]
		maxHeap(hArr, larInd, size)
	}
}

func heapCreation(inArr []int) {
	var i int = (len(inArr) - 1) / 2
	for ; i >= 0; i-- {
		maxHeap(inArr, i, len(inArr))
	}
}

func HeapSort(inArr []int) []int {
	heapCreation(inArr)
	size := len(inArr)
	for i := len(inArr) - 1; i >= 0; i-- {
		inArr[0], inArr[i] = inArr[i], inArr[0]
		size -= 1
		maxHeap(inArr, 0, size)
	}
	return inArr
}
