grammar Monkey;

IDENT  : LETER (LETER | DIFIT)* ;
INT    : '0' | [1-9] DIFIT* ;
STRING : '"' (~["])* '"' ;
WS     : [ \t\r\n]+ -> skip ;

fragment LETER : [a-zA-Z_] ;
fragment DIFIT : [0-9] ;

prog : stat* ;

stat : letStat
     | retStat
     | exprStat
     ;

block : '{' stat* '}' ;

letStat : 'let' IDENT '=' expr  ';'? ;

retStat : 'return' expr ';' ;

exprStat : expr ';' ;

expr : expr '[' expr ']'
     | expr '(' exprs ')'
     | op=('-'|'!') expr
     | expr op=('*'|'/') expr
     | expr op=('+'|'-') expr
     | expr op=('<'|'>') expr
     | expr op=('=='|'!=') expr
     | 'if' '(' expr ')' block ('else' block)?
     | 'fn' '(' params ')' block
     | '[' exprs ']'
     | '{' pairs '}'
     | IDENT
     | INT
     | STRING
     | ('true'|'false')
     | '(' expr ')'
     ;

exprs : expr (',' expr)* | ;

params : IDENT (',' IDENT)* | ;

pairs : pair (',' pair)* | ;

pair: expr ':' expr ;
