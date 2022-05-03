// Code generated from SkeleLexer.g4 by ANTLR 4.9.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 18, 133,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3,
	7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10,
	3, 10, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 7, 11, 78, 10, 11, 12, 11, 14,
	11, 81, 11, 11, 3, 12, 3, 12, 7, 12, 85, 10, 12, 12, 12, 14, 12, 88, 11,
	12, 3, 13, 3, 13, 3, 13, 7, 13, 93, 10, 13, 12, 13, 14, 13, 96, 11, 13,
	3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 14, 6,
	14, 108, 10, 14, 13, 14, 14, 14, 109, 3, 14, 3, 14, 7, 14, 114, 10, 14,
	12, 14, 14, 14, 117, 11, 14, 3, 15, 5, 15, 120, 10, 15, 3, 15, 3, 15, 6,
	15, 124, 10, 15, 13, 15, 14, 15, 125, 3, 16, 3, 16, 3, 16, 3, 16, 3, 17,
	3, 17, 2, 2, 18, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19,
	11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 3, 2, 7, 5,
	2, 12, 12, 15, 15, 37, 37, 4, 2, 12, 12, 15, 15, 4, 2, 50, 59, 97, 97,
	5, 2, 11, 12, 15, 15, 34, 34, 4, 2, 67, 92, 99, 124, 2, 142, 2, 3, 3, 2,
	2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2,
	2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3,
	2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27,
	3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 3,
	35, 3, 2, 2, 2, 5, 39, 3, 2, 2, 2, 7, 43, 3, 2, 2, 2, 9, 47, 3, 2, 2, 2,
	11, 52, 3, 2, 2, 2, 13, 56, 3, 2, 2, 2, 15, 60, 3, 2, 2, 2, 17, 64, 3,
	2, 2, 2, 19, 68, 3, 2, 2, 2, 21, 75, 3, 2, 2, 2, 23, 82, 3, 2, 2, 2, 25,
	89, 3, 2, 2, 2, 27, 107, 3, 2, 2, 2, 29, 123, 3, 2, 2, 2, 31, 127, 3, 2,
	2, 2, 33, 131, 3, 2, 2, 2, 35, 36, 7, 114, 2, 2, 36, 37, 7, 109, 2, 2,
	37, 38, 7, 105, 2, 2, 38, 4, 3, 2, 2, 2, 39, 40, 7, 104, 2, 2, 40, 41,
	7, 113, 2, 2, 41, 42, 7, 110, 2, 2, 42, 6, 3, 2, 2, 2, 43, 44, 7, 102,
	2, 2, 44, 45, 7, 113, 2, 2, 45, 46, 7, 101, 2, 2, 46, 8, 3, 2, 2, 2, 47,
	48, 7, 104, 2, 2, 48, 49, 7, 107, 2, 2, 49, 50, 7, 110, 2, 2, 50, 51, 7,
	103, 2, 2, 51, 10, 3, 2, 2, 2, 52, 53, 7, 104, 2, 2, 53, 54, 7, 119, 2,
	2, 54, 55, 7, 112, 2, 2, 55, 12, 3, 2, 2, 2, 56, 57, 7, 114, 2, 2, 57,
	58, 7, 116, 2, 2, 58, 59, 7, 103, 2, 2, 59, 14, 3, 2, 2, 2, 60, 61, 7,
	114, 2, 2, 61, 62, 7, 113, 2, 2, 62, 63, 7, 117, 2, 2, 63, 16, 3, 2, 2,
	2, 64, 65, 7, 104, 2, 2, 65, 66, 7, 117, 2, 2, 66, 67, 7, 111, 2, 2, 67,
	18, 3, 2, 2, 2, 68, 69, 7, 117, 2, 2, 69, 70, 7, 118, 2, 2, 70, 71, 7,
	99, 2, 2, 71, 72, 7, 118, 2, 2, 72, 73, 7, 103, 2, 2, 73, 74, 7, 117, 2,
	2, 74, 20, 3, 2, 2, 2, 75, 79, 7, 94, 2, 2, 76, 78, 10, 2, 2, 2, 77, 76,
	3, 2, 2, 2, 78, 81, 3, 2, 2, 2, 79, 77, 3, 2, 2, 2, 79, 80, 3, 2, 2, 2,
	80, 22, 3, 2, 2, 2, 81, 79, 3, 2, 2, 2, 82, 86, 7, 37, 2, 2, 83, 85, 10,
	3, 2, 2, 84, 83, 3, 2, 2, 2, 85, 88, 3, 2, 2, 2, 86, 84, 3, 2, 2, 2, 86,
	87, 3, 2, 2, 2, 87, 24, 3, 2, 2, 2, 88, 86, 3, 2, 2, 2, 89, 94, 5, 33,
	17, 2, 90, 93, 5, 33, 17, 2, 91, 93, 9, 4, 2, 2, 92, 90, 3, 2, 2, 2, 92,
	91, 3, 2, 2, 2, 93, 96, 3, 2, 2, 2, 94, 92, 3, 2, 2, 2, 94, 95, 3, 2, 2,
	2, 95, 97, 3, 2, 2, 2, 96, 94, 3, 2, 2, 2, 97, 98, 7, 97, 2, 2, 98, 99,
	7, 118, 2, 2, 99, 100, 7, 103, 2, 2, 100, 101, 7, 117, 2, 2, 101, 102,
	7, 118, 2, 2, 102, 103, 7, 48, 2, 2, 103, 104, 7, 105, 2, 2, 104, 105,
	7, 113, 2, 2, 105, 26, 3, 2, 2, 2, 106, 108, 5, 33, 17, 2, 107, 106, 3,
	2, 2, 2, 108, 109, 3, 2, 2, 2, 109, 107, 3, 2, 2, 2, 109, 110, 3, 2, 2,
	2, 110, 115, 3, 2, 2, 2, 111, 114, 5, 33, 17, 2, 112, 114, 7, 97, 2, 2,
	113, 111, 3, 2, 2, 2, 113, 112, 3, 2, 2, 2, 114, 117, 3, 2, 2, 2, 115,
	113, 3, 2, 2, 2, 115, 116, 3, 2, 2, 2, 116, 28, 3, 2, 2, 2, 117, 115, 3,
	2, 2, 2, 118, 120, 7, 15, 2, 2, 119, 118, 3, 2, 2, 2, 119, 120, 3, 2, 2,
	2, 120, 121, 3, 2, 2, 2, 121, 124, 7, 12, 2, 2, 122, 124, 7, 15, 2, 2,
	123, 119, 3, 2, 2, 2, 123, 122, 3, 2, 2, 2, 124, 125, 3, 2, 2, 2, 125,
	123, 3, 2, 2, 2, 125, 126, 3, 2, 2, 2, 126, 30, 3, 2, 2, 2, 127, 128, 9,
	5, 2, 2, 128, 129, 3, 2, 2, 2, 129, 130, 8, 16, 2, 2, 130, 32, 3, 2, 2,
	2, 131, 132, 9, 6, 2, 2, 132, 34, 3, 2, 2, 2, 13, 2, 79, 86, 92, 94, 109,
	113, 115, 119, 123, 125, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'pkg'", "'fol'", "'doc'", "'file'", "'fun'", "'pre'", "'pos'", "'fsm'",
	"'states'",
}

