package main

import (
	"fmt"
	"strconv"
)

func Calculate(expression string) (float64, error) {
	values := make(tokenStack, 0)
	operators := make(tokenStack, 0)
	tokens, err := tokenize(expression)
	if err != nil {
		return 0, err
	}
	for _, token := range tokens {
		switch token.Type {
		case TokenNumber:
			values = values.Push(token)
		case TokenOpenParenthesis:
			operators = operators.Push(token)
		case TokenCloseParenthesis:
			for !operators.IsEmpty() && operators.Top().Type != TokenOpenParenthesis {
				err := evaluateOperationOnStack(&operators, &values)
				if err != nil {
					return 0, err
				}
			}
			if operators.IsEmpty() {
				return 0, fmt.Errorf("parenthesis mismatch")
			}
			operators, _ = operators.Pop()
		case TokenOperator:
			for !operators.IsEmpty() &&
				getOperatorPrecedence(operators.Top()) >= getOperatorPrecedence(token) {
				err := evaluateOperationOnStack(&operators, &values)
				if err != nil {
					return 0, err
				}
			}
			operators = operators.Push(token)
		}
	}
	for !operators.IsEmpty() {
		if operators.Top().Type == TokenOpenParenthesis {
			return 0, fmt.Errorf("parenthesis mismatch")
		}
		err := evaluateOperationOnStack(&operators, &values)
		if err != nil {
			return 0, err
		}
	}
	if len(values) != 1 {
		return 0, fmt.Errorf("invalid expression")
	}
	result, _ := strconv.ParseFloat(values.Top().Value, 64)
	return result, nil
}

func evaluateOperationOnStack(operators *tokenStack, values *tokenStack) error {
	var operator, operand1, operand2 Token
	*operators, operator = operators.Pop()
	if len(*values) < 2 {
		return fmt.Errorf("invalid expression")
	}
	*values, operand2 = values.Pop()
	*values, operand1 = values.Pop()
	result := evaluateOperation(operand1.Value, operand2.Value, operator.Value)
	*values = values.Push(Token{result, TokenNumber})
	return nil
}

func evaluateOperation(operand1, operand2, operator string) string {
	number1, err1 := strconv.ParseFloat(operand1, 64)
	number2, err2 := strconv.ParseFloat(operand2, 64)
	if err1 != nil {
		panic("invalid numeric value " + operand1)
	}
	if err2 != nil {
		panic("invalid numeric value " + operand2)
	}
	var result float64
	switch operator {
	case "+":
		result = number1 + number2
	case "-":
		result = number1 - number2
	case "*":
		result = number1 * number2
	case "/":
		result = number1 / number2
	default:
		panic("invalid operator " + operator)
	}
	return fmt.Sprintf("%f", result)
}

func getOperatorPrecedence(token Token) int {
	switch token.Value {
	case "*", "/":
		return 2
	case "+", "-":
		return 1
	default:
		return 0
	}
}

type tokenStack []Token

func (s tokenStack) Push(value Token) tokenStack {
	return append(s, value)
}

func (s tokenStack) Pop() (tokenStack, Token) {
	l := len(s)
	return s[:l-1], s[l-1]
}

func (s tokenStack) Top() Token {
	return s[len(s)-1]
}

func (s tokenStack) IsEmpty() bool {
	return len(s) == 0
}
