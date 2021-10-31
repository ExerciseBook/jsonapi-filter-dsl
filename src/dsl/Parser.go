package dsl

type FilterParser struct {
	lexer          FilterLexer
	customFunction map[string]FunctionType
}

func CreateParser(lexer FilterLexer, customFunction map[string]FunctionType) (FilterParser, error) {
	ret := FilterParser{
		lexer: lexer,
	}
	if customFunction == nil {
		ret.customFunction = FunctionDefinition
	} else {
		ret.customFunction = customFunction
	}
	_, err := ret.lexer.FetchNextNonBlankToken()
	if _, ok := err.(*EndOfFile); ok {
		return ret, nil
	}
	return ret, err
}

func (parser *FilterParser) ParseExpression() (AstNode, error) {
	if parser.lexer.IsEnded() {
		return TerminalNode{}, nil
	}

	// 消费函数名
	id := parser.lexer.nowToken
	if id._type != TokenID {
		return nil, &TokenError{
			expected: TokenID,
			token:    id,
		}
	}
	_, err := parser.GetFunction(id.value, id)
	if err != nil {
		return nil, err
	}

	// 读取括号
	_, err = parser.lexer.FetchNextNonBlankToken()
	if err != nil {
		return nil, err
	}
	return parser.ParseFunction(id, true)
}

func (parser *FilterParser) ParseFunction(id Token, ignoreEOF bool) (AstNode, error) {
	var ret = new(NonTerminalNode)
	ret.children = append(ret.children, TerminalNode{_type: NodeFunction, token: id})

	// 消费括号
	open := parser.lexer.nowToken
	if open._type != TokenLeftParen {
		return nil, &TokenError{
			expected: TokenLeftParen,
			token:    open,
		}
	}
	ret.children = append(ret.children, TerminalNode{_type: NodeLeftParen, token: open})

	// 读取括号的下一个元素
	_, err := parser.lexer.FetchNextNonBlankToken()
	if err != nil {
		return nil, err
	}

	for i := 0; ; i++ {
		param, err := parser.ParseGeneric()
		if err != nil {
			return nil, err
		}
		ret.children = append(ret.children, param)

		if parser.lexer.nowToken._type == TokenComma {
			ret.children = append(ret.children, TerminalNode{_type: NodeComma, token: parser.lexer.nowToken})
			_, err = parser.lexer.FetchNextNonBlankToken()
			if err != nil {
				return nil, err
			}
			continue
		}

		if parser.lexer.nowToken._type == TokenRightParen {
			break
		}

		return nil, &TokenError{
			expected: TokenComma,
			token:    parser.lexer.nowToken,
		}
	}

	_close := parser.lexer.nowToken
	_, err = parser.lexer.FetchNextNonBlankToken()
	if err != nil {
		if !ignoreEOF {
			return nil, err
		}
		if _, ok := err.(*EndOfFile); !ok {
			return nil, err
		}
	}
	if _close._type != TokenRightParen {
		return nil, &TokenError{
			expected: TokenLeftParen,
			token:    _close,
		}
	}
	ret.children = append(ret.children, TerminalNode{_type: NodeRightParen, token: _close})

	return ret, nil
}

func (parser *FilterParser) ParseGeneric() (AstNode, error) {
	// 字面量
	if parser.lexer.nowToken._type == TokenLiteralValue {
		ret := TerminalNode{_type: NodeLiteralValue, token: parser.lexer.nowToken}
		_, err := parser.lexer.FetchNextNonBlankToken()
		if err != nil {
			return nil, err
		}
		return ret, nil
	}

	if parser.lexer.nowToken._type != TokenID {
		return nil, &TokenError{
			expected: TokenID,
			token:    parser.lexer.nowToken,
		}
	}
	// 读取 ID
	id := parser.lexer.nowToken
	// 消费 ID
	_, err := parser.lexer.FetchNextNonBlankToken()
	if err != nil {
		return nil, err
	}

	// ID + '(' => 函数
	if parser.lexer.nowToken._type == TokenLeftParen {
		_, err := parser.GetFunction(id.value, id)
		if err != nil {
			return nil, err
		}
		return parser.ParseFunction(id, false)
	}

	// ID + '.' => 字段名
	//if parser.lexer.nowToken._type == TokenDot {
		// TODO
	//}

	// ID + ',' => 直接走人
	if parser.lexer.nowToken._type == TokenComma {
		return TerminalNode{_type: NodeID, token: id}, nil
	}

	// ID + ')' => 直接走人
	if parser.lexer.nowToken._type == TokenRightParen {
		return TerminalNode{_type: NodeID, token: id}, nil
	}

	// ID 的 Next 没有其他的了
	return nil, &TokenError{
		expected: TokenComma,
		token:    parser.lexer.nowToken,
	}
}

func (parser *FilterParser) GetFunction(name string, token Token) (FunctionType, error) {
	ret, ok := parser.customFunction[name]
	if ok {
		return ret, nil
	}

	ret, ok = FunctionDefinition[name]
	if ok {
		return ret, nil
	}

	return FunctionType{}, &FunctionNotFound{name: name, token: token}
}
