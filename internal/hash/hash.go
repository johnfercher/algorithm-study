package hash

func Generate(value string, maxValue int) int {
	characters := []rune(value)

	hash := 0
	for i := 0; i < len(characters); i++ {
		asci := int(characters[i])
		hash = (hash + asci*i) % maxValue
	}

	return hash
}
