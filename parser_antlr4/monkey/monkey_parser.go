// Code generated from Monkey.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Monkey

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type MonkeyParser struct {
	*antlr.BaseParser
}

var MonkeyParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func monkeyParserInit() {
	staticData := &MonkeyParserStaticData
	staticData.LiteralNames = []string{
		"", "'let'", "'='", "';'", "'return'", "'['", "']'", "'('", "')'", "'-'",
		"'!'", "'*'", "'/'", "'+'", "'<'", "'>'", "'=='", "'!='", "'if'", "'else'",
		"'fn'", "'{'", "'}'", "'true'", "'false'", "','", "':'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "IDENT", "INT", "STRING", "WS",
	}
	staticData.RuleNames = []string{
		"prog", "stat", "expr", "exprs", "block", "params", "pairs", "pair",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 30, 150, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 1, 0, 5, 0, 18, 8, 0, 10, 0, 12,
		0, 21, 9, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 28, 8, 1, 1, 1, 1, 1,
		1, 1, 3, 1, 33, 8, 1, 1, 1, 1, 1, 3, 1, 37, 8, 1, 3, 1, 39, 8, 1, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 51, 8, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 75, 8, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 5, 2, 99, 8,
		2, 10, 2, 12, 2, 102, 9, 2, 1, 3, 1, 3, 1, 3, 5, 3, 107, 8, 3, 10, 3, 12,
		3, 110, 9, 3, 1, 3, 3, 3, 113, 8, 3, 1, 4, 1, 4, 5, 4, 117, 8, 4, 10, 4,
		12, 4, 120, 9, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 5, 5, 127, 8, 5, 10, 5,
		12, 5, 130, 9, 5, 1, 5, 3, 5, 133, 8, 5, 1, 6, 1, 6, 1, 6, 5, 6, 138, 8,
		6, 10, 6, 12, 6, 141, 9, 6, 1, 6, 3, 6, 144, 8, 6, 1, 7, 1, 7, 1, 7, 1,
		7, 1, 7, 0, 1, 4, 8, 0, 2, 4, 6, 8, 10, 12, 14, 0, 6, 1, 0, 9, 10, 1, 0,
		23, 24, 1, 0, 11, 12, 2, 0, 9, 9, 13, 13, 1, 0, 14, 15, 1, 0, 16, 17, 170,
		0, 19, 1, 0, 0, 0, 2, 38, 1, 0, 0, 0, 4, 74, 1, 0, 0, 0, 6, 112, 1, 0,
		0, 0, 8, 114, 1, 0, 0, 0, 10, 132, 1, 0, 0, 0, 12, 143, 1, 0, 0, 0, 14,
		145, 1, 0, 0, 0, 16, 18, 3, 2, 1, 0, 17, 16, 1, 0, 0, 0, 18, 21, 1, 0,
		0, 0, 19, 17, 1, 0, 0, 0, 19, 20, 1, 0, 0, 0, 20, 1, 1, 0, 0, 0, 21, 19,
		1, 0, 0, 0, 22, 23, 5, 1, 0, 0, 23, 24, 5, 27, 0, 0, 24, 25, 5, 2, 0, 0,
		25, 27, 3, 4, 2, 0, 26, 28, 5, 3, 0, 0, 27, 26, 1, 0, 0, 0, 27, 28, 1,
		0, 0, 0, 28, 39, 1, 0, 0, 0, 29, 30, 5, 4, 0, 0, 30, 32, 3, 4, 2, 0, 31,
		33, 5, 3, 0, 0, 32, 31, 1, 0, 0, 0, 32, 33, 1, 0, 0, 0, 33, 39, 1, 0, 0,
		0, 34, 36, 3, 4, 2, 0, 35, 37, 5, 3, 0, 0, 36, 35, 1, 0, 0, 0, 36, 37,
		1, 0, 0, 0, 37, 39, 1, 0, 0, 0, 38, 22, 1, 0, 0, 0, 38, 29, 1, 0, 0, 0,
		38, 34, 1, 0, 0, 0, 39, 3, 1, 0, 0, 0, 40, 41, 6, 2, -1, 0, 41, 42, 7,
		0, 0, 0, 42, 75, 3, 4, 2, 14, 43, 44, 5, 18, 0, 0, 44, 45, 5, 7, 0, 0,
		45, 46, 3, 4, 2, 0, 46, 47, 5, 8, 0, 0, 47, 50, 3, 8, 4, 0, 48, 49, 5,
		19, 0, 0, 49, 51, 3, 8, 4, 0, 50, 48, 1, 0, 0, 0, 50, 51, 1, 0, 0, 0, 51,
		75, 1, 0, 0, 0, 52, 53, 5, 20, 0, 0, 53, 54, 5, 7, 0, 0, 54, 55, 3, 10,
		5, 0, 55, 56, 5, 8, 0, 0, 56, 57, 3, 8, 4, 0, 57, 75, 1, 0, 0, 0, 58, 59,
		5, 5, 0, 0, 59, 60, 3, 6, 3, 0, 60, 61, 5, 6, 0, 0, 61, 75, 1, 0, 0, 0,
		62, 63, 5, 21, 0, 0, 63, 64, 3, 12, 6, 0, 64, 65, 5, 22, 0, 0, 65, 75,
		1, 0, 0, 0, 66, 75, 5, 27, 0, 0, 67, 75, 5, 28, 0, 0, 68, 75, 5, 29, 0,
		0, 69, 75, 7, 1, 0, 0, 70, 71, 5, 7, 0, 0, 71, 72, 3, 4, 2, 0, 72, 73,
		5, 8, 0, 0, 73, 75, 1, 0, 0, 0, 74, 40, 1, 0, 0, 0, 74, 43, 1, 0, 0, 0,
		74, 52, 1, 0, 0, 0, 74, 58, 1, 0, 0, 0, 74, 62, 1, 0, 0, 0, 74, 66, 1,
		0, 0, 0, 74, 67, 1, 0, 0, 0, 74, 68, 1, 0, 0, 0, 74, 69, 1, 0, 0, 0, 74,
		70, 1, 0, 0, 0, 75, 100, 1, 0, 0, 0, 76, 77, 10, 13, 0, 0, 77, 78, 7, 2,
		0, 0, 78, 99, 3, 4, 2, 14, 79, 80, 10, 12, 0, 0, 80, 81, 7, 3, 0, 0, 81,
		99, 3, 4, 2, 13, 82, 83, 10, 11, 0, 0, 83, 84, 7, 4, 0, 0, 84, 99, 3, 4,
		2, 12, 85, 86, 10, 10, 0, 0, 86, 87, 7, 5, 0, 0, 87, 99, 3, 4, 2, 11, 88,
		89, 10, 16, 0, 0, 89, 90, 5, 5, 0, 0, 90, 91, 3, 4, 2, 0, 91, 92, 5, 6,
		0, 0, 92, 99, 1, 0, 0, 0, 93, 94, 10, 15, 0, 0, 94, 95, 5, 7, 0, 0, 95,
		96, 3, 6, 3, 0, 96, 97, 5, 8, 0, 0, 97, 99, 1, 0, 0, 0, 98, 76, 1, 0, 0,
		0, 98, 79, 1, 0, 0, 0, 98, 82, 1, 0, 0, 0, 98, 85, 1, 0, 0, 0, 98, 88,
		1, 0, 0, 0, 98, 93, 1, 0, 0, 0, 99, 102, 1, 0, 0, 0, 100, 98, 1, 0, 0,
		0, 100, 101, 1, 0, 0, 0, 101, 5, 1, 0, 0, 0, 102, 100, 1, 0, 0, 0, 103,
		108, 3, 4, 2, 0, 104, 105, 5, 25, 0, 0, 105, 107, 3, 4, 2, 0, 106, 104,
		1, 0, 0, 0, 107, 110, 1, 0, 0, 0, 108, 106, 1, 0, 0, 0, 108, 109, 1, 0,
		0, 0, 109, 113, 1, 0, 0, 0, 110, 108, 1, 0, 0, 0, 111, 113, 1, 0, 0, 0,
		112, 103, 1, 0, 0, 0, 112, 111, 1, 0, 0, 0, 113, 7, 1, 0, 0, 0, 114, 118,
		5, 21, 0, 0, 115, 117, 3, 2, 1, 0, 116, 115, 1, 0, 0, 0, 117, 120, 1, 0,
		0, 0, 118, 116, 1, 0, 0, 0, 118, 119, 1, 0, 0, 0, 119, 121, 1, 0, 0, 0,
		120, 118, 1, 0, 0, 0, 121, 122, 5, 22, 0, 0, 122, 9, 1, 0, 0, 0, 123, 128,
		5, 27, 0, 0, 124, 125, 5, 25, 0, 0, 125, 127, 5, 27, 0, 0, 126, 124, 1,
		0, 0, 0, 127, 130, 1, 0, 0, 0, 128, 126, 1, 0, 0, 0, 128, 129, 1, 0, 0,
		0, 129, 133, 1, 0, 0, 0, 130, 128, 1, 0, 0, 0, 131, 133, 1, 0, 0, 0, 132,
		123, 1, 0, 0, 0, 132, 131, 1, 0, 0, 0, 133, 11, 1, 0, 0, 0, 134, 139, 3,
		14, 7, 0, 135, 136, 5, 25, 0, 0, 136, 138, 3, 14, 7, 0, 137, 135, 1, 0,
		0, 0, 138, 141, 1, 0, 0, 0, 139, 137, 1, 0, 0, 0, 139, 140, 1, 0, 0, 0,
		140, 144, 1, 0, 0, 0, 141, 139, 1, 0, 0, 0, 142, 144, 1, 0, 0, 0, 143,
		134, 1, 0, 0, 0, 143, 142, 1, 0, 0, 0, 144, 13, 1, 0, 0, 0, 145, 146, 3,
		4, 2, 0, 146, 147, 5, 26, 0, 0, 147, 148, 3, 4, 2, 0, 148, 15, 1, 0, 0,
		0, 16, 19, 27, 32, 36, 38, 50, 74, 98, 100, 108, 112, 118, 128, 132, 139,
		143,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// MonkeyParserInit initializes any static state used to implement MonkeyParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewMonkeyParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func MonkeyParserInit() {
	staticData := &MonkeyParserStaticData
	staticData.once.Do(monkeyParserInit)
}

// NewMonkeyParser produces a new parser instance for the optional input antlr.TokenStream.
func NewMonkeyParser(input antlr.TokenStream) *MonkeyParser {
	MonkeyParserInit()
	this := new(MonkeyParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &MonkeyParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Monkey.g4"

	return this
}

// MonkeyParser tokens.
const (
	MonkeyParserEOF    = antlr.TokenEOF
	MonkeyParserT__0   = 1
	MonkeyParserT__1   = 2
	MonkeyParserT__2   = 3
	MonkeyParserT__3   = 4
	MonkeyParserT__4   = 5
	MonkeyParserT__5   = 6
	MonkeyParserT__6   = 7
	MonkeyParserT__7   = 8
	MonkeyParserT__8   = 9
	MonkeyParserT__9   = 10
	MonkeyParserT__10  = 11
	MonkeyParserT__11  = 12
	MonkeyParserT__12  = 13
	MonkeyParserT__13  = 14
	MonkeyParserT__14  = 15
	MonkeyParserT__15  = 16
	MonkeyParserT__16  = 17
	MonkeyParserT__17  = 18
	MonkeyParserT__18  = 19
	MonkeyParserT__19  = 20
	MonkeyParserT__20  = 21
	MonkeyParserT__21  = 22
	MonkeyParserT__22  = 23
	MonkeyParserT__23  = 24
	MonkeyParserT__24  = 25
	MonkeyParserT__25  = 26
	MonkeyParserIDENT  = 27
	MonkeyParserINT    = 28
	MonkeyParserSTRING = 29
	MonkeyParserWS     = 30
)

// MonkeyParser rules.
const (
	MonkeyParserRULE_prog   = 0
	MonkeyParserRULE_stat   = 1
	MonkeyParserRULE_expr   = 2
	MonkeyParserRULE_exprs  = 3
	MonkeyParserRULE_block  = 4
	MonkeyParserRULE_params = 5
	MonkeyParserRULE_pairs  = 6
	MonkeyParserRULE_pair   = 7
)

// IProgContext is an interface to support dynamic dispatch.
type IProgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStat() []IStatContext
	Stat(i int) IStatContext

	// IsProgContext differentiates from other interfaces.
	IsProgContext()
}

type ProgContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgContext() *ProgContext {
	var p = new(ProgContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MonkeyParserRULE_prog
	return p
}

func InitEmptyProgContext(p *ProgContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MonkeyParserRULE_prog
}

func (*ProgContext) IsProgContext() {}

func NewProgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgContext {
	var p = new(ProgContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MonkeyParserRULE_prog

	return p
}

func (s *ProgContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgContext) AllStat() []IStatContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatContext); ok {
			len++
		}
	}

	tst := make([]IStatContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatContext); ok {
			tst[i] = t.(IStatContext)
			i++
		}
	}

	return tst
}

