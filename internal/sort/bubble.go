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
			if condition(final[j], final[j+1]) {
				final[j], final[j+1] = final[j+1], final[j]
			}
		}
	}

	return final
}
