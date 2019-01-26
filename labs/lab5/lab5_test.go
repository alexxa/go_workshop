package lab5

import (
	"testing"
)

func TestForRecursion(t *testing.T) {
	if fibResult := fibonacciRecursion(10); fibResult != 55 {
		t.Errorf("fibonacciRecursion(10) returned '%d', but expected '55'", fibResult)
	}
}

func TestForLoop(t *testing.T) {
	if fibResult := fibonacciLoop(10); fibResult != 55 {
		t.Errorf("fibonacciLoop(10) returned '%d', but expected '55'", fibResult)
	}
}
