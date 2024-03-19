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

func (p *Parser) ExitRetStat(ctx *monkey.RetStatContext) {
	returnValue := p.getValue(ctx.Expr()).(ast.Expression)
	p.setValue(ctx, &ast.ReturnStatement{
		Token:       token.Token{Type: token.RETURN, Literal: "return"},
		ReturnValue: returnValue,
	})
}

func (p *Parser) ExitExprStat(ctx *monkey.ExprStatContext) {
	expr := p.getValue(ctx.Expr()).(ast.Expression)
	p.setValue(ctx, &ast.ExpressionStatement{
		Token:      token.Token{}, // TODO
		Expression: expr,
	})
}

func (p *Parser) ExitIndexExpr(ctx *monkey.IndexExprContext) {
	index := p.getValue(ctx.Expr(1)).(ast.Expression)
	left := p.getValue(ctx.Expr(0)).(ast.Expression)
	p.setValue(ctx, &ast.IndexExpression{
		Token: token.Token{Type: token.LBRACKET, Literal: "["},
		Left:  left,
		Index: index,
	})
}

func (p *Parser) ExitCallExpr(ctx *monkey.CallExprContext) {
	args := p.getValue(ctx.Exprs()).([]ast.Expression)
	fnExpr := p.getValue(ctx.Expr()).(ast.Expression)
	p.setValue(ctx, &ast.CallExpression{
		Token:     token.Token{Type: token.LPAREN, Literal: "("},
		Function:  fnExpr,
		Arguments: args,
	})
}

func (p *Parser) ExitUnOpExpr(ctx *monkey.UnOpExprContext) {
	right := p.getValue(ctx.Expr()).(ast.Expression)
	op := ctx.GetOp().GetText()
	p.setValue(ctx, newPrefixExpression(op, right))
}

func (p *Parser) ExitMulDivExpr(ctx *monkey.MulDivExprContext) {
	right := p.getValue(ctx.Expr(1)).(ast.Expression)
	left := p.getValue(ctx.Expr(0)).(ast.Expression)
	op := ctx.GetOp().GetText()
	p.setValue(ctx, newInfixExpression(op, left, right))
}

func (p *Parser) ExitAddSubExpr(ctx *monkey.AddSubExprContext) {
	right := p.getValue(ctx.Expr(1)).(ast.Expression)
	left := p.getValue(ctx.Expr(0)).(ast.Expression)
	op := ctx.GetOp().GetText()
	p.setValue(ctx, newInfixExpression(op, left, right))
}

func (p *Parser) ExitLtGtExpr(ctx *monkey.LtGtExprContext) {
	right := p.getValue(ctx.Expr(1)).(ast.Expression)
	left := p.getValue(ctx.Expr(0)).(ast.Expression)
	op := ctx.GetOp().GetText()
	p.setValue(ctx, newInfixExpression(op, left, right))
}

func (p *Parser) ExitEqNeExpr(ctx *monkey.EqNeExprContext) {
	right := p.getValue(ctx.Expr(1)).(ast.Expression)
	left := p.getValue(ctx.Expr(0)).(ast.Expression)
	op := ctx.GetOp().GetText()
	p.setValue(ctx, newInfixExpression(op, left, right))
}

func (p *Parser) ExitIfExpr(ctx *monkey.IfExprContext) {
	var conseq, alter *ast.BlockStatement
	switch len(ctx.AllBlock()) {
	case 1:
		conseq = p.getValue(ctx.Block(0)).(*ast.BlockStatement)
	case 2:
		alter = p.getValue(ctx.Block(1)).(*ast.BlockStatement)
		conseq = p.getValue(ctx.Block(0)).(*ast.BlockStatement)
	}
	cond := p.getValue(ctx.Expr()).(ast.Expression)
	p.setValue(ctx, &ast.IfExpression{
		Token:       token.Token{Type: token.IF, Literal: "if"},
		Condition:   cond,
		Consequence: conseq,
		Alternative: alter,
	})
}

func (p *Parser) ExitFnLit(ctx *monkey.FnLitContext) {
	body := p.getValue(ctx.Block()).(*ast.BlockStatement)
	params := p.getValue(ctx.Params()).([]*ast.Identifier)
	p.setValue(ctx, &ast.FunctionLiteral{
		Token:      token.Token{Type: token.FUNCTION, Literal: "fn"},
		Parameters: params,
		Body:       body,
	})
}

