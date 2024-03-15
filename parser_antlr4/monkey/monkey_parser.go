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
		"", "'{'", "'}'", "'let'", "'='", "';'", "'return'", "'['", "']'", "'('",
		"')'", "'-'", "'!'", "'*'", "'/'", "'+'", "'<'", "'>'", "'=='", "'!='",
		"'if'", "'else'", "'fn'", "'true'", "'false'", "','", "':'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "IDENT", "INT", "STRING", "WS",
	}
	staticData.RuleNames = []string{
		"prog", "stat", "block", "letStat", "retStat", "exprStat", "expr", "exprs",
		"params", "pairs", "pair",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 30, 157, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 1, 0, 5, 0, 24, 8, 0, 10, 0, 12, 0, 27, 9, 0, 1, 1, 1, 1, 1, 1, 3,
		1, 32, 8, 1, 1, 2, 1, 2, 5, 2, 36, 8, 2, 10, 2, 12, 2, 39, 9, 2, 1, 2,
		1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 48, 8, 3, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1,
		6, 1, 6, 3, 6, 67, 8, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 3, 6, 91, 8, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1,
		6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1,
		6, 1, 6, 1, 6, 5, 6, 115, 8, 6, 10, 6, 12, 6, 118, 9, 6, 1, 7, 1, 7, 1,
		7, 5, 7, 123, 8, 7, 10, 7, 12, 7, 126, 9, 7, 1, 7, 3, 7, 129, 8, 7, 1,
		8, 1, 8, 1, 8, 5, 8, 134, 8, 8, 10, 8, 12, 8, 137, 9, 8, 1, 8, 3, 8, 140,
		8, 8, 1, 9, 1, 9, 1, 9, 5, 9, 145, 8, 9, 10, 9, 12, 9, 148, 9, 9, 1, 9,
		3, 9, 151, 8, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 0, 1, 12, 11, 0, 2,
		4, 6, 8, 10, 12, 14, 16, 18, 20, 0, 6, 1, 0, 11, 12, 1, 0, 23, 24, 1, 0,
		13, 14, 2, 0, 11, 11, 15, 15, 1, 0, 16, 17, 1, 0, 18, 19, 172, 0, 25, 1,
		0, 0, 0, 2, 31, 1, 0, 0, 0, 4, 33, 1, 0, 0, 0, 6, 42, 1, 0, 0, 0, 8, 49,
		1, 0, 0, 0, 10, 53, 1, 0, 0, 0, 12, 90, 1, 0, 0, 0, 14, 128, 1, 0, 0, 0,
		16, 139, 1, 0, 0, 0, 18, 150, 1, 0, 0, 0, 20, 152, 1, 0, 0, 0, 22, 24,
		3, 2, 1, 0, 23, 22, 1, 0, 0, 0, 24, 27, 1, 0, 0, 0, 25, 23, 1, 0, 0, 0,
		25, 26, 1, 0, 0, 0, 26, 1, 1, 0, 0, 0, 27, 25, 1, 0, 0, 0, 28, 32, 3, 6,
		3, 0, 29, 32, 3, 8, 4, 0, 30, 32, 3, 10, 5, 0, 31, 28, 1, 0, 0, 0, 31,
		29, 1, 0, 0, 0, 31, 30, 1, 0, 0, 0, 32, 3, 1, 0, 0, 0, 33, 37, 5, 1, 0,
		0, 34, 36, 3, 2, 1, 0, 35, 34, 1, 0, 0, 0, 36, 39, 1, 0, 0, 0, 37, 35,
		1, 0, 0, 0, 37, 38, 1, 0, 0, 0, 38, 40, 1, 0, 0, 0, 39, 37, 1, 0, 0, 0,
		40, 41, 5, 2, 0, 0, 41, 5, 1, 0, 0, 0, 42, 43, 5, 3, 0, 0, 43, 44, 5, 27,
		0, 0, 44, 45, 5, 4, 0, 0, 45, 47, 3, 12, 6, 0, 46, 48, 5, 5, 0, 0, 47,
		46, 1, 0, 0, 0, 47, 48, 1, 0, 0, 0, 48, 7, 1, 0, 0, 0, 49, 50, 5, 6, 0,
		0, 50, 51, 3, 12, 6, 0, 51, 52, 5, 5, 0, 0, 52, 9, 1, 0, 0, 0, 53, 54,
		3, 12, 6, 0, 54, 55, 5, 5, 0, 0, 55, 11, 1, 0, 0, 0, 56, 57, 6, 6, -1,
		0, 57, 58, 7, 0, 0, 0, 58, 91, 3, 12, 6, 14, 59, 60, 5, 20, 0, 0, 60, 61,
		5, 9, 0, 0, 61, 62, 3, 12, 6, 0, 62, 63, 5, 10, 0, 0, 63, 66, 3, 4, 2,
		0, 64, 65, 5, 21, 0, 0, 65, 67, 3, 4, 2, 0, 66, 64, 1, 0, 0, 0, 66, 67,
		1, 0, 0, 0, 67, 91, 1, 0, 0, 0, 68, 69, 5, 22, 0, 0, 69, 70, 5, 9, 0, 0,
		70, 71, 3, 16, 8, 0, 71, 72, 5, 10, 0, 0, 72, 73, 3, 4, 2, 0, 73, 91, 1,
		0, 0, 0, 74, 75, 5, 7, 0, 0, 75, 76, 3, 14, 7, 0, 76, 77, 5, 8, 0, 0, 77,
		91, 1, 0, 0, 0, 78, 79, 5, 1, 0, 0, 79, 80, 3, 18, 9, 0, 80, 81, 5, 2,
		0, 0, 81, 91, 1, 0, 0, 0, 82, 91, 5, 27, 0, 0, 83, 91, 5, 28, 0, 0, 84,
		91, 5, 29, 0, 0, 85, 91, 7, 1, 0, 0, 86, 87, 5, 9, 0, 0, 87, 88, 3, 12,
		6, 0, 88, 89, 5, 10, 0, 0, 89, 91, 1, 0, 0, 0, 90, 56, 1, 0, 0, 0, 90,
		59, 1, 0, 0, 0, 90, 68, 1, 0, 0, 0, 90, 74, 1, 0, 0, 0, 90, 78, 1, 0, 0,
		0, 90, 82, 1, 0, 0, 0, 90, 83, 1, 0, 0, 0, 90, 84, 1, 0, 0, 0, 90, 85,
		1, 0, 0, 0, 90, 86, 1, 0, 0, 0, 91, 116, 1, 0, 0, 0, 92, 93, 10, 13, 0,
		0, 93, 94, 7, 2, 0, 0, 94, 115, 3, 12, 6, 14, 95, 96, 10, 12, 0, 0, 96,
		97, 7, 3, 0, 0, 97, 115, 3, 12, 6, 13, 98, 99, 10, 11, 0, 0, 99, 100, 7,
		4, 0, 0, 100, 115, 3, 12, 6, 12, 101, 102, 10, 10, 0, 0, 102, 103, 7, 5,
		0, 0, 103, 115, 3, 12, 6, 11, 104, 105, 10, 16, 0, 0, 105, 106, 5, 7, 0,
		0, 106, 107, 3, 12, 6, 0, 107, 108, 5, 8, 0, 0, 108, 115, 1, 0, 0, 0, 109,
		110, 10, 15, 0, 0, 110, 111, 5, 9, 0, 0, 111, 112, 3, 14, 7, 0, 112, 113,
		5, 10, 0, 0, 113, 115, 1, 0, 0, 0, 114, 92, 1, 0, 0, 0, 114, 95, 1, 0,
		0, 0, 114, 98, 1, 0, 0, 0, 114, 101, 1, 0, 0, 0, 114, 104, 1, 0, 0, 0,
		114, 109, 1, 0, 0, 0, 115, 118, 1, 0, 0, 0, 116, 114, 1, 0, 0, 0, 116,
		117, 1, 0, 0, 0, 117, 13, 1, 0, 0, 0, 118, 116, 1, 0, 0, 0, 119, 124, 3,
		12, 6, 0, 120, 121, 5, 25, 0, 0, 121, 123, 3, 12, 6, 0, 122, 120, 1, 0,
		0, 0, 123, 126, 1, 0, 0, 0, 124, 122, 1, 0, 0, 0, 124, 125, 1, 0, 0, 0,
		125, 129, 1, 0, 0, 0, 126, 124, 1, 0, 0, 0, 127, 129, 1, 0, 0, 0, 128,
		119, 1, 0, 0, 0, 128, 127, 1, 0, 0, 0, 129, 15, 1, 0, 0, 0, 130, 135, 5,
		27, 0, 0, 131, 132, 5, 25, 0, 0, 132, 134, 5, 27, 0, 0, 133, 131, 1, 0,
		0, 0, 134, 137, 1, 0, 0, 0, 135, 133, 1, 0, 0, 0, 135, 136, 1, 0, 0, 0,
		136, 140, 1, 0, 0, 0, 137, 135, 1, 0, 0, 0, 138, 140, 1, 0, 0, 0, 139,
		130, 1, 0, 0, 0, 139, 138, 1, 0, 0, 0, 140, 17, 1, 0, 0, 0, 141, 146, 3,
		20, 10, 0, 142, 143, 5, 25, 0, 0, 143, 145, 3, 20, 10, 0, 144, 142, 1,
		0, 0, 0, 145, 148, 1, 0, 0, 0, 146, 144, 1, 0, 0, 0, 146, 147, 1, 0, 0,
		0, 147, 151, 1, 0, 0, 0, 148, 146, 1, 0, 0, 0, 149, 151, 1, 0, 0, 0, 150,
		141, 1, 0, 0, 0, 150, 149, 1, 0, 0, 0, 151, 19, 1, 0, 0, 0, 152, 153, 3,
		12, 6, 0, 153, 154, 5, 26, 0, 0, 154, 155, 3, 12, 6, 0, 155, 21, 1, 0,
		0, 0, 14, 25, 31, 37, 47, 66, 90, 114, 116, 124, 128, 135, 139, 146, 150,
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
	MonkeyParserRULE_prog     = 0
	MonkeyParserRULE_stat     = 1
	MonkeyParserRULE_block    = 2
	MonkeyParserRULE_letStat  = 3
	MonkeyParserRULE_retStat  = 4
	MonkeyParserRULE_exprStat = 5
	MonkeyParserRULE_expr     = 6
	MonkeyParserRULE_exprs    = 7
	MonkeyParserRULE_params   = 8
	MonkeyParserRULE_pairs    = 9
	MonkeyParserRULE_pair     = 10
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
	p.SetState(25)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&969939658) != 0 {
		{
			p.SetState(22)
			p.Stat()
		}

		p.SetState(27)
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

	// Getter signatures
	LetStat() ILetStatContext
	RetStat() IRetStatContext
	ExprStat() IExprStatContext

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

func (s *StatContext) LetStat() ILetStatContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILetStatContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILetStatContext)
}

