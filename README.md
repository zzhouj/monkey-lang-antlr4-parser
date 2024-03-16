# Monkey Lang

The Monkey programming language designed in [_Writing An Interpreter In Go_](https://interpreterbook.com) and [_Writing a Compiler in Go_](https://compilerbook.com) by [Thorsten Ball](https://github.com/mrnugget). I **_highly_** recommend picking up a copy of his books.

# ANTLR4 parser

1. Write a ANTLR4 grammer rule file for Monkey Lang:
```
grammar Monkey;

IDENT  : LETER (LETER | DIFIT)* ;
INT    : '0' | [1-9] DIFIT* ;
STRING : '"' (~["])* '"' ;
WS     : [ \t\r\n]+ -> skip ;

fragment LETER : [a-zA-Z_] ;
fragment DIFIT : [0-9] ;

prog : stat* ;

stat : 'let' IDENT '=' expr  ';'?  # letStat
     | 'return' expr ';'?          # retStat
     | expr ';'?                   # exprStat
     ;

expr : expr '[' expr ']'                          # indexExpr
     | expr '(' exprs ')'                         # callExpr
     | op=('-'|'!') expr                          # unOpExpr
     | expr op=('*'|'/') expr                     # mulDivExpr
     | expr op=('+'|'-') expr                     # addSubExpr
     | expr op=('<'|'>') expr                     # ltGtExpr
     | expr op=('=='|'!=') expr                   # eqNeExpr
     | 'if' '(' expr ')' block ('else' block)?    # ifExpr
     | 'fn' '(' params ')' block                  # fnLit
     | '[' exprs ']'                              # arrLit
     | '{' pairs '}'                              # hashLit
     | IDENT                                      # ident
     | INT                                        # intLit
     | STRING                                     # strLit
     | ('true'|'false')                           # boolLit
     | '(' expr ')'                               # parenExpr
     ;

exprs : expr (',' expr)* | ;

block : '{' stat* '}' ;

params : IDENT (',' IDENT)* | ;

pairs : pair (',' pair)* | ;

pair: expr ':' expr ;
```


2. Generate parser source code for golang:
```
java -jar /usr/local/lib/antlr-4.13.1-complete.jar -Dlanguage=Go -o monkey Monkey.g4
```

3. Extends *monkey.BaseMonkeyListener to translate antlr4 AST to monkey AST
