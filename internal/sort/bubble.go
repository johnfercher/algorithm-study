package sort

func BubbleInt(arr []int, orderDesc bool) []int {
	final := arr
	length := len(final)

	condition := lessThan
	if orderDesc {
		condition = greaterThan
	}

	for i := 0; i < length; i++ {
		for j := 0; j < length-1; j++ {
			if condition(arr[j], arr[j+1]) {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	return final
}
