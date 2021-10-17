package dsl

type AstNode interface {
	IsTerminal() bool
	GetType() int
	GetChildren() []AstNode
	GetToken() Token
}

type TerminalNode struct {
	_type int
	token Token
}

func (t TerminalNode) IsTerminal() bool {
	return true
}
func (t TerminalNode) GetType() int {
	return t._type
}
func (t TerminalNode) GetChildren() []AstNode {
	return make([]AstNode, 0)
}
func (t TerminalNode) GetToken() Token {
	return t.token
}

type NonTerminalNode struct {
	_type    int
	children []AstNode
}

func (t NonTerminalNode) IsTerminal() bool {
	return false
}
func (t NonTerminalNode) GetType() int {
	return t._type
}
func (t NonTerminalNode) GetChildren() []AstNode {
	return t.children
}
func (t NonTerminalNode) GetToken() Token {
	return Token{}
}

const NodeID = 100 + TokenID
const NodeLeftParen = 100 + TokenLeftParen
const NodeRightParen = 100 + TokenRightParen
const NodeComma = 100 + TokenComma
const NodeLiteralValue = 100 + TokenLiteralValue
const NodeFunction = 200
