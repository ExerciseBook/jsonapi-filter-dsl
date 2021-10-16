grammar Filter;

// https://www.jsonapi.net/usage/reading/filtering.html

filterExpression:
    notExpression
    | logicalExpression
    | comparisonExpression
    | matchTextExpression
    | anyExpression
    | hasExpression;

// 一元逻辑函数
notExpression:
    'not' LPAREN filterExpression RPAREN;

// 多元逻辑函数
logicalExpression:
    ( 'and' | 'or' ) LPAREN filterExpression ( COMMA filterExpression )* RPAREN;

// 二元比较运算
comparisonExpression:
    ( 'equals' | 'greaterThan' | 'greaterOrEqual' | 'lessThan' | 'lessOrEqual' ) LPAREN (
        fieldChain
    ) COMMA (
        literalConstant | 'null' | fieldChain
    ) RPAREN;

// 函数调用 1键 1值
matchTextExpression:
    ( 'contains' | 'startsWith' | 'endsWith' ) LPAREN fieldChain COMMA literalConstant RPAREN;

// 函数调用 1键 多值
anyExpression:
    'any' LPAREN fieldChain COMMA literalConstant ( COMMA literalConstant )+ RPAREN;

// 函数调用 1键 1值
hasExpression:
    'has' LPAREN fieldChain ( COMMA filterExpression )? RPAREN;

fieldChain:
    FIELD ( '.' FIELD )*;

literalConstant:
    ESCAPED_TEXT;

LPAREN: '(';
RPAREN: ')';
COMMA: ',';

fragment OUTER_FIELD_CHARACTER: [A-Za-z0-9];
fragment INNER_FIELD_CHARACTER: [A-Za-z0-9_-];
FIELD: OUTER_FIELD_CHARACTER ( INNER_FIELD_CHARACTER* OUTER_FIELD_CHARACTER )?;

ESCAPED_TEXT: '\'' ( ~['] | '\'\'' )* '\'' ;

LINE_BREAKS: [\r\n]+ -> skip;

WS: [ \t\n] -> skip;