func (s *ProgContext) Stat(i int) IStatContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatContext)
}

func (s *ProgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyListener); ok {
		listenerT.EnterProg(s)
	}
}

func (s *ProgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyListener); ok {
		listenerT.ExitProg(s)
	}
}

func (p *MonkeyParser) Prog() (localctx IProgContext) {
	localctx = NewProgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, MonkeyParserRULE_prog)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(19)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&968099506) != 0 {
		{
			p.SetState(16)
			p.Stat()
		}

		p.SetState(21)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStatContext is an interface to support dynamic dispatch.
type IStatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsStatContext differentiates from other interfaces.
	IsStatContext()
}

type StatContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatContext() *StatContext {
	var p = new(StatContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MonkeyParserRULE_stat
	return p
}

func InitEmptyStatContext(p *StatContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MonkeyParserRULE_stat
}

func (*StatContext) IsStatContext() {}

func NewStatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatContext {
	var p = new(StatContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MonkeyParserRULE_stat

	return p
}

func (s *StatContext) GetParser() antlr.Parser { return s.parser }

func (s *StatContext) CopyAll(ctx *StatContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *StatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ExprStatContext struct {
	StatContext
}

func NewExprStatContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprStatContext {
	var p = new(ExprStatContext)

	InitEmptyStatContext(&p.StatContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatContext))

	return p
}

func (s *ExprStatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprStatContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprStatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyListener); ok {
		listenerT.EnterExprStat(s)
	}
}

func (s *ExprStatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyListener); ok {
		listenerT.ExitExprStat(s)
	}
}

