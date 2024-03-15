// Code generated from Monkey.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Monkey

import "github.com/antlr4-go/antlr/v4"

// BaseMonkeyListener is a complete listener for a parse tree produced by MonkeyParser.
type BaseMonkeyListener struct{}

var _ MonkeyListener = &BaseMonkeyListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseMonkeyListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseMonkeyListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseMonkeyListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseMonkeyListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProg is called when production prog is entered.
func (s *BaseMonkeyListener) EnterProg(ctx *ProgContext) {}

// ExitProg is called when production prog is exited.
func (s *BaseMonkeyListener) ExitProg(ctx *ProgContext) {}

// EnterStat is called when production stat is entered.
func (s *BaseMonkeyListener) EnterStat(ctx *StatContext) {}

// ExitStat is called when production stat is exited.
func (s *BaseMonkeyListener) ExitStat(ctx *StatContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseMonkeyListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseMonkeyListener) ExitBlock(ctx *BlockContext) {}

// EnterLetStat is called when production letStat is entered.
func (s *BaseMonkeyListener) EnterLetStat(ctx *LetStatContext) {}

// ExitLetStat is called when production letStat is exited.
func (s *BaseMonkeyListener) ExitLetStat(ctx *LetStatContext) {}

// EnterRetStat is called when production retStat is entered.
func (s *BaseMonkeyListener) EnterRetStat(ctx *RetStatContext) {}

// ExitRetStat is called when production retStat is exited.
func (s *BaseMonkeyListener) ExitRetStat(ctx *RetStatContext) {}

// EnterExprStat is called when production exprStat is entered.
func (s *BaseMonkeyListener) EnterExprStat(ctx *ExprStatContext) {}

// ExitExprStat is called when production exprStat is exited.
func (s *BaseMonkeyListener) ExitExprStat(ctx *ExprStatContext) {}

// EnterExpr is called when production expr is entered.
func (s *BaseMonkeyListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseMonkeyListener) ExitExpr(ctx *ExprContext) {}

// EnterExprs is called when production exprs is entered.
func (s *BaseMonkeyListener) EnterExprs(ctx *ExprsContext) {}

// ExitExprs is called when production exprs is exited.
func (s *BaseMonkeyListener) ExitExprs(ctx *ExprsContext) {}

// EnterParams is called when production params is entered.
func (s *BaseMonkeyListener) EnterParams(ctx *ParamsContext) {}

// ExitParams is called when production params is exited.
func (s *BaseMonkeyListener) ExitParams(ctx *ParamsContext) {}

// EnterPairs is called when production pairs is entered.
func (s *BaseMonkeyListener) EnterPairs(ctx *PairsContext) {}

// ExitPairs is called when production pairs is exited.
func (s *BaseMonkeyListener) ExitPairs(ctx *PairsContext) {}

// EnterPair is called when production pair is entered.
func (s *BaseMonkeyListener) EnterPair(ctx *PairContext) {}

// ExitPair is called when production pair is exited.
func (s *BaseMonkeyListener) ExitPair(ctx *PairContext) {}
