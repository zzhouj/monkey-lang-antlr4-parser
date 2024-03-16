// Code generated from Monkey.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type MonkeyLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var MonkeyLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func monkeylexerLexerInit() {
	staticData := &MonkeyLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
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
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
		"T__17", "T__18", "T__19", "T__20", "T__21", "T__22", "T__23", "T__24",
		"T__25", "IDENT", "INT", "STRING", "WS", "LETER", "DIFIT",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 30, 176, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7,
		1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12,
		1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1,
		17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19,
		1, 20, 1, 20, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 23, 1,
		23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 25, 1, 25, 1, 26, 1, 26,
		1, 26, 5, 26, 142, 8, 26, 10, 26, 12, 26, 145, 9, 26, 1, 27, 1, 27, 1,
		27, 5, 27, 150, 8, 27, 10, 27, 12, 27, 153, 9, 27, 3, 27, 155, 8, 27, 1,
		28, 1, 28, 5, 28, 159, 8, 28, 10, 28, 12, 28, 162, 9, 28, 1, 28, 1, 28,
		1, 29, 4, 29, 167, 8, 29, 11, 29, 12, 29, 168, 1, 29, 1, 29, 1, 30, 1,
		30, 1, 31, 1, 31, 0, 0, 32, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7,
		15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33,
		17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51,
		26, 53, 27, 55, 28, 57, 29, 59, 30, 61, 0, 63, 0, 1, 0, 5, 1, 0, 49, 57,
		1, 0, 34, 34, 3, 0, 9, 10, 13, 13, 32, 32, 3, 0, 65, 90, 95, 95, 97, 122,
		1, 0, 48, 57, 179, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0,
		0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0,
		0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0,
		0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1,
		0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37,
		1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0,
		45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0,
		0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0,
		0, 1, 65, 1, 0, 0, 0, 3, 69, 1, 0, 0, 0, 5, 71, 1, 0, 0, 0, 7, 73, 1, 0,
		0, 0, 9, 80, 1, 0, 0, 0, 11, 82, 1, 0, 0, 0, 13, 84, 1, 0, 0, 0, 15, 86,
		1, 0, 0, 0, 17, 88, 1, 0, 0, 0, 19, 90, 1, 0, 0, 0, 21, 92, 1, 0, 0, 0,
		23, 94, 1, 0, 0, 0, 25, 96, 1, 0, 0, 0, 27, 98, 1, 0, 0, 0, 29, 100, 1,
		0, 0, 0, 31, 102, 1, 0, 0, 0, 33, 105, 1, 0, 0, 0, 35, 108, 1, 0, 0, 0,
		37, 111, 1, 0, 0, 0, 39, 116, 1, 0, 0, 0, 41, 119, 1, 0, 0, 0, 43, 121,
		1, 0, 0, 0, 45, 123, 1, 0, 0, 0, 47, 128, 1, 0, 0, 0, 49, 134, 1, 0, 0,
		0, 51, 136, 1, 0, 0, 0, 53, 138, 1, 0, 0, 0, 55, 154, 1, 0, 0, 0, 57, 156,
		1, 0, 0, 0, 59, 166, 1, 0, 0, 0, 61, 172, 1, 0, 0, 0, 63, 174, 1, 0, 0,
		0, 65, 66, 5, 108, 0, 0, 66, 67, 5, 101, 0, 0, 67, 68, 5, 116, 0, 0, 68,
		2, 1, 0, 0, 0, 69, 70, 5, 61, 0, 0, 70, 4, 1, 0, 0, 0, 71, 72, 5, 59, 0,
		0, 72, 6, 1, 0, 0, 0, 73, 74, 5, 114, 0, 0, 74, 75, 5, 101, 0, 0, 75, 76,
		5, 116, 0, 0, 76, 77, 5, 117, 0, 0, 77, 78, 5, 114, 0, 0, 78, 79, 5, 110,
		0, 0, 79, 8, 1, 0, 0, 0, 80, 81, 5, 91, 0, 0, 81, 10, 1, 0, 0, 0, 82, 83,
		5, 93, 0, 0, 83, 12, 1, 0, 0, 0, 84, 85, 5, 40, 0, 0, 85, 14, 1, 0, 0,
		0, 86, 87, 5, 41, 0, 0, 87, 16, 1, 0, 0, 0, 88, 89, 5, 45, 0, 0, 89, 18,
		1, 0, 0, 0, 90, 91, 5, 33, 0, 0, 91, 20, 1, 0, 0, 0, 92, 93, 5, 42, 0,
		0, 93, 22, 1, 0, 0, 0, 94, 95, 5, 47, 0, 0, 95, 24, 1, 0, 0, 0, 96, 97,
		5, 43, 0, 0, 97, 26, 1, 0, 0, 0, 98, 99, 5, 60, 0, 0, 99, 28, 1, 0, 0,
		0, 100, 101, 5, 62, 0, 0, 101, 30, 1, 0, 0, 0, 102, 103, 5, 61, 0, 0, 103,
		104, 5, 61, 0, 0, 104, 32, 1, 0, 0, 0, 105, 106, 5, 33, 0, 0, 106, 107,
		5, 61, 0, 0, 107, 34, 1, 0, 0, 0, 108, 109, 5, 105, 0, 0, 109, 110, 5,
		102, 0, 0, 110, 36, 1, 0, 0, 0, 111, 112, 5, 101, 0, 0, 112, 113, 5, 108,
		0, 0, 113, 114, 5, 115, 0, 0, 114, 115, 5, 101, 0, 0, 115, 38, 1, 0, 0,
		0, 116, 117, 5, 102, 0, 0, 117, 118, 5, 110, 0, 0, 118, 40, 1, 0, 0, 0,
		119, 120, 5, 123, 0, 0, 120, 42, 1, 0, 0, 0, 121, 122, 5, 125, 0, 0, 122,
		44, 1, 0, 0, 0, 123, 124, 5, 116, 0, 0, 124, 125, 5, 114, 0, 0, 125, 126,
		5, 117, 0, 0, 126, 127, 5, 101, 0, 0, 127, 46, 1, 0, 0, 0, 128, 129, 5,
		102, 0, 0, 129, 130, 5, 97, 0, 0, 130, 131, 5, 108, 0, 0, 131, 132, 5,
		115, 0, 0, 132, 133, 5, 101, 0, 0, 133, 48, 1, 0, 0, 0, 134, 135, 5, 44,
		0, 0, 135, 50, 1, 0, 0, 0, 136, 137, 5, 58, 0, 0, 137, 52, 1, 0, 0, 0,
		138, 143, 3, 61, 30, 0, 139, 142, 3, 61, 30, 0, 140, 142, 3, 63, 31, 0,
		141, 139, 1, 0, 0, 0, 141, 140, 1, 0, 0, 0, 142, 145, 1, 0, 0, 0, 143,
		141, 1, 0, 0, 0, 143, 144, 1, 0, 0, 0, 144, 54, 1, 0, 0, 0, 145, 143, 1,
		0, 0, 0, 146, 155, 5, 48, 0, 0, 147, 151, 7, 0, 0, 0, 148, 150, 3, 63,
		31, 0, 149, 148, 1, 0, 0, 0, 150, 153, 1, 0, 0, 0, 151, 149, 1, 0, 0, 0,
		151, 152, 1, 0, 0, 0, 152, 155, 1, 0, 0, 0, 153, 151, 1, 0, 0, 0, 154,
		146, 1, 0, 0, 0, 154, 147, 1, 0, 0, 0, 155, 56, 1, 0, 0, 0, 156, 160, 5,
		34, 0, 0, 157, 159, 8, 1, 0, 0, 158, 157, 1, 0, 0, 0, 159, 162, 1, 0, 0,
		0, 160, 158, 1, 0, 0, 0, 160, 161, 1, 0, 0, 0, 161, 163, 1, 0, 0, 0, 162,
		160, 1, 0, 0, 0, 163, 164, 5, 34, 0, 0, 164, 58, 1, 0, 0, 0, 165, 167,
		7, 2, 0, 0, 166, 165, 1, 0, 0, 0, 167, 168, 1, 0, 0, 0, 168, 166, 1, 0,
		0, 0, 168, 169, 1, 0, 0, 0, 169, 170, 1, 0, 0, 0, 170, 171, 6, 29, 0, 0,
		171, 60, 1, 0, 0, 0, 172, 173, 7, 3, 0, 0, 173, 62, 1, 0, 0, 0, 174, 175,
		7, 4, 0, 0, 175, 64, 1, 0, 0, 0, 7, 0, 141, 143, 151, 154, 160, 168, 1,
		6, 0, 0,
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

