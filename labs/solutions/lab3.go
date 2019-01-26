package lab3

import "fmt"

func minMax(input []int) (int, int, error) {
	// if input is invalid return error
	if input == nil || len(input) == 0 {
		return -1, -1, fmt.Errorf("Empty slice")
	}
	min := input[0]
	max := input[0]

	for _, val := range input {
		if val < min {
			min = val
		}
		if val > max {
			max = val
		}
	}
	return min, max, nil
}
