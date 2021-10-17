package dsl

import "fmt"

type EndOfFile struct{}

func (m *EndOfFile) Error() string {
	return ""
}

type TokenError struct {
	expected int
	actual   int
}

func (m *TokenError) Error() string {
	return fmt.Sprintf("Token Error, %d expected but %d found", m.expected, m.actual)
}

type FunctionNotFound struct {
	name string
}

func (m *FunctionNotFound) Error() string {
	return fmt.Sprintf("Function %s not found", m.name)
}

type NodeError struct{}

func (m *NodeError) Error() string {
	return ""
}