// MonkeyLexerInit initializes any static state used to implement MonkeyLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewMonkeyLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func MonkeyLexerInit() {
	staticData := &MonkeyLexerLexerStaticData
	staticData.once.Do(monkeylexerLexerInit)
}

// NewMonkeyLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewMonkeyLexer(input antlr.CharStream) *MonkeyLexer {
	MonkeyLexerInit()
	l := new(MonkeyLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &MonkeyLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "Monkey.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// MonkeyLexer tokens.
const (
	MonkeyLexerT__0   = 1
	MonkeyLexerT__1   = 2
	MonkeyLexerT__2   = 3
	MonkeyLexerT__3   = 4
	MonkeyLexerT__4   = 5
	MonkeyLexerT__5   = 6
	MonkeyLexerT__6   = 7
	MonkeyLexerT__7   = 8
	MonkeyLexerT__8   = 9
	MonkeyLexerT__9   = 10
	MonkeyLexerT__10  = 11
	MonkeyLexerT__11  = 12
	MonkeyLexerT__12  = 13
	MonkeyLexerT__13  = 14
	MonkeyLexerT__14  = 15
	MonkeyLexerT__15  = 16
	MonkeyLexerT__16  = 17
	MonkeyLexerT__17  = 18
	MonkeyLexerT__18  = 19
	MonkeyLexerT__19  = 20
	MonkeyLexerT__20  = 21
	MonkeyLexerT__21  = 22
	MonkeyLexerT__22  = 23
	MonkeyLexerT__23  = 24
	MonkeyLexerT__24  = 25
	MonkeyLexerT__25  = 26
	MonkeyLexerIDENT  = 27
	MonkeyLexerINT    = 28
	MonkeyLexerSTRING = 29
	MonkeyLexerWS     = 30
)
