package sort

func MergeInt(arr []int, orderDesc bool) []int {
	if len(arr) < 2 {
		return arr
	}

	left := arr[:len(arr)/2]
	right := arr[len(arr)/2:]

	sortedLeft := MergeInt(left, orderDesc)
	sortedRight := MergeInt(right, orderDesc)

	return mergeArrays(sortedLeft, sortedRight, orderDesc)
}

func mergeArrays(a []int, b []int, orderDesc bool) []int {
	var merged []int

	condition := lessThan
	if orderDesc {
		condition = greaterThan
	}

	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if condition(a[i], b[j]) {
			merged = append(merged, a[i])
			i++
		} else {
			merged = append(merged, b[j])
			j++
		}
	}

	for ; i < len(a); i++ {
		merged = append(merged, a[i])
	}

	for ; j < len(b); j++ {
		merged = append(merged, b[j])
	}

	return merged
}
