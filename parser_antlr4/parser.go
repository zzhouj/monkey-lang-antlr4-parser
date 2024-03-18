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
	stk    []any
}

func New(input string) *Parser {
	lexer := monkey.NewMonkeyLexer(antlr.NewInputStream(input))
	parser := monkey.NewMonkeyParser(antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel))

	p := &Parser{
		BaseMonkeyListener: &monkey.BaseMonkeyListener{},
		parser:             parser,
	}

	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(p)

	parser.RemoveErrorListeners()
	parser.AddErrorListener(p)

	return p
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
	prog := p.parser.Prog()
	if len(p.errors) > 0 {
		return &ast.Program{
			Statements: []ast.Statement{},
		}
	}
	antlr.NewParseTreeWalker().Walk(p, prog)
	return p.pop().(*ast.Program)
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

func (p *Parser) ExitIndexExpr(ctx *monkey.IndexExprContext) {
	index := p.pop().(ast.Expression)
	left := p.pop().(ast.Expression)
	p.push(&ast.IndexExpression{
		Token: token.Token{Type: token.LBRACKET, Literal: "["},
		Left:  left,
		Index: index,
	})
}

func (p *Parser) ExitCallExpr(ctx *monkey.CallExprContext) {
	args := p.pop().([]ast.Expression)
	fnExpr := p.pop().(ast.Expression)
	p.push(&ast.CallExpression{
		Token:     token.Token{Type: token.LPAREN, Literal: "("},
		Function:  fnExpr,
		Arguments: args,
	})
}

func (p *Parser) ExitUnOpExpr(ctx *monkey.UnOpExprContext) {
	right := p.pop().(ast.Expression)
	op := ctx.GetOp().GetText()
	p.push(newPrefixExpression(op, right))
}

func (p *Parser) ExitMulDivExpr(ctx *monkey.MulDivExprContext) {
	right := p.pop().(ast.Expression)
	left := p.pop().(ast.Expression)
	op := ctx.GetOp().GetText()
	p.push(newInfixExpression(op, left, right))
}

func (p *Parser) ExitAddSubExpr(ctx *monkey.AddSubExprContext) {
	right := p.pop().(ast.Expression)
	left := p.pop().(ast.Expression)
	op := ctx.GetOp().GetText()
	p.push(newInfixExpression(op, left, right))
}

func (p *Parser) ExitLtGtExpr(ctx *monkey.LtGtExprContext) {
	right := p.pop().(ast.Expression)
	left := p.pop().(ast.Expression)
	op := ctx.GetOp().GetText()
	p.push(newInfixExpression(op, left, right))
}

func (p *Parser) ExitEqNeExpr(ctx *monkey.EqNeExprContext) {
	right := p.pop().(ast.Expression)
	left := p.pop().(ast.Expression)
	op := ctx.GetOp().GetText()
	p.push(newInfixExpression(op, left, right))
}

func (p *Parser) ExitIfExpr(ctx *monkey.IfExprContext) {
	var conseq, alter *ast.BlockStatement
	switch len(ctx.AllBlock()) {
	case 1:
		conseq = p.pop().(*ast.BlockStatement)
	case 2:
		alter = p.pop().(*ast.BlockStatement)
		conseq = p.pop().(*ast.BlockStatement)
	}
	cond := p.pop().(ast.Expression)
	p.push(&ast.IfExpression{
		Token:       token.Token{Type: token.IF, Literal: "if"},
		Condition:   cond,
		Consequence: conseq,
		Alternative: alter,
	})
}

func (p *Parser) ExitFnLit(ctx *monkey.FnLitContext) {
	body := p.pop().(*ast.BlockStatement)
	params := p.pop().([]*ast.Identifier)
	p.push(&ast.FunctionLiteral{
		Token:      token.Token{Type: token.FUNCTION, Literal: "fn"},
		Parameters: params,
		Body:       body,
	})
}

func (p *Parser) ExitArrLit(ctx *monkey.ArrLitContext) {
	elements := p.pop().([]ast.Expression)
	p.push(&ast.ArrayLiteral{
		Token:    token.Token{Type: token.LBRACKET, Literal: "["},
		Elements: elements,
	})
}

func (p *Parser) ExitHashLit(ctx *monkey.HashLitContext) {
	pairs := p.pop().(map[ast.Expression]ast.Expression)
	p.push(&ast.HashLiteral{
		Token: token.Token{Type: token.LBRACE, Literal: "{"},
		Pairs: pairs,
	})
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

func (p *Parser) ExitExprs(ctx *monkey.ExprsContext) {
	n := len(ctx.AllExpr())
	exprs := make([]ast.Expression, n)
	for i := n - 1; i >= 0; i-- {
		exprs[i] = p.pop().(ast.Expression)
	}
	p.push(exprs)
}

func (p *Parser) ExitBlock(ctx *monkey.BlockContext) {
	n := len(ctx.AllStat())
	stats := make([]ast.Statement, n)
	for i := n - 1; i >= 0; i-- {
		stats[i] = p.pop().(ast.Statement)
	}
	p.push(&ast.BlockStatement{
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
	p.push(params)
}

func (p *Parser) ExitPairs(ctx *monkey.PairsContext) {
	n := len(ctx.AllPair())
	keys := make([]ast.Expression, n)
	values := make([]ast.Expression, n)
	for i := n - 1; i >= 0; i-- {
		values[i] = p.pop().(ast.Expression)
		keys[i] = p.pop().(ast.Expression)
	}
	pairs := map[ast.Expression]ast.Expression{}
	for i := 0; i < n; i++ {
		pairs[keys[i]] = values[i]
	}
	p.push(pairs)
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