type LetStatContext struct {
	StatContext
}

func NewLetStatContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LetStatContext {
	var p = new(LetStatContext)

	InitEmptyStatContext(&p.StatContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatContext))

	return p
}

func (s *LetStatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LetStatContext) IDENT() antlr.TerminalNode {
	return s.GetToken(MonkeyParserIDENT, 0)
}

func (s *LetStatContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *LetStatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyListener); ok {
		listenerT.EnterLetStat(s)
	}
}

func (s *LetStatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyListener); ok {
		listenerT.ExitLetStat(s)
	}
}

type RetStatContext struct {
	StatContext
}

func NewRetStatContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RetStatContext {
	var p = new(RetStatContext)

	InitEmptyStatContext(&p.StatContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatContext))

	return p
}

func (s *RetStatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RetStatContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *RetStatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyListener); ok {
		listenerT.EnterRetStat(s)
	}
}

func (s *RetStatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyListener); ok {
		listenerT.ExitRetStat(s)
	}
}

func (p *MonkeyParser) Stat() (localctx IStatContext) {
	localctx = NewStatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, MonkeyParserRULE_stat)
	var _la int

	p.SetState(38)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MonkeyParserT__0:
		localctx = NewLetStatContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(22)
			p.Match(MonkeyParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(23)
			p.Match(MonkeyParserIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(24)
			p.Match(MonkeyParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(25)
			p.expr(0)
		}
		p.SetState(27)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == MonkeyParserT__2 {
			{
				p.SetState(26)
				p.Match(MonkeyParserT__2)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case MonkeyParserT__3:
		localctx = NewRetStatContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(29)
			p.Match(MonkeyParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(30)
			p.expr(0)
		}
		p.SetState(32)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == MonkeyParserT__2 {
			{
				p.SetState(31)
				p.Match(MonkeyParserT__2)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case MonkeyParserT__4, MonkeyParserT__6, MonkeyParserT__8, MonkeyParserT__9, MonkeyParserT__17, MonkeyParserT__19, MonkeyParserT__20, MonkeyParserT__22, MonkeyParserT__23, MonkeyParserIDENT, MonkeyParserINT, MonkeyParserSTRING:
		localctx = NewExprStatContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(34)
			p.expr(0)
		}
		p.SetState(36)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == MonkeyParserT__2 {
			{
				p.SetState(35)
				p.Match(MonkeyParserT__2)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MonkeyParserRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MonkeyParserRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MonkeyParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) CopyAll(ctx *ExprContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type AddSubExprContext struct {
	ExprContext
	op antlr.Token
}

func NewAddSubExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AddSubExprContext {
	var p = new(AddSubExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *AddSubExprContext) GetOp() antlr.Token { return s.op }

func (s *AddSubExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *AddSubExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddSubExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *AddSubExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AddSubExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyListener); ok {
		listenerT.EnterAddSubExpr(s)
	}
}

func (s *AddSubExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyListener); ok {
		listenerT.ExitAddSubExpr(s)
	}
}

type IdentContext struct {
	ExprContext
}

func NewIdentContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdentContext {
	var p = new(IdentContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *IdentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentContext) IDENT() antlr.TerminalNode {
	return s.GetToken(MonkeyParserIDENT, 0)
}

func (s *IdentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyListener); ok {
		listenerT.EnterIdent(s)
	}
}

func (s *IdentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyListener); ok {
		listenerT.ExitIdent(s)
	}
}

type ParenExprContext struct {
	ExprContext
}

func NewParenExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParenExprContext {
	var p = new(ParenExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ParenExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenExprContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ParenExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyListener); ok {
		listenerT.EnterParenExpr(s)
	}
}

func (s *ParenExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyListener); ok {
		listenerT.ExitParenExpr(s)
	}
}

type IndexExprContext struct {
	ExprContext
}

func NewIndexExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IndexExprContext {
	var p = new(IndexExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *IndexExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *IndexExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *IndexExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyListener); ok {
		listenerT.EnterIndexExpr(s)
	}
}

func (s *IndexExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyListener); ok {
		listenerT.ExitIndexExpr(s)
	}
}

type StrLitContext struct {
	ExprContext
}

func NewStrLitContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StrLitContext {
	var p = new(StrLitContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *StrLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StrLitContext) STRING() antlr.TerminalNode {
	return s.GetToken(MonkeyParserSTRING, 0)
}

func (s *StrLitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyListener); ok {
		listenerT.EnterStrLit(s)
	}
}

func (s *StrLitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyListener); ok {
		listenerT.ExitStrLit(s)
	}
}

type BoolLitContext struct {
	ExprContext
}

func NewBoolLitContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BoolLitContext {
	var p = new(BoolLitContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *BoolLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolLitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyListener); ok {
		listenerT.EnterBoolLit(s)
	}
}

func (s *BoolLitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyListener); ok {
		listenerT.ExitBoolLit(s)
	}
}

type ArrLitContext struct {
	ExprContext
}

func NewArrLitContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArrLitContext {
	var p = new(ArrLitContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ArrLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrLitContext) Exprs() IExprsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprsContext)
}

