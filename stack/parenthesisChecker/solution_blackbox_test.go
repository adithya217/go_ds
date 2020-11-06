package parenthesisChecker_test

import (
	"ds/stack/parenthesisChecker"
	"testing"
)

func testInput(input string, results []bool, t *testing.T) {
	flags := []bool{false, true}

	for index := 0; index < len(results); index++ {
		flag := flags[index]
		actual := parenthesisChecker.IsBalanced(input, flag)
		expected := results[index]

		if actual != expected {
			t.Errorf("Expected = %v, Actual = %v", expected, actual)
			return
		}
	}
}

func TestParenthesisChecker_WithEmptyString(t *testing.T) {
	input := ""
	testInput(input, []bool{true, true}, t)

}

func TestParenthesisChecker_WithBalancedString(t *testing.T) {
	input := "[()]{}{[()()]()}"
	testInput(input, []bool{true, true}, t)
}

func TestParenthesisChecker_WithUnbalancedString(t *testing.T) {
	input := "[(])"
	testInput(input, []bool{false, false}, t)
}

func TestParenthesisChecker_WithInvalidBalancedString(t *testing.T) {
	input := "]{()}["
	testInput(input, []bool{false, true}, t)
}
