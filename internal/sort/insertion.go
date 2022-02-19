package sort

func InsertionInt(arr []int, orderDesc bool) []int {
	final := arr
	length := len(final)

	condition := lessThan
	if orderDesc {
		condition = greaterThan
	}

	for i := 0; i < length; i++ {
		for j := i; j > 0; j-- {
			if condition(final[j-1], final[j]) {
				final[j], final[j-1] = final[j-1], final[j]
			}
		}
	}

	return final
}