func (s *ArrLitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyListener); ok {
		listenerT.EnterArrLit(s)
	}
}

func (s *ArrLitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyListener); ok {
		listenerT.ExitArrLit(s)
	}
}

type LtGtExprContext struct {
	ExprContext
	op antlr.Token
}

func NewLtGtExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LtGtExprContext {
	var p = new(LtGtExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *LtGtExprContext) GetOp() antlr.Token { return s.op }

func (s *LtGtExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *LtGtExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LtGtExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *LtGtExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *LtGtExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyListener); ok {
		listenerT.EnterLtGtExpr(s)
	}
}

func (s *LtGtExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyListener); ok {
		listenerT.ExitLtGtExpr(s)
	}
}

type IntLitContext struct {
	ExprContext
}

func NewIntLitContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IntLitContext {
	var p = new(IntLitContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *IntLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntLitContext) INT() antlr.TerminalNode {
	return s.GetToken(MonkeyParserINT, 0)
}

func (s *IntLitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyListener); ok {
		listenerT.EnterIntLit(s)
	}
}

func (s *IntLitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyListener); ok {
		listenerT.ExitIntLit(s)
	}
}

type IfExprContext struct {
	ExprContext
}

func NewIfExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IfExprContext {
	var p = new(IfExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *IfExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfExprContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *IfExprContext) AllBlock() []IBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBlockContext); ok {
			len++
		}
	}

	tst := make([]IBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBlockContext); ok {
			tst[i] = t.(IBlockContext)
			i++
		}
	}

	return tst
}

