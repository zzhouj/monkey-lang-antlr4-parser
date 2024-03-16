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
	p.push(&ast.LetStatement{
		Token: token.Token{Type: token.LET, Literal: "let"},
		Name:  newIdentifier(ctx.IDENT().GetText()),
		Value: value,
	})
}

func (p *Parser) ExitRetStat(ctx *monkey.RetStatContext) {
	returnValue := p.pop().(ast.Expression)
	p.push(&ast.ReturnStatement{
		Token:       token.Token{Type: token.RETURN, Literal: "return"},
		ReturnValue: returnValue,
	})
}

func (p *Parser) ExitExprStat(ctx *monkey.ExprStatContext) {
	expr := p.pop().(ast.Expression)
	p.push(&ast.ExpressionStatement{
		Token:      token.Token{}, // TODO
		Expression: expr,
	})
}

func (p *Parser) ExitAddSubExpr(ctx *monkey.AddSubExprContext) {
	right := p.pop().(ast.Expression)
	left := p.pop().(ast.Expression)
	op := ctx.GetOp().GetText()
	p.push(newInfixExpression(op, left, right))
}

func (p *Parser) ExitIdent(ctx *monkey.IdentContext) {
	p.push(newIdentifier(ctx.IDENT().GetText()))
}

func (p *Parser) ExitIntLit(ctx *monkey.IntLitContext) {
	intLit := ctx.INT().GetText()
	value, _ := strconv.ParseInt(intLit, 10, 64)
	p.push(&ast.IntegerLiteral{
		Token: token.Token{Type: token.INT, Literal: intLit},
		Value: value,
	})
}

func (p *Parser) ExitStrLit(ctx *monkey.StrLitContext) {
	strLit := ctx.STRING().GetText()
	if len(strLit) >= 2 {
		strLit = strLit[1 : len(strLit)-1]
	}
	p.push(&ast.StringLiteral{
		Token: token.Token{Type: token.STRING, Literal: strLit},
		Value: strLit,
	})
}

func (p *Parser) ExitBoolLit(ctx *monkey.BoolLitContext) {
	boolLit := ctx.GetText()
	var tokenType token.TokenType
	var value bool
	if boolLit == "true" {
		tokenType = token.TRUE
		value = true
	} else {
		tokenType = token.FALSE
		value = false
	}
	p.push(&ast.BooleanLiteral{
		Token: token.Token{Type: tokenType, Literal: boolLit},
		Value: value,
	})
}

func newIdentifier(identLit string) *ast.Identifier {
	return &ast.Identifier{
		Token: token.Token{Type: token.IDENT, Literal: identLit},
		Value: identLit,
	}
}

func newInfixExpression(op string, left, right ast.Expression) *ast.InfixExpression {
	var tokenType token.TokenType
	switch op {
	case "+":
		tokenType = token.PLUS
	case "-":
		tokenType = token.MINUS
	case "*":
		tokenType = token.ASTERISK
	case "/":
		tokenType = token.SLASH
	case "<":
		tokenType = token.LT
	case ">":
		tokenType = token.GT
	case "==":
		tokenType = token.EQ
	case "!=":
		tokenType = token.NOT_EQ
	}
	return &ast.InfixExpression{
		Token:    token.Token{Type: tokenType, Literal: op},
		Left:     left,
		Operator: op,
		Right:    right,
	}
}
