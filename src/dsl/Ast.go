package dsl

type NodeType int

const (
	NodeID           NodeType = 100 + 6
	NodeLeftParen    NodeType = 100 + 1
	NodeRightParen   NodeType = 100 + 2
	NodeComma        NodeType = 100 + 3
	NodeLiteralValue NodeType = 100 + 5
	NodeFunction     NodeType = 200
)

type AstNode interface {
	IsTerminal() bool
	GetType() NodeType
	GetChildren() []AstNode
	GetToken() Token
}

type TerminalNode struct {
	_type NodeType
	token Token
}

func (t TerminalNode) IsTerminal() bool {
	return true
}
func (t TerminalNode) GetType() NodeType {
	return t._type
}
func (t TerminalNode) GetChildren() []AstNode {
	return make([]AstNode, 0)
}
func (t TerminalNode) GetToken() Token {
	return t.token
}

type NonTerminalNode struct {
	_type    NodeType
	children []AstNode
}

func (t NonTerminalNode) IsTerminal() bool {
	return false
}
func (t NonTerminalNode) GetType() NodeType {
	return t._type
}
func (t NonTerminalNode) GetChildren() []AstNode {
	return t.children
}
func (t NonTerminalNode) GetToken() Token {
	return Token{}
}
