package sort

func MergeInt(arr []int, orderDesc bool) []int {
	if len(arr) <= 1 {
		return arr
	}

	unsortedLeft := arr[:len(arr)/2]
	unsortedRight := arr[len(arr)/2:]

	sortedLeft := MergeInt(unsortedLeft, orderDesc)
	sortedRight := MergeInt(unsortedRight, orderDesc)

	condition := lessThan
	if orderDesc {
		condition = greaterThan
	}

	return mergeArrays(sortedLeft, sortedRight, condition)
}

func greaterThan(a int, b int) bool {
	return a > b
}

func lessThan(a int, b int) bool {
	return a < b
}

func mergeArrays(a []int, b []int, condition func(int, int) bool) []int {
	var final []int

	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if condition(a[i], b[j]) {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}

	for ; i < len(a); i++ {
		final = append(final, a[i])
	}

	for ; j < len(b); j++ {
		final = append(final, b[j])
	}

	return final
}
