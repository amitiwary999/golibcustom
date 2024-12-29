package algorithmgo

func merge(arL []int, arR []int) []int {
	var mergeArr []int
	var i = 0
	var j = 0
	for i < len(arL) && j < len(arR) {
		if arL[i] < arR[j] {
			mergeArr = append(mergeArr, arL[i])
			i++
		} else {
			mergeArr = append(mergeArr, arR[j])
			j++
		}
	}
	for i < len(arL) {
		mergeArr = append(mergeArr, arL[i])
		i++
	}
	for j < len(arR) {
		mergeArr = append(mergeArr, arR[j])
		j++
	}
	return mergeArr
}

func SplitMerge(datArr []int) []int {
	if len(datArr) == 1 {
		return datArr
	}

	var mid = len(datArr) / 2
	left := SplitMerge(datArr[0:mid])
	right := SplitMerge(datArr[mid:len(datArr)])
	mergeA := merge(left, right)
	return mergeA
}
