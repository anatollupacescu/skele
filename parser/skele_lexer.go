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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 17, 122,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3,
	5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3,
	8, 3, 8, 3, 8, 3, 9, 3, 9, 7, 9, 65, 10, 9, 12, 9, 14, 9, 68, 11, 9, 3,
	10, 3, 10, 7, 10, 72, 10, 10, 12, 10, 14, 10, 75, 11, 10, 3, 11, 3, 11,
	3, 11, 7, 11, 80, 10, 11, 12, 11, 14, 11, 83, 11, 11, 3, 11, 3, 11, 3,
	11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 12, 6, 12, 95, 10, 12,
	13, 12, 14, 12, 96, 3, 12, 3, 12, 7, 12, 101, 10, 12, 12, 12, 14, 12, 104,
	11, 12, 3, 13, 5, 13, 107, 10, 13, 3, 13, 3, 13, 6, 13, 111, 10, 13, 13,
	13, 14, 13, 112, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16,
	2, 2, 17, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11,
	21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17, 3, 2, 7, 5, 2, 12, 12,
	15, 15, 37, 37, 4, 2, 12, 12, 15, 15, 4, 2, 50, 59, 97, 97, 5, 2, 11, 12,
	15, 15, 34, 34, 4, 2, 67, 92, 99, 124, 2, 131, 2, 3, 3, 2, 2, 2, 2, 5,
	3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13,
	3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2,
	21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2,
	2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 3, 33, 3, 2, 2, 2, 5, 37, 3, 2, 2,
	2, 7, 41, 3, 2, 2, 2, 9, 45, 3, 2, 2, 2, 11, 50, 3, 2, 2, 2, 13, 54, 3,
	2, 2, 2, 15, 58, 3, 2, 2, 2, 17, 62, 3, 2, 2, 2, 19, 69, 3, 2, 2, 2, 21,
	76, 3, 2, 2, 2, 23, 94, 3, 2, 2, 2, 25, 110, 3, 2, 2, 2, 27, 114, 3, 2,
	2, 2, 29, 116, 3, 2, 2, 2, 31, 120, 3, 2, 2, 2, 33, 34, 7, 114, 2, 2, 34,
	35, 7, 109, 2, 2, 35, 36, 7, 105, 2, 2, 36, 4, 3, 2, 2, 2, 37, 38, 7, 104,
	2, 2, 38, 39, 7, 113, 2, 2, 39, 40, 7, 110, 2, 2, 40, 6, 3, 2, 2, 2, 41,
	42, 7, 102, 2, 2, 42, 43, 7, 113, 2, 2, 43, 44, 7, 101, 2, 2, 44, 8, 3,
	2, 2, 2, 45, 46, 7, 104, 2, 2, 46, 47, 7, 107, 2, 2, 47, 48, 7, 110, 2,
	2, 48, 49, 7, 103, 2, 2, 49, 10, 3, 2, 2, 2, 50, 51, 7, 104, 2, 2, 51,
	52, 7, 119, 2, 2, 52, 53, 7, 112, 2, 2, 53, 12, 3, 2, 2, 2, 54, 55, 7,
	114, 2, 2, 55, 56, 7, 116, 2, 2, 56, 57, 7, 103, 2, 2, 57, 14, 3, 2, 2,
	2, 58, 59, 7, 114, 2, 2, 59, 60, 7, 113, 2, 2, 60, 61, 7, 117, 2, 2, 61,
	16, 3, 2, 2, 2, 62, 66, 7, 94, 2, 2, 63, 65, 10, 2, 2, 2, 64, 63, 3, 2,
	2, 2, 65, 68, 3, 2, 2, 2, 66, 64, 3, 2, 2, 2, 66, 67, 3, 2, 2, 2, 67, 18,
	3, 2, 2, 2, 68, 66, 3, 2, 2, 2, 69, 73, 7, 37, 2, 2, 70, 72, 10, 3, 2,
	2, 71, 70, 3, 2, 2, 2, 72, 75, 3, 2, 2, 2, 73, 71, 3, 2, 2, 2, 73, 74,
	3, 2, 2, 2, 74, 20, 3, 2, 2, 2, 75, 73, 3, 2, 2, 2, 76, 81, 5, 31, 16,
	2, 77, 80, 5, 31, 16, 2, 78, 80, 9, 4, 2, 2, 79, 77, 3, 2, 2, 2, 79, 78,
	3, 2, 2, 2, 80, 83, 3, 2, 2, 2, 81, 79, 3, 2, 2, 2, 81, 82, 3, 2, 2, 2,
	82, 84, 3, 2, 2, 2, 83, 81, 3, 2, 2, 2, 84, 85, 7, 97, 2, 2, 85, 86, 7,
	118, 2, 2, 86, 87, 7, 103, 2, 2, 87, 88, 7, 117, 2, 2, 88, 89, 7, 118,
	2, 2, 89, 90, 7, 48, 2, 2, 90, 91, 7, 105, 2, 2, 91, 92, 7, 113, 2, 2,
	92, 22, 3, 2, 2, 2, 93, 95, 5, 31, 16, 2, 94, 93, 3, 2, 2, 2, 95, 96, 3,
	2, 2, 2, 96, 94, 3, 2, 2, 2, 96, 97, 3, 2, 2, 2, 97, 102, 3, 2, 2, 2, 98,
	101, 5, 31, 16, 2, 99, 101, 7, 97, 2, 2, 100, 98, 3, 2, 2, 2, 100, 99,
	3, 2, 2, 2, 101, 104, 3, 2, 2, 2, 102, 100, 3, 2, 2, 2, 102, 103, 3, 2,
	2, 2, 103, 24, 3, 2, 2, 2, 104, 102, 3, 2, 2, 2, 105, 107, 7, 15, 2, 2,
	106, 105, 3, 2, 2, 2, 106, 107, 3, 2, 2, 2, 107, 108, 3, 2, 2, 2, 108,
	111, 7, 12, 2, 2, 109, 111, 7, 15, 2, 2, 110, 106, 3, 2, 2, 2, 110, 109,
	3, 2, 2, 2, 111, 112, 3, 2, 2, 2, 112, 110, 3, 2, 2, 2, 112, 113, 3, 2,
	2, 2, 113, 26, 3, 2, 2, 2, 114, 115, 7, 49, 2, 2, 115, 28, 3, 2, 2, 2,
	116, 117, 9, 5, 2, 2, 117, 118, 3, 2, 2, 2, 118, 119, 8, 15, 2, 2, 119,
	30, 3, 2, 2, 2, 120, 121, 9, 6, 2, 2, 121, 32, 3, 2, 2, 2, 13, 2, 66, 73,
	79, 81, 96, 100, 102, 106, 110, 112, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'pkg'", "'fol'", "'doc'", "'file'", "'fun'", "'pre'", "'pos'", "",
	"", "", "", "", "'/'",
}

var lexerSymbolicNames = []string{
	"", "PKG", "FOL", "DOC", "FILE", "FUN", "PRE", "POS", "LINE", "COMMENT",
	"FILENAME", "WORD", "NEWLINE", "FS", "WS", "ID",
}

var lexerRuleNames = []string{
	"PKG", "FOL", "DOC", "FILE", "FUN", "PRE", "POS", "LINE", "COMMENT", "FILENAME",
	"WORD", "NEWLINE", "FS", "WS", "ID",
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
	SkeleLexerLINE     = 8
	SkeleLexerCOMMENT  = 9
	SkeleLexerFILENAME = 10
	SkeleLexerWORD     = 11
	SkeleLexerNEWLINE  = 12
	SkeleLexerFS       = 13
	SkeleLexerWS       = 14
	SkeleLexerID       = 15
)
