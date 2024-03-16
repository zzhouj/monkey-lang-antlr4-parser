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

// EnterAddSubExpr is called when production addSubExpr is entered.
func (s *BaseMonkeyListener) EnterAddSubExpr(ctx *AddSubExprContext) {}

// ExitAddSubExpr is called when production addSubExpr is exited.
func (s *BaseMonkeyListener) ExitAddSubExpr(ctx *AddSubExprContext) {}

// EnterIdent is called when production ident is entered.
func (s *BaseMonkeyListener) EnterIdent(ctx *IdentContext) {}

// ExitIdent is called when production ident is exited.
func (s *BaseMonkeyListener) ExitIdent(ctx *IdentContext) {}

// EnterParenExpr is called when production parenExpr is entered.
func (s *BaseMonkeyListener) EnterParenExpr(ctx *ParenExprContext) {}

// ExitParenExpr is called when production parenExpr is exited.
func (s *BaseMonkeyListener) ExitParenExpr(ctx *ParenExprContext) {}

// EnterIndexExpr is called when production indexExpr is entered.
func (s *BaseMonkeyListener) EnterIndexExpr(ctx *IndexExprContext) {}

// ExitIndexExpr is called when production indexExpr is exited.
func (s *BaseMonkeyListener) ExitIndexExpr(ctx *IndexExprContext) {}

// EnterStrLit is called when production strLit is entered.
func (s *BaseMonkeyListener) EnterStrLit(ctx *StrLitContext) {}

// ExitStrLit is called when production strLit is exited.
func (s *BaseMonkeyListener) ExitStrLit(ctx *StrLitContext) {}

// EnterBoolLit is called when production boolLit is entered.
func (s *BaseMonkeyListener) EnterBoolLit(ctx *BoolLitContext) {}

// ExitBoolLit is called when production boolLit is exited.
func (s *BaseMonkeyListener) ExitBoolLit(ctx *BoolLitContext) {}

// EnterArrLit is called when production arrLit is entered.
func (s *BaseMonkeyListener) EnterArrLit(ctx *ArrLitContext) {}

// ExitArrLit is called when production arrLit is exited.
func (s *BaseMonkeyListener) ExitArrLit(ctx *ArrLitContext) {}

// EnterLtGtExpr is called when production ltGtExpr is entered.
func (s *BaseMonkeyListener) EnterLtGtExpr(ctx *LtGtExprContext) {}

// ExitLtGtExpr is called when production ltGtExpr is exited.
func (s *BaseMonkeyListener) ExitLtGtExpr(ctx *LtGtExprContext) {}

// EnterIntLit is called when production intLit is entered.
func (s *BaseMonkeyListener) EnterIntLit(ctx *IntLitContext) {}

// ExitIntLit is called when production intLit is exited.
func (s *BaseMonkeyListener) ExitIntLit(ctx *IntLitContext) {}

// EnterIfExpr is called when production ifExpr is entered.
func (s *BaseMonkeyListener) EnterIfExpr(ctx *IfExprContext) {}

// ExitIfExpr is called when production ifExpr is exited.
func (s *BaseMonkeyListener) ExitIfExpr(ctx *IfExprContext) {}

// EnterEqNeExpr is called when production eqNeExpr is entered.
func (s *BaseMonkeyListener) EnterEqNeExpr(ctx *EqNeExprContext) {}

// ExitEqNeExpr is called when production eqNeExpr is exited.
func (s *BaseMonkeyListener) ExitEqNeExpr(ctx *EqNeExprContext) {}

// EnterHashLit is called when production hashLit is entered.
func (s *BaseMonkeyListener) EnterHashLit(ctx *HashLitContext) {}

// ExitHashLit is called when production hashLit is exited.
func (s *BaseMonkeyListener) ExitHashLit(ctx *HashLitContext) {}

// EnterCallExpr is called when production callExpr is entered.
func (s *BaseMonkeyListener) EnterCallExpr(ctx *CallExprContext) {}

// ExitCallExpr is called when production callExpr is exited.
func (s *BaseMonkeyListener) ExitCallExpr(ctx *CallExprContext) {}

// EnterMulDivExpr is called when production mulDivExpr is entered.
func (s *BaseMonkeyListener) EnterMulDivExpr(ctx *MulDivExprContext) {}

// ExitMulDivExpr is called when production mulDivExpr is exited.
func (s *BaseMonkeyListener) ExitMulDivExpr(ctx *MulDivExprContext) {}

// EnterUnOpExpr is called when production unOpExpr is entered.
func (s *BaseMonkeyListener) EnterUnOpExpr(ctx *UnOpExprContext) {}

// ExitUnOpExpr is called when production unOpExpr is exited.
func (s *BaseMonkeyListener) ExitUnOpExpr(ctx *UnOpExprContext) {}

// EnterFnLit is called when production fnLit is entered.
func (s *BaseMonkeyListener) EnterFnLit(ctx *FnLitContext) {}

// ExitFnLit is called when production fnLit is exited.
func (s *BaseMonkeyListener) ExitFnLit(ctx *FnLitContext) {}

// EnterExprs is called when production exprs is entered.
func (s *BaseMonkeyListener) EnterExprs(ctx *ExprsContext) {}

// ExitExprs is called when production exprs is exited.
func (s *BaseMonkeyListener) ExitExprs(ctx *ExprsContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseMonkeyListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseMonkeyListener) ExitBlock(ctx *BlockContext) {}

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
