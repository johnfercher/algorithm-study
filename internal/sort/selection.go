package sort

func SelectionInt(arr []int, orderDesc bool) []int {
	final := arr
	length := len(final)

	condition := lessThan
	if orderDesc {
		condition = greaterThan
	}

	for i := 0; i < length; i++ {
		minValue := final[i]
		minIndex := i
		for j := i + 1; j < length; j++ {
			if condition(minValue, final[j]) {
				minValue = final[j]
				minIndex = j
			}
		}

		final[i], final[minIndex] = final[minIndex], final[i]
	}

	return final
}
