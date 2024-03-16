// Code generated from Monkey.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Monkey

import "github.com/antlr4-go/antlr/v4"

// MonkeyListener is a complete listener for a parse tree produced by MonkeyParser.
type MonkeyListener interface {
	antlr.ParseTreeListener

	// EnterProg is called when entering the prog production.
	EnterProg(c *ProgContext)

	// EnterLetStat is called when entering the letStat production.
	EnterLetStat(c *LetStatContext)

	// EnterRetStat is called when entering the retStat production.
	EnterRetStat(c *RetStatContext)

	// EnterExprStat is called when entering the exprStat production.
	EnterExprStat(c *ExprStatContext)

	// EnterAddSubExpr is called when entering the addSubExpr production.
	EnterAddSubExpr(c *AddSubExprContext)

	// EnterIdent is called when entering the ident production.
	EnterIdent(c *IdentContext)

	// EnterParenExpr is called when entering the parenExpr production.
	EnterParenExpr(c *ParenExprContext)

	// EnterIndexExpr is called when entering the indexExpr production.
	EnterIndexExpr(c *IndexExprContext)

	// EnterStrLit is called when entering the strLit production.
	EnterStrLit(c *StrLitContext)

	// EnterBoolLit is called when entering the boolLit production.
	EnterBoolLit(c *BoolLitContext)

	// EnterArrLit is called when entering the arrLit production.
	EnterArrLit(c *ArrLitContext)

	// EnterLtGtExpr is called when entering the ltGtExpr production.
	EnterLtGtExpr(c *LtGtExprContext)

	// EnterIntLit is called when entering the intLit production.
	EnterIntLit(c *IntLitContext)

	// EnterIfExpr is called when entering the ifExpr production.
	EnterIfExpr(c *IfExprContext)

	// EnterEqNeExpr is called when entering the eqNeExpr production.
	EnterEqNeExpr(c *EqNeExprContext)

	// EnterHashLit is called when entering the hashLit production.
	EnterHashLit(c *HashLitContext)

	// EnterCallExpr is called when entering the callExpr production.
	EnterCallExpr(c *CallExprContext)

	// EnterMulDivExpr is called when entering the mulDivExpr production.
	EnterMulDivExpr(c *MulDivExprContext)

	// EnterUnOpExpr is called when entering the unOpExpr production.
	EnterUnOpExpr(c *UnOpExprContext)

	// EnterFnLit is called when entering the fnLit production.
	EnterFnLit(c *FnLitContext)

	// EnterExprs is called when entering the exprs production.
	EnterExprs(c *ExprsContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterParams is called when entering the params production.
	EnterParams(c *ParamsContext)

	// EnterPairs is called when entering the pairs production.
	EnterPairs(c *PairsContext)

	// EnterPair is called when entering the pair production.
	EnterPair(c *PairContext)

	// ExitProg is called when exiting the prog production.
	ExitProg(c *ProgContext)

	// ExitLetStat is called when exiting the letStat production.
	ExitLetStat(c *LetStatContext)

	// ExitRetStat is called when exiting the retStat production.
	ExitRetStat(c *RetStatContext)

	// ExitExprStat is called when exiting the exprStat production.
	ExitExprStat(c *ExprStatContext)

	// ExitAddSubExpr is called when exiting the addSubExpr production.
	ExitAddSubExpr(c *AddSubExprContext)

	// ExitIdent is called when exiting the ident production.
	ExitIdent(c *IdentContext)

	// ExitParenExpr is called when exiting the parenExpr production.
	ExitParenExpr(c *ParenExprContext)

	// ExitIndexExpr is called when exiting the indexExpr production.
	ExitIndexExpr(c *IndexExprContext)

	// ExitStrLit is called when exiting the strLit production.
	ExitStrLit(c *StrLitContext)

	// ExitBoolLit is called when exiting the boolLit production.
	ExitBoolLit(c *BoolLitContext)

	// ExitArrLit is called when exiting the arrLit production.
	ExitArrLit(c *ArrLitContext)

	// ExitLtGtExpr is called when exiting the ltGtExpr production.
	ExitLtGtExpr(c *LtGtExprContext)

	// ExitIntLit is called when exiting the intLit production.
	ExitIntLit(c *IntLitContext)

	// ExitIfExpr is called when exiting the ifExpr production.
	ExitIfExpr(c *IfExprContext)

	// ExitEqNeExpr is called when exiting the eqNeExpr production.
	ExitEqNeExpr(c *EqNeExprContext)

	// ExitHashLit is called when exiting the hashLit production.
	ExitHashLit(c *HashLitContext)

	// ExitCallExpr is called when exiting the callExpr production.
	ExitCallExpr(c *CallExprContext)

	// ExitMulDivExpr is called when exiting the mulDivExpr production.
	ExitMulDivExpr(c *MulDivExprContext)

	// ExitUnOpExpr is called when exiting the unOpExpr production.
	ExitUnOpExpr(c *UnOpExprContext)

	// ExitFnLit is called when exiting the fnLit production.
	ExitFnLit(c *FnLitContext)

	// ExitExprs is called when exiting the exprs production.
	ExitExprs(c *ExprsContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitParams is called when exiting the params production.
	ExitParams(c *ParamsContext)

	// ExitPairs is called when exiting the pairs production.
	ExitPairs(c *PairsContext)

	// ExitPair is called when exiting the pair production.
	ExitPair(c *PairContext)
}
