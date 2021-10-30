package dsl

import "fmt"

type EndOfFile struct{}

func (m *EndOfFile) Error() string {
	return ""
}

type TokenError struct {
	expected TokenType
	token    Token
}

func (m *TokenError) Error() string {
	return fmt.Sprintf("[%d]: Token Error, %s expected but %s found", m.token.position, m.expected.GetName(), m.token._type.GetName())
}

type FunctionNotFound struct {
	name  string
	token Token
}

func (m *FunctionNotFound) Error() string {
	return fmt.Sprintf(`[%d]: Function "%s" not found`, m.token.position, m.name)
}

type NodeError struct{}

func (m *NodeError) Error() string {
	return ""
}
