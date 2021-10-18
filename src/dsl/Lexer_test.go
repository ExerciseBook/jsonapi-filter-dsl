package dsl

import (
	"fmt"
	"testing"
)

func TestLexer(t *testing.T) {
	cases := []string{
		"equals(lastName,'Smith')",
		"lessThan(age,'25')",
		"lessOrEqual(lastModified,'2001-01-01')",
		"greaterThan(duration,'6:12:14')",
		"greaterOrEqual(percentage,'33.33')",
		"contains(description,'cooking')",
		"startsWith(description,'The')",
		"endsWith(description,'End')",
		"any(chapter,'Intro','Summary','Conclusion')",
		"has(articles)",
		"not(equals(lastName,null))",
		"or(has(orders),has(invoices))",
		"and(has(orders),has(invoices))",
		"equals(displayName,'Brian O''Connor')",
	}

	for _, source := range cases {
		println(source)
		lexer := CreateLexer(source)
		for {
			ret, err := lexer.FetchNextNonBlankToken()
			if err != nil {
				break
			}
			println(fmt.Sprintf("%v", ret))
		}
		println()
	}
}
