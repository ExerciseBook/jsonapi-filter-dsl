package dsl

import (
	"fmt"
	"testing"
)

func TestParser(t *testing.T) {
	var function = make(map[string]FunctionType)
	//function["and"] = FunctionType{
	//	Type:           FunctionLogistic,
	//	ParameterCount: -1, // 负数表示不限制数量
	//	Function: func() {
	//	},
	//}
	//function["equals"] = FunctionType{
	//	Type:           FunctionLogistic,
	//	ParameterCount: 2, // 负数表示不限制数量
	//	Function: func() {
	//	},
	//}
	//function["has"] = FunctionType{
	//	Type:           FunctionLogistic,
	//	ParameterCount: 1, // 负数表示不限制数量
	//	Function: func() {
	//	},
	//}

	cases := []string{
		//"equals(lastName,'Smith')",
		//"lessThan(age,'25')",
		//"lessOrEqual(lastModified,'2001-01-01')",
		//"greaterThan(duration,'6:12:14')",
		//"greaterOrEqual(percentage,'33.33')",
		//"contains(description,'cooking')",
		//"startsWith(description,'The')",
		//"endsWith(description,'End')",
		//"any(chapter,'Intro','Summary','Conclusion')",
		//"has(articles)",
		//"not(equals(lastName,null))",
		//"or(has(orders),has(invoices))",
		//"and(has(orders),has(invoices))",
		//"equals(displayName,'Brian O''Connor')",
		"and(equals(displayName,'Brian O''Connor'), has(invoices))",
		//"       and(equals(displayName,'Brian O''Connor'), has(invoices))",
	}

	for _, source := range cases {
		println(source)
		lexer := CreateLexer(source)
		parser, err := CreateParser(lexer, function)
		if err != nil {
			t.Error(err.Error())
			return
		}
		ret, err := parser.ParseExpression()
		if err != nil {
			t.Error(err.Error())
			return
		}
		println(fmt.Sprintf("%v", ret))
		println()
	}
}
