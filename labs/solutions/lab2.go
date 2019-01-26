package lab2

import (
	"fmt"
)

func bubbleSort(array []int) []int {
	arrayLength := len(array)

	for i := 0; i < arrayLength; i++ {
		for j := 0; j < arrayLength-i-1; j++ {

			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	return array
}

func main() {
	a := []int{10, 0, -3, 8, 5, -2}
	fmt.Println(bubbleSort(a))
}