var lexerSymbolicNames = []string{
	"", "PKG", "FOL", "DOC", "FILE", "FUN", "PRE", "POS", "FSM", "STS", "LINE",
	"COMMENT", "FILENAME", "WORD", "NEWLINE", "WS", "ID",
}

var lexerRuleNames = []string{
	"PKG", "FOL", "DOC", "FILE", "FUN", "PRE", "POS", "FSM", "STS", "LINE",
	"COMMENT", "FILENAME", "WORD", "NEWLINE", "WS", "ID",
}

type SkeleLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewSkeleLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *SkeleLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewSkeleLexer(input antlr.CharStream) *SkeleLexer {
	l := new(SkeleLexer)
	lexerDeserializer := antlr.NewATNDeserializer(nil)
	lexerAtn := lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)
	lexerDecisionToDFA := make([]*antlr.DFA, len(lexerAtn.DecisionToState))
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "SkeleLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// SkeleLexer tokens.
const (
	SkeleLexerPKG      = 1
	SkeleLexerFOL      = 2
	SkeleLexerDOC      = 3
	SkeleLexerFILE     = 4
	SkeleLexerFUN      = 5
	SkeleLexerPRE      = 6
	SkeleLexerPOS      = 7
	SkeleLexerFSM      = 8
	SkeleLexerSTS      = 9
	SkeleLexerLINE     = 10
	SkeleLexerCOMMENT  = 11
	SkeleLexerFILENAME = 12
	SkeleLexerWORD     = 13
	SkeleLexerNEWLINE  = 14
	SkeleLexerWS       = 15
	SkeleLexerID       = 16
)
