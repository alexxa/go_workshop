package lab1

func getDuplicateCharacters(input string) map[rune]int {

	frequency := make(map[rune]int)
	for _, char := range input {
		if _, ok := frequency[char]; ok {
			frequency[char]++
		} else {
			frequency[char] = 1
		}
	}
	for key, value := range frequency {
		if value == 1 {
			delete(frequency, key)
		}
	}
	return frequency
}
