package dsl

type Token struct {
	value    string
	position int
	_type    int
}

const TokenWhiteSpace = 0
const TokenLeftParen = 1
const TokenRightParen = 2
const TokenComma = 3
const TokenDot = 4
const TokenLiteralValue = 5
const TokenID = 6

func (token Token) GetValue() string {
	return token.value
}

func (token Token) GetPosition() int {
	return token.position
}

func (token Token) GetType() int {
	return token._type
}
