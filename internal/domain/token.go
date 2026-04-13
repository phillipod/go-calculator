package domain

import "fmt"

type TokenType int

const (
	TokenNumber TokenType = iota
	TokenOperator
	TokenFunction
	TokenLeftParen
	TokenRightParen
	TokenConstant
)

type Token struct {
	Type TokenType

	Value    string
	NumValue float64
}

func (t Token) String() string {
	return t.
		Value
}
func NewNumberToken(val float64) Token {
	return Token{Type: TokenNumber,
		Value:    fmt.Sprintf("%g", val),
		NumValue: val}
}
func NewOperatorToken(op string) Token {
	return Token{Type: TokenOperator,
		Value: op}
}
func NewFunctionToken(fn string) Token {
	return Token{Type: TokenFunction,
		Value: fn}
}
func NewLeftParenToken() Token {
	return Token{Type: TokenLeftParen, Value: "("}
}
func NewRightParenToken() Token {

	return Token{Type: TokenRightParen, Value: ")"}
}
func NewConstantToken(name string,

	val float64) Token {
	return Token{Type: TokenConstant,
		Value:    name,
		NumValue: val}
}
