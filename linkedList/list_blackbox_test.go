package linkedList_test

import (
	. "ds/linkedList"
	"testing"
)

func TestReverse_WithEmptyList(t *testing.T) {
	var list IList
	list = New()
	list.Reverse()
	if list.Len() != 0 {
		t.Errorf("Size not 0 after reversal!")
	}
}

func TestReverse_WithSingleSizedList(t *testing.T) {
	var list IList
	list = New()
	list.Append(1)

	list.Reverse()

	expected := []int{1}
	result := []int{*list.PopFront()}

	if expected[0] != result[0] {
		t.Errorf(
			"Expected = %q, Actual = %q",
			expected[0], result[0])
	}
}

func TestReverse_WithOddSizedList(t *testing.T) {
	original := []int{1, 2, 3, 4, 5}

	var list IList
	list = New()
	for _, number := range original {
		list.Append(number)
	}

	list.Reverse()

	expected := []int{5, 4, 3, 2, 1}
	result := make([]int, 0)
	for list.Len() > 0 {
		result = append(result, *list.PopFront())
	}

	for index := 0; index < len(expected); index++ {
		expectedNum := expected[index]
		actualNum := result[index]

		if expectedNum != actualNum {
			t.Errorf(
				"Expected = %q, Actual = %q at index = %q",
				expectedNum, actualNum, index)
		}
	}
}

func TestReverse_WithEvenSizedList(t *testing.T) {
	original := []int{1, 2, 3, 4, 5, 6, 7, 8}

	var list IList
	list = New()
	for _, number := range original {
		list.Append(number)
	}

	list.Reverse()

	expected := []int{1, 2, 3, 4, 5, 6, 7, 8}
	result := make([]int, 0)
	for list.Len() > 0 {
		result = append(result, *list.PopBack())
	}

	for index := 0; index < len(expected); index++ {
		expectedNum := expected[index]
		actualNum := result[index]

		if expectedNum != actualNum {
			t.Errorf(
				"Expected = %q, Actual = %q at index = %q",
				expectedNum, actualNum, index)
		}
	}
}
