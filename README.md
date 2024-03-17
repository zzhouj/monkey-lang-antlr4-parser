# Monkey Lang

The Monkey programming language designed in [_Writing An Interpreter In Go_](https://interpreterbook.com) and [_Writing a Compiler in Go_](https://compilerbook.com) by [Thorsten Ball](https://github.com/mrnugget). I **_highly_** recommend picking up a copy of his books.

# ANTLR4 parser

1. Write an ANTLR4 grammer rule file for Monkey Lang:
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

2. Generate Golang target source code for Lexer & Parser of Monkey Lang:
```
% java -jar /usr/local/lib/antlr-4.13.1-complete.jar -Dlanguage=Go -o monkey Monkey.g4

% ls monkey 
Monkey.interp
Monkey.tokens
MonkeyLexer.interp
MonkeyLexer.tokens
monkey_base_listener.go
monkey_lexer.go
monkey_listener.go
monkey_parser.go
```

3. Extends *monkey.BaseMonkeyListener to translate antlr4 AST to monkey AST:
```
package parser_antlr4

import (
	"monkey/ast"
	monkey "monkey/parser_antlr4/monkey"
	"monkey/token"
	"strconv"

	"github.com/antlr4-go/antlr/v4"
)

type Parser struct {
	*monkey.BaseMonkeyListener
	parser *monkey.MonkeyParser

	errors []string
	stk    []any
}

func New(input string) *Parser {
	lexer := monkey.NewMonkeyLexer(antlr.NewInputStream(input))
	parser := monkey.NewMonkeyParser(antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel))

	return &Parser{
		BaseMonkeyListener: &monkey.BaseMonkeyListener{},
		parser:             parser,
	}
}

func (p *Parser) push(n any) {
	p.stk = append(p.stk, n)
}

func (p *Parser) pop() any {
	if len(p.stk) > 0 {
		result := p.stk[len(p.stk)-1]
		p.stk = p.stk[:len(p.stk)-1]
		return result
	}
	return nil
}

func (p *Parser) ParseProgram() *ast.Program {
	antlr.NewParseTreeWalker().Walk(p, p.parser.Prog())
	return p.pop().(*ast.Program)
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) VisitErrorNode(node antlr.ErrorNode) {
	p.errors = append(p.errors, node.GetText())
}

func (p *Parser) ExitProg(ctx *monkey.ProgContext) {
	n := len(ctx.AllStat())
	stats := make([]ast.Statement, n)
	for i := n - 1; i >= 0; i-- {
		stats[i] = p.pop().(ast.Statement)
	}
	p.push(&ast.Program{
		Statements: stats,
	})
}

func (p *Parser) ExitLetStat(ctx *monkey.LetStatContext) {
	value := p.pop().(ast.Expression)
	name := ctx.IDENT().GetText()
	if fnLit, ok := value.(*ast.FunctionLiteral); ok {
		fnLit.Name = name
	}
	p.push(&ast.LetStatement{
		Token: token.Token{Type: token.LET, Literal: "let"},
		Name:  newIdentifier(name),
		Value: value,
	})
}

...
```