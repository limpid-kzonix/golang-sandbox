package katas

func CountSheeps(numbers []bool) int {
	count := 0
	for i := 0; i < len(numbers); i++ {
		if numbers[i] {
			count++
		}
	}
	return count
}
