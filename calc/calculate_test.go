package main

import (
	"fmt"
	"testing"
)

func TestCalculateSimpleExpressions(t *testing.T) {
	s := calculateSuite{t}
	s.expect("", 0, fmt.Errorf("invalid expression"))
	s.expect("0", 0, nil)
	s.expect("54", 54, nil)
}

func TestCalculateSignleOperation(t *testing.T) {
	s := calculateSuite{t}
	s.expect("1+3", 4, nil)
	s.expect("100-98", 2, nil)
	s.expect("98-100", -2, nil)
	s.expect("45*2", 90, nil)
	s.expect("35/7", 5, nil)

	s.expect("*", 0, fmt.Errorf("invalid expression"))
	s.expect("/71", 0, fmt.Errorf("invalid expression"))
	s.expect("512+", 0, fmt.Errorf("invalid expression"))
}

func TestCalculateMultipleOperations(t *testing.T) {
	s := calculateSuite{t}
	s.expect("4*20-5", 75, nil)
	s.expect("1+2+3+4+5+6+7*11", 98, nil)
	s.expect("55/11*67-340", -5, nil)
}

func TestCalculateWithParenthesis(t *testing.T) {
	s := calculateSuite{t}
	s.expect("(1+2)*9", 27, nil)
	s.expect("(1-5)*2+12", 4, nil)
	s.expect("(1-5)*(2+12)", -56, nil)

	s.expect("34-(124", 0, fmt.Errorf("parenthesis mismatch"))
	s.expect("34-124)-2", 0, fmt.Errorf("parenthesis mismatch"))
	s.expect("(1-5)*2+12)", 0, fmt.Errorf("parenthesis mismatch"))
}

type calculateSuite struct {
	t *testing.T
}

func (s calculateSuite) expect(expression string, expectedResult float64, expectedError error) {
	result, err := Calculate(expression)
	testError := ""
	if result != expectedResult {
		testError += fmt.Sprintf("Expected: %v\nActual:   %v\n", expectedResult, result)
	}
	testError += expectError(err, expectedError)
	if testError != "" {
		s.t.Errorf("\nCalculating %v\n%v", expression, testError)
	}
}