func (s *IfExprContext) Block(i int) IBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *IfExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyListener); ok {
		listenerT.EnterIfExpr(s)
	}
}

func (s *IfExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyListener); ok {
		listenerT.ExitIfExpr(s)
	}
}

type EqNeExprContext struct {
	ExprContext
	op antlr.Token
}

func NewEqNeExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EqNeExprContext {
	var p = new(EqNeExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *EqNeExprContext) GetOp() antlr.Token { return s.op }

func (s *EqNeExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *EqNeExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqNeExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *EqNeExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *EqNeExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyListener); ok {
		listenerT.EnterEqNeExpr(s)
	}
}

func (s *EqNeExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyListener); ok {
		listenerT.ExitEqNeExpr(s)
	}
}

type HashLitContext struct {
	ExprContext
}

func NewHashLitContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *HashLitContext {
	var p = new(HashLitContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *HashLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HashLitContext) Pairs() IPairsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPairsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPairsContext)
}

func (s *HashLitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyListener); ok {
		listenerT.EnterHashLit(s)
	}
}

func (s *HashLitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyListener); ok {
		listenerT.ExitHashLit(s)
	}
}

type CallExprContext struct {
	ExprContext
}

func NewCallExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CallExprContext {
	var p = new(CallExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *CallExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CallExprContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *CallExprContext) Exprs() IExprsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprsContext)
}

