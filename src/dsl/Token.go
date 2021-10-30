package dsl

type TokenType int

const (
	TokenWhiteSpace   TokenType = 0
	TokenLeftParen    TokenType = 1
	TokenRightParen   TokenType = 2
	TokenComma        TokenType = 3
	TokenDot          TokenType = 4
	TokenLiteralValue TokenType = 5
	TokenID           TokenType = 6
)

func (t TokenType) GetName() string {
	switch t {
	case TokenWhiteSpace:
		return "TokenWhiteSpace"
	case TokenLeftParen:
		return "TokenLeftParen"
	case TokenRightParen:
		return "TokenRightParen"
	case TokenComma:
		return "TokenComma"
	case TokenDot:
		return "TokenDot"
	case TokenLiteralValue:
		return "TokenLiteralValue"
	case TokenID:
		return "TokenID"
	}
	return ""
}

type Token struct {
	value    string
	position int
	_type    TokenType
}

func (token Token) GetValue() string {
	return token.value
}

func (token Token) GetPosition() int {
	return token.position
}

func (token Token) GetType() TokenType {
	return token._type
}
