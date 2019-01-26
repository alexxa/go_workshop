package lab3

import (
	"testing"
)

func TestForSuccess(t *testing.T) {
	expectedMin := -5
	expectedMax := 200

	min, max, err := minMax([]int{44, 20, 14, 200, 9, -5})
	if err != nil {
		t.Error("minMax failed")
	}

	if expectedMin != min {
		t.Errorf("Expected min: '%d' but got '%d'", expectedMin, min)
	}
	if expectedMax != max {
		t.Errorf("Expected max: '%d' but got '%d'", expectedMax, max)
	}
}

func TestForFailure(t *testing.T) {
	_, _, err := minMax([]int{})
	if err == nil {
		t.Error("minMax isn't protected from working on empty lists")
	}
}