func (s *StatContext) RetStat() IRetStatContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRetStatContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRetStatContext)
}

func (s *StatContext) ExprStat() IExprStatContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprStatContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprStatContext)
}

func (s *StatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyListener); ok {
		listenerT.EnterStat(s)
	}
}

func (s *StatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyListener); ok {
		listenerT.ExitStat(s)
	}
}

func (p *MonkeyParser) Stat() (localctx IStatContext) {
	localctx = NewStatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, MonkeyParserRULE_stat)
	p.SetState(31)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MonkeyParserT__2:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(28)
			p.LetStat()
		}

	case MonkeyParserT__5:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(29)
			p.RetStat()
		}

	case MonkeyParserT__0, MonkeyParserT__6, MonkeyParserT__8, MonkeyParserT__10, MonkeyParserT__11, MonkeyParserT__19, MonkeyParserT__21, MonkeyParserT__22, MonkeyParserT__23, MonkeyParserIDENT, MonkeyParserINT, MonkeyParserSTRING:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(30)
			p.ExprStat()
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
	p.EnterRule(localctx, 4, MonkeyParserRULE_block)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(33)
		p.Match(MonkeyParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(37)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&969939658) != 0 {
		{
			p.SetState(34)
			p.Stat()
		}

		p.SetState(39)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(40)
		p.Match(MonkeyParserT__1)
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

// ILetStatContext is an interface to support dynamic dispatch.
type ILetStatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENT() antlr.TerminalNode
	Expr() IExprContext

	// IsLetStatContext differentiates from other interfaces.
	IsLetStatContext()
}

type LetStatContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLetStatContext() *LetStatContext {
	var p = new(LetStatContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MonkeyParserRULE_letStat
	return p
}

func InitEmptyLetStatContext(p *LetStatContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MonkeyParserRULE_letStat
}

func (*LetStatContext) IsLetStatContext() {}

func NewLetStatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LetStatContext {
	var p = new(LetStatContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MonkeyParserRULE_letStat

	return p
}

func (s *LetStatContext) GetParser() antlr.Parser { return s.parser }

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

func (s *LetStatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LetStatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
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

func (p *MonkeyParser) LetStat() (localctx ILetStatContext) {
	localctx = NewLetStatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, MonkeyParserRULE_letStat)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(42)
		p.Match(MonkeyParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(43)
		p.Match(MonkeyParserIDENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(44)
		p.Match(MonkeyParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(45)
		p.expr(0)
	}
	p.SetState(47)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == MonkeyParserT__4 {
		{
			p.SetState(46)
			p.Match(MonkeyParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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

// IRetStatContext is an interface to support dynamic dispatch.
type IRetStatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expr() IExprContext

	// IsRetStatContext differentiates from other interfaces.
	IsRetStatContext()
}

type RetStatContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRetStatContext() *RetStatContext {
	var p = new(RetStatContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MonkeyParserRULE_retStat
	return p
}

func InitEmptyRetStatContext(p *RetStatContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MonkeyParserRULE_retStat
}

func (*RetStatContext) IsRetStatContext() {}

func NewRetStatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RetStatContext {
	var p = new(RetStatContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MonkeyParserRULE_retStat

	return p
}

func (s *RetStatContext) GetParser() antlr.Parser { return s.parser }

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

func (s *RetStatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RetStatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
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

func (p *MonkeyParser) RetStat() (localctx IRetStatContext) {
	localctx = NewRetStatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, MonkeyParserRULE_retStat)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(49)
		p.Match(MonkeyParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(50)
		p.expr(0)
	}
	{
		p.SetState(51)
		p.Match(MonkeyParserT__4)
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

// IExprStatContext is an interface to support dynamic dispatch.
type IExprStatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expr() IExprContext

	// IsExprStatContext differentiates from other interfaces.
	IsExprStatContext()
}

type ExprStatContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprStatContext() *ExprStatContext {
	var p = new(ExprStatContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MonkeyParserRULE_exprStat
	return p
}

func InitEmptyExprStatContext(p *ExprStatContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MonkeyParserRULE_exprStat
}

func (*ExprStatContext) IsExprStatContext() {}

func NewExprStatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprStatContext {
	var p = new(ExprStatContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MonkeyParserRULE_exprStat

	return p
}

func (s *ExprStatContext) GetParser() antlr.Parser { return s.parser }

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

func (s *ExprStatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprStatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
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

func (p *MonkeyParser) ExprStat() (localctx IExprStatContext) {
	localctx = NewExprStatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, MonkeyParserRULE_exprStat)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(53)
		p.expr(0)
	}
	{
		p.SetState(54)
		p.Match(MonkeyParserT__4)
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

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// Getter signatures
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	AllBlock() []IBlockContext
	Block(i int) IBlockContext
	Params() IParamsContext
	Exprs() IExprsContext
	Pairs() IPairsContext
	IDENT() antlr.TerminalNode
	INT() antlr.TerminalNode
	STRING() antlr.TerminalNode

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
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

func (s *ExprContext) GetOp() antlr.Token { return s.op }

func (s *ExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *ExprContext) AllExpr() []IExprContext {
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

func (s *ExprContext) Expr(i int) IExprContext {
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

func (s *ExprContext) AllBlock() []IBlockContext {
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

func (s *ExprContext) Block(i int) IBlockContext {
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

func (s *ExprContext) Params() IParamsContext {
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

func (s *ExprContext) Exprs() IExprsContext {
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

func (s *ExprContext) Pairs() IPairsContext {
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

func (s *ExprContext) IDENT() antlr.TerminalNode {
	return s.GetToken(MonkeyParserIDENT, 0)
}

func (s *ExprContext) INT() antlr.TerminalNode {
	return s.GetToken(MonkeyParserINT, 0)
}

func (s *ExprContext) STRING() antlr.TerminalNode {
	return s.GetToken(MonkeyParserSTRING, 0)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyListener); ok {
		listenerT.EnterExpr(s)
	}
}

func (s *ExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonkeyListener); ok {
		listenerT.ExitExpr(s)
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
	_startState := 12
	p.EnterRecursionRule(localctx, 12, MonkeyParserRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(90)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MonkeyParserT__10, MonkeyParserT__11:
		{
			p.SetState(57)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*ExprContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == MonkeyParserT__10 || _la == MonkeyParserT__11) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*ExprContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(58)
			p.expr(14)
		}

	case MonkeyParserT__19:
		{
			p.SetState(59)
			p.Match(MonkeyParserT__19)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(60)
			p.Match(MonkeyParserT__8)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(61)
			p.expr(0)
		}
		{
			p.SetState(62)
			p.Match(MonkeyParserT__9)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(63)
			p.Block()
		}
		p.SetState(66)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(64)
				p.Match(MonkeyParserT__20)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(65)
				p.Block()
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case MonkeyParserT__21:
		{
			p.SetState(68)
			p.Match(MonkeyParserT__21)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(69)
			p.Match(MonkeyParserT__8)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(70)
			p.Params()
		}
		{
			p.SetState(71)
			p.Match(MonkeyParserT__9)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(72)
			p.Block()
		}

	case MonkeyParserT__6:
		{
			p.SetState(74)
			p.Match(MonkeyParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(75)
			p.Exprs()
		}
		{
			p.SetState(76)
			p.Match(MonkeyParserT__7)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MonkeyParserT__0:
		{
			p.SetState(78)
			p.Match(MonkeyParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(79)
			p.Pairs()
		}
		{
			p.SetState(80)
			p.Match(MonkeyParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MonkeyParserIDENT:
		{
			p.SetState(82)
			p.Match(MonkeyParserIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MonkeyParserINT:
		{
			p.SetState(83)
			p.Match(MonkeyParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MonkeyParserSTRING:
		{
			p.SetState(84)
			p.Match(MonkeyParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MonkeyParserT__22, MonkeyParserT__23:
		{
			p.SetState(85)
			_la = p.GetTokenStream().LA(1)

			if !(_la == MonkeyParserT__22 || _la == MonkeyParserT__23) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case MonkeyParserT__8:
		{
			p.SetState(86)
			p.Match(MonkeyParserT__8)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(87)
			p.expr(0)
		}
		{
			p.SetState(88)
			p.Match(MonkeyParserT__9)
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
	p.SetState(116)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(114)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, MonkeyParserRULE_expr)
				p.SetState(92)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
					goto errorExit
				}
				{
					p.SetState(93)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == MonkeyParserT__12 || _la == MonkeyParserT__13) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(94)
					p.expr(14)
				}

			case 2:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, MonkeyParserRULE_expr)
				p.SetState(95)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
					goto errorExit
				}
				{
					p.SetState(96)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == MonkeyParserT__10 || _la == MonkeyParserT__14) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(97)
					p.expr(13)
				}

			case 3:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, MonkeyParserRULE_expr)
				p.SetState(98)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
					goto errorExit
				}
				{
					p.SetState(99)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == MonkeyParserT__15 || _la == MonkeyParserT__16) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(100)
					p.expr(12)
				}

			case 4:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, MonkeyParserRULE_expr)
				p.SetState(101)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(102)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == MonkeyParserT__17 || _la == MonkeyParserT__18) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(103)
					p.expr(11)
				}

			case 5:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, MonkeyParserRULE_expr)
				p.SetState(104)

				if !(p.Precpred(p.GetParserRuleContext(), 16)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 16)", ""))
					goto errorExit
				}
				{
					p.SetState(105)
					p.Match(MonkeyParserT__6)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(106)
					p.expr(0)
				}
				{
					p.SetState(107)
					p.Match(MonkeyParserT__7)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case 6:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, MonkeyParserRULE_expr)
				p.SetState(109)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
					goto errorExit
				}
				{
					p.SetState(110)
					p.Match(MonkeyParserT__8)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(111)
					p.Exprs()
				}
				{
					p.SetState(112)
					p.Match(MonkeyParserT__9)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(118)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 14, MonkeyParserRULE_exprs)
	var _la int

	p.SetState(128)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MonkeyParserT__0, MonkeyParserT__6, MonkeyParserT__8, MonkeyParserT__10, MonkeyParserT__11, MonkeyParserT__19, MonkeyParserT__21, MonkeyParserT__22, MonkeyParserT__23, MonkeyParserIDENT, MonkeyParserINT, MonkeyParserSTRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(119)
			p.expr(0)
		}
		p.SetState(124)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == MonkeyParserT__24 {
			{
				p.SetState(120)
				p.Match(MonkeyParserT__24)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(121)
				p.expr(0)
			}

			p.SetState(126)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case MonkeyParserT__7, MonkeyParserT__9:
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
	p.EnterRule(localctx, 16, MonkeyParserRULE_params)
	var _la int

	p.SetState(139)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MonkeyParserIDENT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(130)
			p.Match(MonkeyParserIDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(135)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == MonkeyParserT__24 {
			{
				p.SetState(131)
				p.Match(MonkeyParserT__24)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(132)
				p.Match(MonkeyParserIDENT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(137)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case MonkeyParserT__9:
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
	p.EnterRule(localctx, 18, MonkeyParserRULE_pairs)
	var _la int

	p.SetState(150)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MonkeyParserT__0, MonkeyParserT__6, MonkeyParserT__8, MonkeyParserT__10, MonkeyParserT__11, MonkeyParserT__19, MonkeyParserT__21, MonkeyParserT__22, MonkeyParserT__23, MonkeyParserIDENT, MonkeyParserINT, MonkeyParserSTRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(141)
			p.Pair()
		}
		p.SetState(146)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == MonkeyParserT__24 {
			{
				p.SetState(142)
				p.Match(MonkeyParserT__24)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(143)
				p.Pair()
			}

			p.SetState(148)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case MonkeyParserT__1:
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
	p.EnterRule(localctx, 20, MonkeyParserRULE_pair)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(152)
		p.expr(0)
	}
	{
		p.SetState(153)
		p.Match(MonkeyParserT__25)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(154)
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
	case 6:
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
