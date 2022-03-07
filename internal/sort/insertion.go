package sort

func InsertionInt(arr []int, orderDesc bool) []int {
	length := len(arr)

	condition := lessThan
	if orderDesc {
		condition = greaterThan
	}

	for i := 0; i < length; i++ {
		for j := i; j > 0; j-- {
			if condition(arr[j-1], arr[j]) {
				arr[j-1], arr[j] = arr[j], arr[j-1]
			}
		}
	}

	return arr
}
