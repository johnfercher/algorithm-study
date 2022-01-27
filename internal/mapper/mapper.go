package mapper

import "algorithm-study/internal/domain"

func ArrayToMap(arr []int, counter *domain.Counter) map[int]bool {
	m := make(map[int]bool)

	for _, value := range arr {
		counter.Increment()
		m[value] = true
	}

	return m
}
