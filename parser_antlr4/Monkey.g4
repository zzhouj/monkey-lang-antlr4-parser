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
