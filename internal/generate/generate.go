package generate

import (
	"algorithm-study/internal/domain"
	"crypto/rand"
	"math/big"
)

func RandomObjectAvoidingID(max, avoidID int) *domain.Object {
	randomNumber, _ := rand.Int(rand.Reader, big.NewInt(int64(max)))
	randNumber := int(randomNumber.Int64())

	if randNumber == avoidID {
		return &domain.Object{
			ID:   avoidID + 1,
			Name: "wrong",
		}
	}

	return &domain.Object{
		ID:   randNumber,
		Name: "wrong",
	}
}

func RandomIntAvoidingNumber(max, avoidID int) int {
	randomNumber, _ := rand.Int(rand.Reader, big.NewInt(int64(max)))
	randNumber := int(randomNumber.Int64())

	if randNumber == avoidID {
		return avoidID + 1
	}

	return randNumber
}

func RandomInt(max int) int {
	number, _ := rand.Int(rand.Reader, big.NewInt(int64(max)))
	n := int(number.Int64())

	negativeFactor, _ := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if int(negativeFactor.Int64())%2 == 0 {
		return n * -1
	}

	return n
}

func RandomPositiveInt(max int) int {
	randomNumber, _ := rand.Int(rand.Reader, big.NewInt(int64(max)))
	randNumber := int(randomNumber.Int64())

	return randNumber
}

func UnsortedObjectArrayWithWorstCase(size int, maxValue int, rightID int) domain.Objects {
	var arr domain.Objects
	for i := 0; i < size-1; i++ {
		arr = append(arr, RandomObjectAvoidingID(maxValue, rightID))
	}

	arr = append(arr, &domain.Object{
		ID:   rightID,
		Name: "right",
	})

	return arr
}

func UnsortedIntArrayWithWorstCase(size int, maxValue int, number int) []int {
	var arr []int
	for i := 0; i < size-1; i++ {
		arr = append(arr, RandomIntAvoidingNumber(maxValue, number))
	}

	arr = append(arr, number)

	return arr
}

func UnsortedPositiveIntArray(size int, maxValue int) []int {
	var arr []int
	for i := 0; i < size; i++ {
		arr = append(arr, RandomPositiveInt(maxValue))
	}

	return arr
}

func UnsortedIntArray(size int, maxValue int) []int {
	var arr []int
	for i := 0; i < size; i++ {
		arr = append(arr, RandomInt(maxValue))
	}

	return arr
}

func RandomText() string {
	return "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum."
}
