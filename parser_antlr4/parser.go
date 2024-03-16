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
	stk    []ast.Node
}

func New(input string) *Parser {
	lexer := monkey.NewMonkeyLexer(antlr.NewInputStream(input))
	parser := monkey.NewMonkeyParser(antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel))

	return &Parser{
		BaseMonkeyListener: &monkey.BaseMonkeyListener{},
		parser:             parser,
	}
}

func (p *Parser) push(n ast.Node) {
	p.stk = append(p.stk, n)
}

func (p *Parser) pop() ast.Node {
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
	identLit := ctx.IDENT().GetText()
	name := &ast.Identifier{
		Token: token.Token{Type: token.IDENT, Literal: identLit},
		Value: identLit,
	}
	p.push(&ast.LetStatement{
		Token: token.Token{Type: token.LET, Literal: "let"},
		Name:  name,
		Value: value,
	})
}

func (p *Parser) ExitIdent(ctx *monkey.IdentContext) {
	identLit := ctx.IDENT().GetText()
	p.push(&ast.Identifier{
		Token: token.Token{Type: token.IDENT, Literal: identLit},
		Value: identLit,
	})
}

func (p *Parser) ExitIntLit(ctx *monkey.IntLitContext) {
	intLit := ctx.INT().GetText()
	value, _ := strconv.ParseInt(intLit, 10, 64)
	p.push(&ast.IntegerLiteral{
		Token: token.Token{Type: token.INT, Literal: intLit},
		Value: value,
	})
}
