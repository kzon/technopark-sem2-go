package main

import (
	"fmt"
	"regexp"
	"strings"
)

type Token struct {
	Value string
	Type  int
}

const (
	TokenNumber = iota
	TokenOperator
	TokenOpenParenthesis
	TokenCloseParenthesis
)

func tokenize(expression string) ([]Token, error) {
	expression = strings.ReplaceAll(expression, " ", "")
	var tokens []Token
	token := getFirstToken(expression)
	for token != nil {
		tokens = append(tokens, *token)
		expression = expression[len(token.Value):]
		token = getFirstToken(expression)
	}
	if expression != "" {
		return nil, fmt.Errorf("invalid expression")
	}
	return tokens, nil
}

var numberRegexp = regexp.MustCompile(`^(\d+)`)

func getFirstToken(expression string) *Token {
	if len(expression) == 0 {
		return nil
	}
	if value := numberRegexp.FindString(expression); value != "" {
		return &Token{value, TokenNumber}
	}
	firstChar := expression[0]
	switch firstChar {
	case '+', '-', '*', '/':
		return &Token{string(firstChar), TokenOperator}
	case '(':
		return &Token{string(firstChar), TokenOpenParenthesis}
	case ')':
		return &Token{string(firstChar), TokenCloseParenthesis}
	}
	return nil
}
