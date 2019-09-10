package main

import (
	"fmt"
	"testing"
)

func TestTokenize(t *testing.T) {
	s := tokenizeSuite{t}
	s.expect("", []Token{}, nil)
	s.expect("a", []Token{}, fmt.Errorf("invalid expression"))
	s.expect("12\ndsa22", []Token{}, fmt.Errorf("invalid expression"))
	s.expect("e12\ndsa22", []Token{}, fmt.Errorf("invalid expression"))
	s.expect("0", []Token{Token{"0", TokenNumber}}, nil)
	s.expect("9290", []Token{Token{"9290", TokenNumber}}, nil)
	s.expect("  123 ", []Token{Token{"123", TokenNumber}}, nil)
	s.expect(
		"80-90*6 + (1 - 0) * 8",
		[]Token{
			Token{"80", TokenNumber},
			Token{"-", TokenOperator},
			Token{"90", TokenNumber},
			Token{"*", TokenOperator},
			Token{"6", TokenNumber},
			Token{"+", TokenOperator},
			Token{"(", TokenOpenParenthesis},
			Token{"1", TokenNumber},
			Token{"-", TokenOperator},
			Token{"0", TokenNumber},
			Token{")", TokenCloseParenthesis},
			Token{"*", TokenOperator},
			Token{"8", TokenNumber},
		},
		nil,
	)
}

type tokenizeSuite struct {
	t *testing.T
}

func (s tokenizeSuite) expect(input string, expectedResult []Token, expectedError error) {
	result, err := tokenize(input)
	testError := ""
	if !s.equal(result, expectedResult) {
		testError += fmt.Sprintf("\nExpected: %v\nActual:   %v\n", expectedResult, result)
	}
	testError += expectError(err, expectedError)
	if testError != "" {
		s.t.Errorf(testError)
	}
}

func expectError(err error, expectedError error) string {
	if err == nil {
		if expectedError != nil {
			return fmt.Sprintf("\nExpected error: %v\nActual error:   nil\n", expectedError)
		}
	} else {
		if expectedError == nil {
			return fmt.Sprintf("\nExpected error: nil\nActual error:   %v\n", err)
		} else if err.Error() != expectedError.Error() {
			return fmt.Sprintf("\nExpected error: %v\nActual error:   %v\n", expectedError, err)
		}
	}
	return ""
}

func (s tokenizeSuite) equal(first []Token, second []Token) bool {
	if len(first) != len(second) {
		return false
	}
	for i := range first {
		if first[i] != second[i] {
			return false
		}
	}
	return true
}
