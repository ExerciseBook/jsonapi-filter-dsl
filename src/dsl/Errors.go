package dsl

type EndOfFile struct{}

func (m *EndOfFile) Error() string {
	return ""
}

type TokenError struct {
	expected int
	actual   int
}

func (m *TokenError) Error() string {
	return ""
}

type FunctionNotFound struct {
	name string
}

func (m *FunctionNotFound) Error() string {
	return ""
}
