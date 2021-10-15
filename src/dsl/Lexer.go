package dsl

import (
	"fmt"
)

type FilterLexer struct {
	source   string
	position int
	length   int

	nowToken Token
}

func CreateLexer(source string) FilterLexer {
	return FilterLexer{
		source:   source,
		length:   len(source),
		position: 0,
	}
}

func (lexer *FilterLexer) GetPosition() int {
	return lexer.position
}

func (lexer *FilterLexer) FetchNextToken() (Token, error) {
	if lexer.length <= lexer.position {
		return Token{}, fmt.Errorf("没了")
	}

	if lexer.source[lexer.position] == ' ' ||
		lexer.source[lexer.position] == '\r' ||
		lexer.source[lexer.position] == '\n' ||
		lexer.source[lexer.position] == '\t' {
		lexer.position++
		return Token{
			value:    string(lexer.source[lexer.position]),
			position: lexer.position,
			_type:    TokenWhiteSpace,
		}, nil
	}

	if lexer.source[lexer.position] == '(' {
		lexer.position++
		return Token{
			value:    "(",
			position: lexer.position,
			_type:    TokenLeftParen,
		}, nil
	}

	if lexer.source[lexer.position] == ')' {
		lexer.position++
		return Token{
			value:    ")",
			position: lexer.position,
			_type:    TokenRightParen,
		}, nil
	}

	if lexer.source[lexer.position] == ',' {
		lexer.position++
		return Token{
			value:    ",",
			position: lexer.position,
			_type:    TokenComma,
		}, nil
	}

	if lexer.source[lexer.position] == '.' {
		lexer.position++
		return Token{
			value:    ".",
			position: lexer.position,
			_type:    TokenDot,
		}, nil
	}

	if lexer.source[lexer.position] == '\'' {
		ret := Token{
			value:    "",
			position: lexer.position,
			_type:    LiteralValue,
		}

		// TODO 找个大手子帮我优化一下这里
		i := lexer.position + 1
		for i < lexer.length {
			if lexer.source[i] == '\'' {
				if i+1 < lexer.length && lexer.source[i+1] == '\'' {
					ret.value = ret.value + "'"
					i = i + 2
				} else {
					i = i + 1
					break
				}
			} else {
				ret.value = ret.value + string(lexer.source[i])
				i++
			}
		}

		lexer.position = i
		return ret, nil
	}

	ret := Token{
		position: lexer.position,
		_type:    TokenID,
	}

	i := lexer.position
	for ; i < lexer.length &&
		lexer.source[i] != ' ' &&
		lexer.source[i] != '\r' &&
		lexer.source[i] != '\n' &&
		lexer.source[i] != '\t' &&
		lexer.source[i] != '(' &&
		lexer.source[i] != ',' &&
		lexer.source[i] != ')' &&
		lexer.source[i] != '.'; i++ {
	}
	lexer.position = i
	ret.value = lexer.source[ret.position : i]

	return ret, nil
}

func (lexer *FilterLexer) FetchNextNonBlankToken() (Token, error) {
	for {
		ret, err := lexer.FetchNextToken()
		if err != nil {
			return ret, err
		}

		if ret.GetType() != TokenWhiteSpace {
			return ret, nil
		}
	}
}
