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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 18, 105,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 7,
	7, 60, 10, 7, 12, 7, 14, 7, 63, 11, 7, 3, 8, 3, 8, 7, 8, 67, 10, 8, 12,
	8, 14, 8, 70, 11, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 6, 10,
	79, 10, 10, 13, 10, 14, 10, 80, 3, 11, 5, 11, 84, 10, 11, 3, 11, 3, 11,
	6, 11, 88, 10, 11, 13, 11, 14, 11, 89, 3, 12, 3, 12, 3, 13, 3, 13, 3, 14,
	3, 14, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 17, 3, 17, 2, 2, 18,
	3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23,
	13, 25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 3, 2, 6, 4, 2, 12, 12, 15,
	15, 3, 2, 43, 43, 5, 2, 11, 12, 15, 15, 34, 34, 4, 2, 67, 92, 99, 124,
	2, 111, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3,
	2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17,
	3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2,
	25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2,
	2, 33, 3, 2, 2, 2, 3, 35, 3, 2, 2, 2, 5, 39, 3, 2, 2, 2, 7, 44, 3, 2, 2,
	2, 9, 49, 3, 2, 2, 2, 11, 53, 3, 2, 2, 2, 13, 57, 3, 2, 2, 2, 15, 64, 3,
	2, 2, 2, 17, 71, 3, 2, 2, 2, 19, 78, 3, 2, 2, 2, 21, 87, 3, 2, 2, 2, 23,
	91, 3, 2, 2, 2, 25, 93, 3, 2, 2, 2, 27, 95, 3, 2, 2, 2, 29, 97, 3, 2, 2,
	2, 31, 99, 3, 2, 2, 2, 33, 103, 3, 2, 2, 2, 35, 36, 7, 104, 2, 2, 36, 37,
	7, 119, 2, 2, 37, 38, 7, 112, 2, 2, 38, 4, 3, 2, 2, 2, 39, 40, 7, 104,
	2, 2, 40, 41, 7, 107, 2, 2, 41, 42, 7, 110, 2, 2, 42, 43, 7, 103, 2, 2,
	43, 6, 3, 2, 2, 2, 44, 45, 7, 107, 2, 2, 45, 46, 7, 118, 2, 2, 46, 47,
	7, 103, 2, 2, 47, 48, 7, 111, 2, 2, 48, 8, 3, 2, 2, 2, 49, 50, 7, 114,
	2, 2, 50, 51, 7, 116, 2, 2, 51, 52, 7, 103, 2, 2, 52, 10, 3, 2, 2, 2, 53,
	54, 7, 114, 2, 2, 54, 55, 7, 113, 2, 2, 55, 56, 7, 117, 2, 2, 56, 12, 3,
	2, 2, 2, 57, 61, 7, 37, 2, 2, 58, 60, 10, 2, 2, 2, 59, 58, 3, 2, 2, 2,
	60, 63, 3, 2, 2, 2, 61, 59, 3, 2, 2, 2, 61, 62, 3, 2, 2, 2, 62, 14, 3,
	2, 2, 2, 63, 61, 3, 2, 2, 2, 64, 68, 7, 42, 2, 2, 65, 67, 10, 3, 2, 2,
	66, 65, 3, 2, 2, 2, 67, 70, 3, 2, 2, 2, 68, 66, 3, 2, 2, 2, 68, 69, 3,
	2, 2, 2, 69, 16, 3, 2, 2, 2, 70, 68, 3, 2, 2, 2, 71, 72, 5, 19, 10, 2,
	72, 73, 7, 48, 2, 2, 73, 74, 7, 105, 2, 2, 74, 75, 7, 113, 2, 2, 75, 18,
	3, 2, 2, 2, 76, 79, 5, 33, 17, 2, 77, 79, 7, 97, 2, 2, 78, 76, 3, 2, 2,
	2, 78, 77, 3, 2, 2, 2, 79, 80, 3, 2, 2, 2, 80, 78, 3, 2, 2, 2, 80, 81,
	3, 2, 2, 2, 81, 20, 3, 2, 2, 2, 82, 84, 7, 15, 2, 2, 83, 82, 3, 2, 2, 2,
	83, 84, 3, 2, 2, 2, 84, 85, 3, 2, 2, 2, 85, 88, 7, 12, 2, 2, 86, 88, 7,
	15, 2, 2, 87, 83, 3, 2, 2, 2, 87, 86, 3, 2, 2, 2, 88, 89, 3, 2, 2, 2, 89,
	87, 3, 2, 2, 2, 89, 90, 3, 2, 2, 2, 90, 22, 3, 2, 2, 2, 91, 92, 7, 43,
	2, 2, 92, 24, 3, 2, 2, 2, 93, 94, 7, 49, 2, 2, 94, 26, 3, 2, 2, 2, 95,
	96, 7, 48, 2, 2, 96, 28, 3, 2, 2, 2, 97, 98, 7, 94, 2, 2, 98, 30, 3, 2,
	2, 2, 99, 100, 9, 4, 2, 2, 100, 101, 3, 2, 2, 2, 101, 102, 8, 16, 2, 2,
	102, 32, 3, 2, 2, 2, 103, 104, 9, 5, 2, 2, 104, 34, 3, 2, 2, 2, 10, 2,
	61, 68, 78, 80, 83, 87, 89, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'fun'", "'file'", "'item'", "'pre'", "'pos'", "", "", "", "", "",
	"')'", "'/'", "'.'", "'\\'",
}

var lexerSymbolicNames = []string{
	"", "FUN", "FILE", "ITEM", "PRE", "POS", "COMMENT", "ARG_NAME", "FILENAME",
	"WORD", "NEWLINE", "CP", "FS", "DT", "ML", "WS", "ID",
}

var lexerRuleNames = []string{
	"FUN", "FILE", "ITEM", "PRE", "POS", "COMMENT", "ARG_NAME", "FILENAME",
	"WORD", "NEWLINE", "CP", "FS", "DT", "ML", "WS", "ID",
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
	SkeleLexerFUN      = 1
	SkeleLexerFILE     = 2
	SkeleLexerITEM     = 3
	SkeleLexerPRE      = 4
	SkeleLexerPOS      = 5
	SkeleLexerCOMMENT  = 6
	SkeleLexerARG_NAME = 7
	SkeleLexerFILENAME = 8
	SkeleLexerWORD     = 9
	SkeleLexerNEWLINE  = 10
	SkeleLexerCP       = 11
	SkeleLexerFS       = 12
	SkeleLexerDT       = 13
	SkeleLexerML       = 14
	SkeleLexerWS       = 15
	SkeleLexerID       = 16
)
