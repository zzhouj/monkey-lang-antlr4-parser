# Monkey Lang

The Monkey programming language designed in [_Writing An Interpreter In Go_](https://interpreterbook.com) and [_Writing a Compiler in Go_](https://compilerbook.com) by [Thorsten Ball](https://github.com/mrnugget). I **_highly_** recommend picking up a copy of his books.

# ANTLR4 parser

1. Write an ANTLR4 grammer rule file for Monkey Lang:
```antlr
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
```bash
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
```go
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
	*antlr.DefaultErrorListener
	parser *monkey.MonkeyParser

	errors []string
	values map[antlr.ParserRuleContext]any
}

func New(input string) *Parser {
	lexer := monkey.NewMonkeyLexer(antlr.NewInputStream(input))
	parser := monkey.NewMonkeyParser(antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel))

	p := &Parser{
		BaseMonkeyListener:   &monkey.BaseMonkeyListener{},
		DefaultErrorListener: &antlr.DefaultErrorListener{},
		parser:               parser,
		errors:               []string{},
		values:               map[antlr.ParserRuleContext]any{},
	}

	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(p)

	parser.RemoveErrorListeners()
	parser.AddErrorListener(p)

	return p
}

func (p *Parser) setValue(ctx antlr.ParserRuleContext, value any) {
	p.values[ctx] = value
}

func (p *Parser) getValue(ctx antlr.ParserRuleContext) any {
	return p.values[ctx]
}

func (p *Parser) ParseProgram() *ast.Program {
	prog := p.parser.Prog()
	if len(p.errors) > 0 {
		return &ast.Program{
			Statements: []ast.Statement{},
		}
	}
	antlr.NewParseTreeWalker().Walk(p, prog)
	return p.getValue(prog).(*ast.Program)
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) SyntaxError(_ antlr.Recognizer, _ interface{}, line, column int, msg string, _ antlr.RecognitionException) {
	p.errors = append(p.errors, "line "+strconv.Itoa(line)+":"+strconv.Itoa(column)+" "+msg)
}

func (p *Parser) VisitErrorNode(node antlr.ErrorNode) {
	p.errors = append(p.errors, node.GetText())
}

func (p *Parser) ExitProg(ctx *monkey.ProgContext) {
	n := len(ctx.AllStat())
	stats := make([]ast.Statement, n)
	for i := 0; i < n; i++ {
		stats[i] = p.getValue(ctx.Stat(i)).(ast.Statement)
	}
	p.setValue(ctx, &ast.Program{
		Statements: stats,
	})
}

func (p *Parser) ExitLetStat(ctx *monkey.LetStatContext) {
	value := p.getValue(ctx.Expr()).(ast.Expression)
	name := ctx.IDENT().GetText()
	if fnLit, ok := value.(*ast.FunctionLiteral); ok {
		fnLit.Name = name
	}
	p.setValue(ctx, &ast.LetStatement{
		Token: token.Token{Type: token.LET, Literal: "let"},
		Name:  newIdentifier(name),
		Value: value,
	})
}

...
```