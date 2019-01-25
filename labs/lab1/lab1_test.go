package lab1

import (
	"reflect"
	"testing"
)

func TestForSuccess(t *testing.T) {
	expected := map[rune]int{
		'o': 5,
		'p': 2,
		'g': 2,
		'n': 2,
		'r': 3,
		'a': 2,
		' ': 5,
		'm': 2,
	}

	test1 := getDuplicateCharacters("Go programming workshop at devconf 2019")

	eq := reflect.DeepEqual(expected, test1)
	if !eq {
		t.Error("assertion error")
	}
}

func TestForFail(t *testing.T) {
	expected := map[rune]int{
		'o': 5,
		'p': 2,
		'g': 2,
		'n': 2,
		'r': 3,
		'a': 2,
		' ': 5,
		'm': 2,
	}

	test1 := getDuplicateCharacters("go programming workshop at devconf 2019")

	eq := reflect.DeepEqual(expected, test1)
	if eq {
		t.Error("assertion error")
	}
}
