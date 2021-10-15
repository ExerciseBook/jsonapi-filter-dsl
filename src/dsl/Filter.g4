grammar Filter;

// https://www.jsonapi.net/usage/reading/filtering.html

filterExpression:
    notExpression
    | logicalExpression
    | comparisonExpression
    | matchTextExpression
    | anyExpression
    | hasExpression;

notExpression:
    'not' LPAREN filterExpression RPAREN;

logicalExpression:
    ( 'and' | 'or' ) LPAREN filterExpression ( COMMA filterExpression )* RPAREN;

comparisonExpression:
    ( 'equals' | 'greaterThan' | 'greaterOrEqual' | 'lessThan' | 'lessOrEqual' ) LPAREN (
        countExpression | fieldChain
    ) COMMA (
        countExpression | literalConstant | 'null' | fieldChain
    ) RPAREN;

matchTextExpression:
    ( 'contains' | 'startsWith' | 'endsWith' ) LPAREN fieldChain COMMA literalConstant RPAREN;

anyExpression:
    'any' LPAREN fieldChain COMMA literalConstant ( COMMA literalConstant )+ RPAREN;

hasExpression:
    'has' LPAREN fieldChain ( COMMA filterExpression )? RPAREN;

countExpression:
    'count' LPAREN fieldChain RPAREN;

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