func (s *CallExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyListener); ok {
		listenerT.EnterCallExpr(s)
	}
}

func (s *CallExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyListener); ok {
		listenerT.ExitCallExpr(s)
	}
}

type MulDivExprContext struct {
	ExprContext
	op antlr.Token
}

func NewMulDivExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MulDivExprContext {
	var p = new(MulDivExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *MulDivExprContext) GetOp() antlr.Token { return s.op }

func (s *MulDivExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *MulDivExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MulDivExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *MulDivExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *MulDivExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyListener); ok {
		listenerT.EnterMulDivExpr(s)
	}
}

func (s *MulDivExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyListener); ok {
		listenerT.ExitMulDivExpr(s)
	}
}

type UnOpExprContext struct {
	ExprContext
	op antlr.Token
}

func NewUnOpExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnOpExprContext {
	var p = new(UnOpExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *UnOpExprContext) GetOp() antlr.Token { return s.op }

func (s *UnOpExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *UnOpExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnOpExprContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *UnOpExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyListener); ok {
		listenerT.EnterUnOpExpr(s)
	}
}

func (s *UnOpExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyListener); ok {
		listenerT.ExitUnOpExpr(s)
	}
}

type FnLitContext struct {
	ExprContext
}

func NewFnLitContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FnLitContext {
	var p = new(FnLitContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *FnLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FnLitContext) Params() IParamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamsContext)
}

func (s *FnLitContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *FnLitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyListener); ok {
		listenerT.EnterFnLit(s)
	}
}

func (s *FnLitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyListener); ok {
		listenerT.ExitFnLit(s)
	}
}

func (p *MonkeyParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *MonkeyParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 4
	p.EnterRecursionRule(localctx, 4, MonkeyParserRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(74)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MonkeyParserT__8, MonkeyParserT__9:
		localctx = NewUnOpExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(41)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*UnOpExprContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == MonkeyParserT__8 || _la == MonkeyParserT__9) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*UnOpExprContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(42)
			p.expr(14)
		}

	case MonkeyParserT__17:
		localctx = NewIfExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(43)
			p.Match(MonkeyParserT__17)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(44)
			p.Match(MonkeyParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(45)
			p.expr(0)
		}
		{
			p.SetState(46)
			p.Match(MonkeyParserT__7)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(47)
			p.Block()
		}
		p.SetState(50)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(48)
				p.Match(MonkeyParserT__18)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(49)
				p.Block()
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case MonkeyParserT__19:
		localctx = NewFnLitContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(52)
			p.Match(MonkeyParserT__19)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(53)
			p.Match(MonkeyParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(54)
			p.Params()
		}
		{
			p.SetState(55)
			p.Match(MonkeyParserT__7)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(56)
			p.Block()
		}

	case MonkeyParserT__4:
		localctx = NewArrLitContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(58)
			p.Match(MonkeyParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(59)
			p.Exprs()
		}
		{
			p.SetState(60)
			p.Match(MonkeyParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MonkeyParserT__20:
		localctx = NewHashLitContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(62)
			p.Match(MonkeyParserT__20)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(63)
			p.Pairs()
		}
		{
			p.SetState(64)
			p.Match(MonkeyParserT__21)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MonkeyParserIDENT:
		localctx = NewIdentContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(66)
			p.Match(MonkeyParserIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MonkeyParserINT:
		localctx = NewIntLitContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(67)
			p.Match(MonkeyParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MonkeyParserSTRING:
		localctx = NewStrLitContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(68)
			p.Match(MonkeyParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MonkeyParserT__22, MonkeyParserT__23:
		localctx = NewBoolLitContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(69)
			_la = p.GetTokenStream().LA(1)

			if !(_la == MonkeyParserT__22 || _la == MonkeyParserT__23) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case MonkeyParserT__6:
		localctx = NewParenExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(70)
			p.Match(MonkeyParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(71)
			p.expr(0)
		}
		{
			p.SetState(72)
			p.Match(MonkeyParserT__7)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(100)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(98)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMulDivExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MonkeyParserRULE_expr)
				p.SetState(76)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
					goto errorExit
				}
				{
					p.SetState(77)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*MulDivExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == MonkeyParserT__10 || _la == MonkeyParserT__11) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*MulDivExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(78)
					p.expr(14)
				}

			case 2:
				localctx = NewAddSubExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MonkeyParserRULE_expr)
				p.SetState(79)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
					goto errorExit
				}
				{
					p.SetState(80)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*AddSubExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == MonkeyParserT__8 || _la == MonkeyParserT__12) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*AddSubExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(81)
					p.expr(13)
				}

			case 3:
				localctx = NewLtGtExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MonkeyParserRULE_expr)
				p.SetState(82)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
					goto errorExit
				}
				{
					p.SetState(83)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*LtGtExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == MonkeyParserT__13 || _la == MonkeyParserT__14) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*LtGtExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(84)
					p.expr(12)
				}

			case 4:
				localctx = NewEqNeExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MonkeyParserRULE_expr)
				p.SetState(85)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(86)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*EqNeExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == MonkeyParserT__15 || _la == MonkeyParserT__16) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*EqNeExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(87)
					p.expr(11)
				}

			case 5:
				localctx = NewIndexExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MonkeyParserRULE_expr)
				p.SetState(88)

				if !(p.Precpred(p.GetParserRuleContext(), 16)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 16)", ""))
					goto errorExit
				}
				{
					p.SetState(89)
					p.Match(MonkeyParserT__4)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(90)
					p.expr(0)
				}
				{
					p.SetState(91)
					p.Match(MonkeyParserT__5)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case 6:
				localctx = NewCallExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, MonkeyParserRULE_expr)
				p.SetState(93)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
					goto errorExit
				}
				{
					p.SetState(94)
					p.Match(MonkeyParserT__6)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(95)
					p.Exprs()
				}
				{
					p.SetState(96)
					p.Match(MonkeyParserT__7)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(102)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExprsContext is an interface to support dynamic dispatch.
type IExprsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpr() []IExprContext
	Expr(i int) IExprContext

	// IsExprsContext differentiates from other interfaces.
	IsExprsContext()
}

type ExprsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprsContext() *ExprsContext {
	var p = new(ExprsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MonkeyParserRULE_exprs
	return p
}

func InitEmptyExprsContext(p *ExprsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MonkeyParserRULE_exprs
}

func (*ExprsContext) IsExprsContext() {}

func NewExprsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprsContext {
	var p = new(ExprsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MonkeyParserRULE_exprs

	return p
}

func (s *ExprsContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprsContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ExprsContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyListener); ok {
		listenerT.EnterExprs(s)
	}
}

func (s *ExprsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyListener); ok {
		listenerT.ExitExprs(s)
	}
}

func (p *MonkeyParser) Exprs() (localctx IExprsContext) {
	localctx = NewExprsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, MonkeyParserRULE_exprs)
	var _la int

	p.SetState(112)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MonkeyParserT__4, MonkeyParserT__6, MonkeyParserT__8, MonkeyParserT__9, MonkeyParserT__17, MonkeyParserT__19, MonkeyParserT__20, MonkeyParserT__22, MonkeyParserT__23, MonkeyParserIDENT, MonkeyParserINT, MonkeyParserSTRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(103)
			p.expr(0)
		}
		p.SetState(108)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == MonkeyParserT__24 {
			{
				p.SetState(104)
				p.Match(MonkeyParserT__24)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(105)
				p.expr(0)
			}

			p.SetState(110)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case MonkeyParserT__5, MonkeyParserT__7:
		p.EnterOuterAlt(localctx, 2)

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStat() []IStatContext
	Stat(i int) IStatContext

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MonkeyParserRULE_block
	return p
}

func InitEmptyBlockContext(p *BlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MonkeyParserRULE_block
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MonkeyParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) AllStat() []IStatContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatContext); ok {
			len++
		}
	}

	tst := make([]IStatContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatContext); ok {
			tst[i] = t.(IStatContext)
			i++
		}
	}

	return tst
}

func (s *BlockContext) Stat(i int) IStatContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (p *MonkeyParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, MonkeyParserRULE_block)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(114)
		p.Match(MonkeyParserT__20)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(118)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&968099506) != 0 {
		{
			p.SetState(115)
			p.Stat()
		}

		p.SetState(120)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(121)
		p.Match(MonkeyParserT__21)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParamsContext is an interface to support dynamic dispatch.
type IParamsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIDENT() []antlr.TerminalNode
	IDENT(i int) antlr.TerminalNode

	// IsParamsContext differentiates from other interfaces.
	IsParamsContext()
}

type ParamsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamsContext() *ParamsContext {
	var p = new(ParamsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MonkeyParserRULE_params
	return p
}

func InitEmptyParamsContext(p *ParamsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MonkeyParserRULE_params
}

func (*ParamsContext) IsParamsContext() {}

func NewParamsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamsContext {
	var p = new(ParamsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MonkeyParserRULE_params

	return p
}

func (s *ParamsContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamsContext) AllIDENT() []antlr.TerminalNode {
	return s.GetTokens(MonkeyParserIDENT)
}

func (s *ParamsContext) IDENT(i int) antlr.TerminalNode {
	return s.GetToken(MonkeyParserIDENT, i)
}

func (s *ParamsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyListener); ok {
		listenerT.EnterParams(s)
	}
}

func (s *ParamsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyListener); ok {
		listenerT.ExitParams(s)
	}
}

func (p *MonkeyParser) Params() (localctx IParamsContext) {
	localctx = NewParamsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, MonkeyParserRULE_params)
	var _la int

	p.SetState(132)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MonkeyParserIDENT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(123)
			p.Match(MonkeyParserIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(128)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == MonkeyParserT__24 {
			{
				p.SetState(124)
				p.Match(MonkeyParserT__24)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(125)
				p.Match(MonkeyParserIDENT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(130)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case MonkeyParserT__7:
		p.EnterOuterAlt(localctx, 2)

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPairsContext is an interface to support dynamic dispatch.
type IPairsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllPair() []IPairContext
	Pair(i int) IPairContext

	// IsPairsContext differentiates from other interfaces.
	IsPairsContext()
}

type PairsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPairsContext() *PairsContext {
	var p = new(PairsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MonkeyParserRULE_pairs
	return p
}

func InitEmptyPairsContext(p *PairsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MonkeyParserRULE_pairs
}

func (*PairsContext) IsPairsContext() {}

func NewPairsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PairsContext {
	var p = new(PairsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MonkeyParserRULE_pairs

	return p
}

func (s *PairsContext) GetParser() antlr.Parser { return s.parser }

func (s *PairsContext) AllPair() []IPairContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPairContext); ok {
			len++
		}
	}

	tst := make([]IPairContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPairContext); ok {
			tst[i] = t.(IPairContext)
			i++
		}
	}

	return tst
}

func (s *PairsContext) Pair(i int) IPairContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPairContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPairContext)
}

func (s *PairsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PairsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PairsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyListener); ok {
		listenerT.EnterPairs(s)
	}
}

func (s *PairsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyListener); ok {
		listenerT.ExitPairs(s)
	}
}

func (p *MonkeyParser) Pairs() (localctx IPairsContext) {
	localctx = NewPairsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, MonkeyParserRULE_pairs)
	var _la int

	p.SetState(143)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MonkeyParserT__4, MonkeyParserT__6, MonkeyParserT__8, MonkeyParserT__9, MonkeyParserT__17, MonkeyParserT__19, MonkeyParserT__20, MonkeyParserT__22, MonkeyParserT__23, MonkeyParserIDENT, MonkeyParserINT, MonkeyParserSTRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(134)
			p.Pair()
		}
		p.SetState(139)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == MonkeyParserT__24 {
			{
				p.SetState(135)
				p.Match(MonkeyParserT__24)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(136)
				p.Pair()
			}

			p.SetState(141)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case MonkeyParserT__21:
		p.EnterOuterAlt(localctx, 2)

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPairContext is an interface to support dynamic dispatch.
type IPairContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpr() []IExprContext
	Expr(i int) IExprContext

	// IsPairContext differentiates from other interfaces.
	IsPairContext()
}

type PairContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPairContext() *PairContext {
	var p = new(PairContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MonkeyParserRULE_pair
	return p
}

func InitEmptyPairContext(p *PairContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MonkeyParserRULE_pair
}

func (*PairContext) IsPairContext() {}

func NewPairContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PairContext {
	var p = new(PairContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MonkeyParserRULE_pair

	return p
}

func (s *PairContext) GetParser() antlr.Parser { return s.parser }

func (s *PairContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *PairContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *PairContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PairContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PairContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyListener); ok {
		listenerT.EnterPair(s)
	}
}

func (s *PairContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyListener); ok {
		listenerT.ExitPair(s)
	}
}

func (p *MonkeyParser) Pair() (localctx IPairContext) {
	localctx = NewPairContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, MonkeyParserRULE_pair)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(145)
		p.expr(0)
	}
	{
		p.SetState(146)
		p.Match(MonkeyParserT__25)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(147)
		p.expr(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *MonkeyParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 2:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *MonkeyParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 16)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 15)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
