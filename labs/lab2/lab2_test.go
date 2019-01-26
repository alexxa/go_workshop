package lab2

import (
	"reflect"
	"testing"
)

func TestForSuccess(t *testing.T) {
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sorted := bubbleSort([]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1})
	if eq := reflect.DeepEqual(expected, sorted); !eq {
		t.Errorf("bubblesort received: %v, but expected %v", sorted, expected)
	}
}
