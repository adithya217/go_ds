package nextLargerElement_test

import (
	"ds/stack/nextLargerElement"
	"testing"
)

func testInput(input []uint, expected []int, t *testing.T) {
	actual := nextLargerElement.FindNextLargerElements(input)

	if len(actual) != len(expected) {
		t.Errorf("Expected = %v, Actual = %v", expected, actual)
		return
	}

	for index := 0; index < len(actual); index++ {
		if expected[index] != actual[index] {
			t.Errorf("Expected = %v, Actual = %v", expected[index], actual[index])
			return
		}
	}
}

func TestNextLargerElement_WithEmptyInput(t *testing.T) {
	input := []uint{}
	testInput(input, []int{}, t)
}

func TestNextLargerElement_WithAscendingOrder(t *testing.T) {
	input := []uint{1, 2, 3, 4}
	testInput(input, []int{2, 3, 4, -1}, t)
}

func TestNextLargerElement_WithDescendingOrder(t *testing.T) {
	input := []uint{4, 3, 2, 1}
	testInput(input, []int{-1, -1, -1, -1}, t)
}

func TestNextLargerElement_WithAscendingMixedOrder(t *testing.T) {
	input := []uint{1, 3, 2, 4}
	testInput(input, []int{3, 4, 4, -1}, t)
}

func TestNextLargerElement_WithDescendingMixedOrder(t *testing.T) {
	input := []uint{4, 3, 5, 1}
	testInput(input, []int{5, 5, -1, -1}, t)
}
