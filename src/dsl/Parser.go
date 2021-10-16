package dsl

type FilterParser struct {
	lexer          FilterLexer
	customFunction map[string]FunctionType
}

func CreateParser(lexer FilterLexer, customFunction map[string]FunctionType) (FilterParser, error) {
	ret := FilterParser{
		lexer:          lexer,
		customFunction: customFunction,
	}
	_, err := ret.lexer.FetchNextNonBlankToken()
	return ret, err
}

func (parser *FilterParser) ParseExpression() (AstNode, error) {
	var ret = new(NonTerminalNode)

	id, err := parser.lexer.FetchNextNonBlankToken()
	if err != nil {
		return nil, err
	}
	if id._type != TokenID {
		return nil, &TokenError{
			expected: TokenID,
			actual:   id._type,
		}
	}
	function, err := parser.GetFunction(id.value)
	if err != nil {
		return nil, err
	}
	ret.children = append(ret.children, TerminalNode{_type: NodeID, token: id})


	open, err := parser.lexer.FetchNextToken()
	if err != nil {
		return nil, err
	}
	if open._type != TokenLeftParen {
		return nil, &TokenError{
			expected: TokenLeftParen,
			actual:   open._type,
		}
	}
	ret.children = append(ret.children, TerminalNode{_type: NodeLeftParen, token: open})

	for i := 0; i < function.ParameterCount; i++ {
		// 0, 1, 2, 3
		// parser.
		if i != function.ParameterCount-1 {
			// 吃个逗号
			comma, err := parser.lexer.FetchNextToken()
			if err != nil {
				return nil, err
			}
			if comma._type != TokenComma {
				return nil, &TokenError{
					expected: TokenComma,
					actual:   comma._type,
				}
			}
			ret.children = append(ret.children, TerminalNode{_type: NodeComma, token: comma})
		}
	}

	_close, err := parser.lexer.FetchNextToken()
	if err != nil {
		return nil, err
	}
	if _close._type != TokenRightParen {
		return nil, &TokenError{
			expected: TokenLeftParen,
			actual:   _close._type,
		}
	}
	ret.children = append(ret.children, TerminalNode{_type: NodeRightParen, token: _close})

	return ret, nil
}

func (parser *FilterParser) GetFunction(name string) (FunctionType, error) {
	ret, ok := parser.customFunction[name]
	if ok {
		return ret, nil
	}

	ret, ok = FunctionDefinition[name]
	if ok {
		return ret, nil
	}

	return FunctionType{}, &FunctionNotFound{name: name}
}
