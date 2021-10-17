package dsl

import (
	"fmt"
	"testing"
)

func TestParser(t *testing.T) {
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
	}

	for _, source := range cases {
		println(source)
		lexer := CreateLexer(source)
		parser, err := CreateParser(lexer, nil)
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
