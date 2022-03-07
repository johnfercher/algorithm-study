package sort

func SelectionInt(arr []int, orderDesc bool) []int {
	length := len(arr)

	condition := lessThan
	if orderDesc {
		condition = greaterThan
	}

	for i := 0; i < length; i++ {
		conditionValue := arr[i]
		conditionIndex := i
		for j := i + 1; j < length; j++ {
			if condition(conditionValue, arr[j]) {
				conditionValue = arr[j]
				conditionIndex = j
			}
		}

		arr[i], arr[conditionIndex] = arr[conditionIndex], arr[i]
	}

	return arr
}
