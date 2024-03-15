// Code generated from Monkey.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Monkey

import "github.com/antlr4-go/antlr/v4"

// MonkeyListener is a complete listener for a parse tree produced by MonkeyParser.
type MonkeyListener interface {
	antlr.ParseTreeListener

	// EnterProg is called when entering the prog production.
	EnterProg(c *ProgContext)

	// EnterStat is called when entering the stat production.
	EnterStat(c *StatContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterLetStat is called when entering the letStat production.
	EnterLetStat(c *LetStatContext)

	// EnterRetStat is called when entering the retStat production.
	EnterRetStat(c *RetStatContext)

	// EnterExprStat is called when entering the exprStat production.
	EnterExprStat(c *ExprStatContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterExprs is called when entering the exprs production.
	EnterExprs(c *ExprsContext)

	// EnterParams is called when entering the params production.
	EnterParams(c *ParamsContext)

	// EnterPairs is called when entering the pairs production.
	EnterPairs(c *PairsContext)

	// EnterPair is called when entering the pair production.
	EnterPair(c *PairContext)

	// ExitProg is called when exiting the prog production.
	ExitProg(c *ProgContext)

	// ExitStat is called when exiting the stat production.
	ExitStat(c *StatContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitLetStat is called when exiting the letStat production.
	ExitLetStat(c *LetStatContext)

	// ExitRetStat is called when exiting the retStat production.
	ExitRetStat(c *RetStatContext)

	// ExitExprStat is called when exiting the exprStat production.
	ExitExprStat(c *ExprStatContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitExprs is called when exiting the exprs production.
	ExitExprs(c *ExprsContext)

	// ExitParams is called when exiting the params production.
	ExitParams(c *ParamsContext)

	// ExitPairs is called when exiting the pairs production.
	ExitPairs(c *PairsContext)

	// ExitPair is called when exiting the pair production.
	ExitPair(c *PairContext)
}