func (p *Parser) ExitArrLit(ctx *monkey.ArrLitContext) {
	elements := p.getValue(ctx.Exprs()).([]ast.Expression)
	p.setValue(ctx, &ast.ArrayLiteral{
		Token:    token.Token{Type: token.LBRACKET, Literal: "["},
		Elements: elements,
	})
}

func (p *Parser) ExitHashLit(ctx *monkey.HashLitContext) {
	pairs := p.getValue(ctx.Pairs()).(map[ast.Expression]ast.Expression)
	p.setValue(ctx, &ast.HashLiteral{
		Token: token.Token{Type: token.LBRACE, Literal: "{"},
		Pairs: pairs,
	})
}

func (p *Parser) ExitIdent(ctx *monkey.IdentContext) {
	p.setValue(ctx, newIdentifier(ctx.IDENT().GetText()))
}

func (p *Parser) ExitIntLit(ctx *monkey.IntLitContext) {
	intLit := ctx.INT().GetText()
	value, _ := strconv.ParseInt(intLit, 10, 64)
	p.setValue(ctx, &ast.IntegerLiteral{
		Token: token.Token{Type: token.INT, Literal: intLit},
		Value: value,
	})
}

func (p *Parser) ExitStrLit(ctx *monkey.StrLitContext) {
	strLit := ctx.STRING().GetText()
	if len(strLit) >= 2 {
		strLit = strLit[1 : len(strLit)-1]
	}
	p.setValue(ctx, &ast.StringLiteral{
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
	p.setValue(ctx, &ast.BooleanLiteral{
		Token: token.Token{Type: tokenType, Literal: boolLit},
		Value: value,
	})
}

func (p *Parser) ExitParenExpr(ctx *monkey.ParenExprContext) {
	p.setValue(ctx, p.getValue(ctx.Expr()))
}

func (p *Parser) ExitExprs(ctx *monkey.ExprsContext) {
	n := len(ctx.AllExpr())
	exprs := make([]ast.Expression, n)
	for i := 0; i < n; i++ {
		exprs[i] = p.getValue(ctx.Expr(i)).(ast.Expression)
	}
	p.setValue(ctx, exprs)
}

func (p *Parser) ExitBlock(ctx *monkey.BlockContext) {
	n := len(ctx.AllStat())
	stats := make([]ast.Statement, n)
	for i := 0; i < n; i++ {
		stats[i] = p.getValue(ctx.Stat(i)).(ast.Statement)
	}
	p.setValue(ctx, &ast.BlockStatement{
		Token:      token.Token{Type: token.LBRACE, Literal: "{"},
		Statements: stats,
	})
}

func (p *Parser) ExitParams(ctx *monkey.ParamsContext) {
	n := len(ctx.AllIDENT())
	params := make([]*ast.Identifier, n)
	for i := 0; i < n; i++ {
		params[i] = newIdentifier(ctx.IDENT(i).GetText())
	}
	p.setValue(ctx, params)
}

func (p *Parser) ExitPairs(ctx *monkey.PairsContext) {
	n := len(ctx.AllPair())
	pairs := map[ast.Expression]ast.Expression{}
	for i := 0; i < n; i++ {
		pair := p.getValue(ctx.Pair(i)).([]ast.Expression)
		pairs[pair[0]] = pair[1]
	}
	p.setValue(ctx, pairs)
}

func (p *Parser) ExitPair(ctx *monkey.PairContext) {
	pair := make([]ast.Expression, 2)
	for i := range pair {
		pair[i] = p.getValue(ctx.Expr(i)).(ast.Expression)
	}
	p.setValue(ctx, pair)
}

func newIdentifier(identLit string) *ast.Identifier {
	return &ast.Identifier{
		Token: token.Token{Type: token.IDENT, Literal: identLit},
		Value: identLit,
	}
}

func newPrefixExpression(op string, right ast.Expression) *ast.PrefixExpression {
	var tokenType token.TokenType
	switch op {
	case "-":
		tokenType = token.MINUS
	case "!":
		tokenType = token.BANG
	}
	return &ast.PrefixExpression{
		Token:    token.Token{Type: tokenType, Literal: op},
		Operator: op,
		Right:    right,
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
