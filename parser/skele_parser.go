// Code generated from Skele.g4 by ANTLR 4.9.1. DO NOT EDIT.

package parser // Skele

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 17, 104,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 3, 2, 6, 2, 28, 10, 2, 13, 2, 14, 2, 29, 3, 2, 3, 2, 3, 3, 3, 3,
	5, 3, 36, 10, 3, 3, 3, 5, 3, 39, 10, 3, 3, 3, 6, 3, 42, 10, 3, 13, 3, 14,
	3, 43, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6,
	6, 6, 57, 10, 6, 13, 6, 14, 6, 58, 3, 7, 3, 7, 3, 7, 3, 7, 6, 7, 65, 10,
	7, 13, 7, 14, 7, 66, 3, 8, 3, 8, 3, 8, 3, 8, 5, 8, 73, 10, 8, 3, 8, 3,
	8, 3, 9, 3, 9, 3, 9, 6, 9, 80, 10, 9, 13, 9, 14, 9, 81, 3, 10, 3, 10, 3,
	10, 6, 10, 87, 10, 10, 13, 10, 14, 10, 88, 3, 11, 3, 11, 3, 11, 3, 12,
	5, 12, 95, 10, 12, 3, 12, 6, 12, 98, 10, 12, 13, 12, 14, 12, 99, 3, 13,
	3, 13, 3, 13, 2, 2, 14, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 2,
	2, 2, 102, 2, 27, 3, 2, 2, 2, 4, 33, 3, 2, 2, 2, 6, 45, 3, 2, 2, 2, 8,
	49, 3, 2, 2, 2, 10, 53, 3, 2, 2, 2, 12, 60, 3, 2, 2, 2, 14, 68, 3, 2, 2,
	2, 16, 76, 3, 2, 2, 2, 18, 83, 3, 2, 2, 2, 20, 90, 3, 2, 2, 2, 22, 97,
	3, 2, 2, 2, 24, 101, 3, 2, 2, 2, 26, 28, 5, 4, 3, 2, 27, 26, 3, 2, 2, 2,
	28, 29, 3, 2, 2, 2, 29, 27, 3, 2, 2, 2, 29, 30, 3, 2, 2, 2, 30, 31, 3,
	2, 2, 2, 31, 32, 7, 2, 2, 3, 32, 3, 3, 2, 2, 2, 33, 35, 5, 8, 5, 2, 34,
	36, 5, 6, 4, 2, 35, 34, 3, 2, 2, 2, 35, 36, 3, 2, 2, 2, 36, 38, 3, 2, 2,
	2, 37, 39, 5, 10, 6, 2, 38, 37, 3, 2, 2, 2, 38, 39, 3, 2, 2, 2, 39, 41,
	3, 2, 2, 2, 40, 42, 5, 12, 7, 2, 41, 40, 3, 2, 2, 2, 42, 43, 3, 2, 2, 2,
	43, 41, 3, 2, 2, 2, 43, 44, 3, 2, 2, 2, 44, 5, 3, 2, 2, 2, 45, 46, 7, 4,
	2, 2, 46, 47, 7, 13, 2, 2, 47, 48, 5, 22, 12, 2, 48, 7, 3, 2, 2, 2, 49,
	50, 7, 3, 2, 2, 50, 51, 7, 13, 2, 2, 51, 52, 5, 22, 12, 2, 52, 9, 3, 2,
	2, 2, 53, 54, 7, 5, 2, 2, 54, 56, 5, 22, 12, 2, 55, 57, 5, 20, 11, 2, 56,
	55, 3, 2, 2, 2, 57, 58, 3, 2, 2, 2, 58, 56, 3, 2, 2, 2, 58, 59, 3, 2, 2,
	2, 59, 11, 3, 2, 2, 2, 60, 61, 7, 6, 2, 2, 61, 62, 7, 12, 2, 2, 62, 64,
	5, 22, 12, 2, 63, 65, 5, 14, 8, 2, 64, 63, 3, 2, 2, 2, 65, 66, 3, 2, 2,
	2, 66, 64, 3, 2, 2, 2, 66, 67, 3, 2, 2, 2, 67, 13, 3, 2, 2, 2, 68, 69,
	7, 7, 2, 2, 69, 70, 5, 22, 12, 2, 70, 72, 5, 20, 11, 2, 71, 73, 5, 16,
	9, 2, 72, 71, 3, 2, 2, 2, 72, 73, 3, 2, 2, 2, 73, 74, 3, 2, 2, 2, 74, 75,
	5, 18, 10, 2, 75, 15, 3, 2, 2, 2, 76, 77, 7, 8, 2, 2, 77, 79, 5, 22, 12,
	2, 78, 80, 5, 20, 11, 2, 79, 78, 3, 2, 2, 2, 80, 81, 3, 2, 2, 2, 81, 79,
	3, 2, 2, 2, 81, 82, 3, 2, 2, 2, 82, 17, 3, 2, 2, 2, 83, 84, 7, 9, 2, 2,
	84, 86, 5, 22, 12, 2, 85, 87, 5, 20, 11, 2, 86, 85, 3, 2, 2, 2, 87, 88,
	3, 2, 2, 2, 88, 86, 3, 2, 2, 2, 88, 89, 3, 2, 2, 2, 89, 19, 3, 2, 2, 2,
	90, 91, 7, 10, 2, 2, 91, 92, 5, 22, 12, 2, 92, 21, 3, 2, 2, 2, 93, 95,
	5, 24, 13, 2, 94, 93, 3, 2, 2, 2, 94, 95, 3, 2, 2, 2, 95, 96, 3, 2, 2,
	2, 96, 98, 7, 14, 2, 2, 97, 94, 3, 2, 2, 2, 98, 99, 3, 2, 2, 2, 99, 97,
	3, 2, 2, 2, 99, 100, 3, 2, 2, 2, 100, 23, 3, 2, 2, 2, 101, 102, 7, 11,
	2, 2, 102, 25, 3, 2, 2, 2, 13, 29, 35, 38, 43, 58, 66, 72, 81, 88, 94,
	99,
}
var literalNames = []string{
	"", "'pkg'", "'fol'", "'doc'", "'file'", "'fun'", "'pre'", "'pos'", "",
	"", "", "", "", "'/'",
}
var symbolicNames = []string{
	"", "PKG", "FOL", "DOC", "FILE", "FUN", "PRE", "POS", "LINE", "COMMENT",
	"FILENAME", "WORD", "NEWLINE", "FS", "WS", "ID",
}

var ruleNames = []string{
	"start", "spec", "fol", "pkg", "doc", "file", "fun", "pre", "pos", "ln",
	"newln", "comment",
}

type Skele struct {
	*antlr.BaseParser
}

// NewSkele produces a new parser instance for the optional input antlr.TokenStream.
//
// The *Skele instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewSkele(input antlr.TokenStream) *Skele {
	this := new(Skele)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Skele.g4"

	return this
}

// Skele tokens.
const (
	SkeleEOF      = antlr.TokenEOF
	SkelePKG      = 1
	SkeleFOL      = 2
	SkeleDOC      = 3
	SkeleFILE     = 4
	SkeleFUN      = 5
	SkelePRE      = 6
	SkelePOS      = 7
	SkeleLINE     = 8
	SkeleCOMMENT  = 9
	SkeleFILENAME = 10
	SkeleWORD     = 11
	SkeleNEWLINE  = 12
	SkeleFS       = 13
	SkeleWS       = 14
	SkeleID       = 15
)

// Skele rules.
const (
	SkeleRULE_start   = 0
	SkeleRULE_spec    = 1
	SkeleRULE_fol     = 2
	SkeleRULE_pkg     = 3
	SkeleRULE_doc     = 4
	SkeleRULE_file    = 5
	SkeleRULE_fun     = 6
	SkeleRULE_pre     = 7
	SkeleRULE_pos     = 8
	SkeleRULE_ln      = 9
	SkeleRULE_newln   = 10
	SkeleRULE_comment = 11
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SkeleRULE_start
	return p
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SkeleRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) EOF() antlr.TerminalNode {
	return s.GetToken(SkeleEOF, 0)
}

func (s *StartContext) AllSpec() []ISpecContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISpecContext)(nil)).Elem())
	var tst = make([]ISpecContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISpecContext)
		}
	}

	return tst
}

func (s *StartContext) Spec(i int) ISpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISpecContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISpecContext)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SkeleListener); ok {
		listenerT.EnterStart(s)
	}
}

func (s *StartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SkeleListener); ok {
		listenerT.ExitStart(s)
	}
}

func (p *Skele) Start() (localctx IStartContext) {
	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SkeleRULE_start)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(25)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == SkelePKG {
		{
			p.SetState(24)
			p.Spec()
		}

		p.SetState(27)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(29)
		p.Match(SkeleEOF)
	}

	return localctx
}

// ISpecContext is an interface to support dynamic dispatch.
type ISpecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSpecContext differentiates from other interfaces.
	IsSpecContext()
}

type SpecContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySpecContext() *SpecContext {
	var p = new(SpecContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SkeleRULE_spec
	return p
}

func (*SpecContext) IsSpecContext() {}

func NewSpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SpecContext {
	var p = new(SpecContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SkeleRULE_spec

	return p
}

func (s *SpecContext) GetParser() antlr.Parser { return s.parser }

func (s *SpecContext) Pkg() IPkgContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPkgContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPkgContext)
}

func (s *SpecContext) Fol() IFolContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFolContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFolContext)
}

func (s *SpecContext) Doc() IDocContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDocContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDocContext)
}

func (s *SpecContext) AllFile() []IFileContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFileContext)(nil)).Elem())
	var tst = make([]IFileContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFileContext)
		}
	}

	return tst
}

func (s *SpecContext) File(i int) IFileContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFileContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFileContext)
}

func (s *SpecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SpecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SpecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SkeleListener); ok {
		listenerT.EnterSpec(s)
	}
}

func (s *SpecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SkeleListener); ok {
		listenerT.ExitSpec(s)
	}
}

func (p *Skele) Spec() (localctx ISpecContext) {
	localctx = NewSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SkeleRULE_spec)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(31)
		p.Pkg()
	}
	p.SetState(33)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SkeleFOL {
		{
			p.SetState(32)
			p.Fol()
		}

	}
	p.SetState(36)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SkeleDOC {
		{
			p.SetState(35)
			p.Doc()
		}

	}
	p.SetState(39)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == SkeleFILE {
		{
			p.SetState(38)
			p.File()
		}

		p.SetState(41)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IFolContext is an interface to support dynamic dispatch.
type IFolContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFolContext differentiates from other interfaces.
	IsFolContext()
}

type FolContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFolContext() *FolContext {
	var p = new(FolContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SkeleRULE_fol
	return p
}

func (*FolContext) IsFolContext() {}

func NewFolContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FolContext {
	var p = new(FolContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SkeleRULE_fol

	return p
}

func (s *FolContext) GetParser() antlr.Parser { return s.parser }

func (s *FolContext) FOL() antlr.TerminalNode {
	return s.GetToken(SkeleFOL, 0)
}

func (s *FolContext) WORD() antlr.TerminalNode {
	return s.GetToken(SkeleWORD, 0)
}

func (s *FolContext) Newln() INewlnContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INewlnContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INewlnContext)
}

func (s *FolContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FolContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FolContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SkeleListener); ok {
		listenerT.EnterFol(s)
	}
}

func (s *FolContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SkeleListener); ok {
		listenerT.ExitFol(s)
	}
}

func (p *Skele) Fol() (localctx IFolContext) {
	localctx = NewFolContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, SkeleRULE_fol)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(43)
		p.Match(SkeleFOL)
	}
	{
		p.SetState(44)
		p.Match(SkeleWORD)
	}
	{
		p.SetState(45)
		p.Newln()
	}

	return localctx
}

// IPkgContext is an interface to support dynamic dispatch.
type IPkgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPkgContext differentiates from other interfaces.
	IsPkgContext()
}

type PkgContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPkgContext() *PkgContext {
	var p = new(PkgContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SkeleRULE_pkg
	return p
}

func (*PkgContext) IsPkgContext() {}

func NewPkgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PkgContext {
	var p = new(PkgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SkeleRULE_pkg

	return p
}

func (s *PkgContext) GetParser() antlr.Parser { return s.parser }

func (s *PkgContext) PKG() antlr.TerminalNode {
	return s.GetToken(SkelePKG, 0)
}

func (s *PkgContext) WORD() antlr.TerminalNode {
	return s.GetToken(SkeleWORD, 0)
}

func (s *PkgContext) Newln() INewlnContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INewlnContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INewlnContext)
}

func (s *PkgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PkgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PkgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SkeleListener); ok {
		listenerT.EnterPkg(s)
	}
}

func (s *PkgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SkeleListener); ok {
		listenerT.ExitPkg(s)
	}
}

func (p *Skele) Pkg() (localctx IPkgContext) {
	localctx = NewPkgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, SkeleRULE_pkg)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(47)
		p.Match(SkelePKG)
	}
	{
		p.SetState(48)
		p.Match(SkeleWORD)
	}
	{
		p.SetState(49)
		p.Newln()
	}

	return localctx
}

// IDocContext is an interface to support dynamic dispatch.
type IDocContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDocContext differentiates from other interfaces.
	IsDocContext()
}

type DocContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDocContext() *DocContext {
	var p = new(DocContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SkeleRULE_doc
	return p
}

func (*DocContext) IsDocContext() {}

func NewDocContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DocContext {
	var p = new(DocContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SkeleRULE_doc

	return p
}

func (s *DocContext) GetParser() antlr.Parser { return s.parser }

func (s *DocContext) DOC() antlr.TerminalNode {
	return s.GetToken(SkeleDOC, 0)
}

func (s *DocContext) Newln() INewlnContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INewlnContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INewlnContext)
}

func (s *DocContext) AllLn() []ILnContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ILnContext)(nil)).Elem())
	var tst = make([]ILnContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ILnContext)
		}
	}

	return tst
}

func (s *DocContext) Ln(i int) ILnContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILnContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ILnContext)
}

func (s *DocContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DocContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DocContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SkeleListener); ok {
		listenerT.EnterDoc(s)
	}
}

func (s *DocContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SkeleListener); ok {
		listenerT.ExitDoc(s)
	}
}

func (p *Skele) Doc() (localctx IDocContext) {
	localctx = NewDocContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, SkeleRULE_doc)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(51)
		p.Match(SkeleDOC)
	}
	{
		p.SetState(52)
		p.Newln()
	}
	p.SetState(54)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == SkeleLINE {
		{
			p.SetState(53)
			p.Ln()
		}

		p.SetState(56)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IFileContext is an interface to support dynamic dispatch.
type IFileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFileContext differentiates from other interfaces.
	IsFileContext()
}

type FileContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFileContext() *FileContext {
	var p = new(FileContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SkeleRULE_file
	return p
}

func (*FileContext) IsFileContext() {}

func NewFileContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FileContext {
	var p = new(FileContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SkeleRULE_file

	return p
}

func (s *FileContext) GetParser() antlr.Parser { return s.parser }

func (s *FileContext) FILE() antlr.TerminalNode {
	return s.GetToken(SkeleFILE, 0)
}

func (s *FileContext) FILENAME() antlr.TerminalNode {
	return s.GetToken(SkeleFILENAME, 0)
}

func (s *FileContext) Newln() INewlnContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INewlnContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INewlnContext)
}

func (s *FileContext) AllFun() []IFunContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFunContext)(nil)).Elem())
	var tst = make([]IFunContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFunContext)
		}
	}

	return tst
}

func (s *FileContext) Fun(i int) IFunContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFunContext)
}

func (s *FileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FileContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FileContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SkeleListener); ok {
		listenerT.EnterFile(s)
	}
}

func (s *FileContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SkeleListener); ok {
		listenerT.ExitFile(s)
	}
}

func (p *Skele) File() (localctx IFileContext) {
	localctx = NewFileContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, SkeleRULE_file)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(58)
		p.Match(SkeleFILE)
	}
	{
		p.SetState(59)
		p.Match(SkeleFILENAME)
	}
	{
		p.SetState(60)
		p.Newln()
	}
	p.SetState(62)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == SkeleFUN {
		{
			p.SetState(61)
			p.Fun()
		}

		p.SetState(64)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IFunContext is an interface to support dynamic dispatch.
type IFunContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunContext differentiates from other interfaces.
	IsFunContext()
}

type FunContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunContext() *FunContext {
	var p = new(FunContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SkeleRULE_fun
	return p
}

func (*FunContext) IsFunContext() {}

func NewFunContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunContext {
	var p = new(FunContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SkeleRULE_fun

	return p
}

func (s *FunContext) GetParser() antlr.Parser { return s.parser }

func (s *FunContext) FUN() antlr.TerminalNode {
	return s.GetToken(SkeleFUN, 0)
}

func (s *FunContext) Newln() INewlnContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INewlnContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INewlnContext)
}

func (s *FunContext) Ln() ILnContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILnContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILnContext)
}

func (s *FunContext) Pos() IPosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPosContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPosContext)
}

func (s *FunContext) Pre() IPreContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPreContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPreContext)
}

func (s *FunContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SkeleListener); ok {
		listenerT.EnterFun(s)
	}
}

func (s *FunContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SkeleListener); ok {
		listenerT.ExitFun(s)
	}
}

func (p *Skele) Fun() (localctx IFunContext) {
	localctx = NewFunContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, SkeleRULE_fun)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(66)
		p.Match(SkeleFUN)
	}
	{
		p.SetState(67)
		p.Newln()
	}
	{
		p.SetState(68)
		p.Ln()
	}
	p.SetState(70)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SkelePRE {
		{
			p.SetState(69)
			p.Pre()
		}

	}
	{
		p.SetState(72)
		p.Pos()
	}

	return localctx
}

// IPreContext is an interface to support dynamic dispatch.
type IPreContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPreContext differentiates from other interfaces.
	IsPreContext()
}

type PreContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPreContext() *PreContext {
	var p = new(PreContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SkeleRULE_pre
	return p
}

func (*PreContext) IsPreContext() {}

func NewPreContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PreContext {
	var p = new(PreContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SkeleRULE_pre

	return p
}

func (s *PreContext) GetParser() antlr.Parser { return s.parser }

func (s *PreContext) PRE() antlr.TerminalNode {
	return s.GetToken(SkelePRE, 0)
}

func (s *PreContext) Newln() INewlnContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INewlnContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INewlnContext)
}

func (s *PreContext) AllLn() []ILnContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ILnContext)(nil)).Elem())
	var tst = make([]ILnContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ILnContext)
		}
	}

	return tst
}

func (s *PreContext) Ln(i int) ILnContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILnContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ILnContext)
}

func (s *PreContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PreContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PreContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SkeleListener); ok {
		listenerT.EnterPre(s)
	}
}

func (s *PreContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SkeleListener); ok {
		listenerT.ExitPre(s)
	}
}

func (p *Skele) Pre() (localctx IPreContext) {
	localctx = NewPreContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, SkeleRULE_pre)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(74)
		p.Match(SkelePRE)
	}
	{
		p.SetState(75)
		p.Newln()
	}
	p.SetState(77)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == SkeleLINE {
		{
			p.SetState(76)
			p.Ln()
		}

		p.SetState(79)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IPosContext is an interface to support dynamic dispatch.
type IPosContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPosContext differentiates from other interfaces.
	IsPosContext()
}

type PosContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPosContext() *PosContext {
	var p = new(PosContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SkeleRULE_pos
	return p
}

func (*PosContext) IsPosContext() {}

func NewPosContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PosContext {
	var p = new(PosContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SkeleRULE_pos

	return p
}

func (s *PosContext) GetParser() antlr.Parser { return s.parser }

func (s *PosContext) POS() antlr.TerminalNode {
	return s.GetToken(SkelePOS, 0)
}

func (s *PosContext) Newln() INewlnContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INewlnContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INewlnContext)
}

func (s *PosContext) AllLn() []ILnContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ILnContext)(nil)).Elem())
	var tst = make([]ILnContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ILnContext)
		}
	}

	return tst
}

func (s *PosContext) Ln(i int) ILnContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILnContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ILnContext)
}

func (s *PosContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PosContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PosContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SkeleListener); ok {
		listenerT.EnterPos(s)
	}
}

func (s *PosContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SkeleListener); ok {
		listenerT.ExitPos(s)
	}
}

func (p *Skele) Pos() (localctx IPosContext) {
	localctx = NewPosContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, SkeleRULE_pos)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(81)
		p.Match(SkelePOS)
	}
	{
		p.SetState(82)
		p.Newln()
	}
	p.SetState(84)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == SkeleLINE {
		{
			p.SetState(83)
			p.Ln()
		}

		p.SetState(86)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ILnContext is an interface to support dynamic dispatch.
type ILnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLnContext differentiates from other interfaces.
	IsLnContext()
}

type LnContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLnContext() *LnContext {
	var p = new(LnContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SkeleRULE_ln
	return p
}

func (*LnContext) IsLnContext() {}

func NewLnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LnContext {
	var p = new(LnContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SkeleRULE_ln

	return p
}

func (s *LnContext) GetParser() antlr.Parser { return s.parser }

func (s *LnContext) LINE() antlr.TerminalNode {
	return s.GetToken(SkeleLINE, 0)
}

func (s *LnContext) Newln() INewlnContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INewlnContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INewlnContext)
}

func (s *LnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SkeleListener); ok {
		listenerT.EnterLn(s)
	}
}

func (s *LnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SkeleListener); ok {
		listenerT.ExitLn(s)
	}
}

func (p *Skele) Ln() (localctx ILnContext) {
	localctx = NewLnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, SkeleRULE_ln)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(88)
		p.Match(SkeleLINE)
	}
	{
		p.SetState(89)
		p.Newln()
	}

	return localctx
}

// INewlnContext is an interface to support dynamic dispatch.
type INewlnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNewlnContext differentiates from other interfaces.
	IsNewlnContext()
}

type NewlnContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNewlnContext() *NewlnContext {
	var p = new(NewlnContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SkeleRULE_newln
	return p
}

func (*NewlnContext) IsNewlnContext() {}

func NewNewlnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NewlnContext {
	var p = new(NewlnContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SkeleRULE_newln

	return p
}

func (s *NewlnContext) GetParser() antlr.Parser { return s.parser }

func (s *NewlnContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(SkeleNEWLINE)
}

func (s *NewlnContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(SkeleNEWLINE, i)
}

func (s *NewlnContext) AllComment() []ICommentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICommentContext)(nil)).Elem())
	var tst = make([]ICommentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICommentContext)
		}
	}

	return tst
}

func (s *NewlnContext) Comment(i int) ICommentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICommentContext)
}

func (s *NewlnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NewlnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NewlnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SkeleListener); ok {
		listenerT.EnterNewln(s)
	}
}

func (s *NewlnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SkeleListener); ok {
		listenerT.ExitNewln(s)
	}
}

func (p *Skele) Newln() (localctx INewlnContext) {
	localctx = NewNewlnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, SkeleRULE_newln)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(95)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == SkeleCOMMENT || _la == SkeleNEWLINE {
		p.SetState(92)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SkeleCOMMENT {
			{
				p.SetState(91)
				p.Comment()
			}

		}
		{
			p.SetState(94)
			p.Match(SkeleNEWLINE)
		}

		p.SetState(97)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ICommentContext is an interface to support dynamic dispatch.
type ICommentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCommentContext differentiates from other interfaces.
	IsCommentContext()
}

type CommentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommentContext() *CommentContext {
	var p = new(CommentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SkeleRULE_comment
	return p
}

func (*CommentContext) IsCommentContext() {}

func NewCommentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommentContext {
	var p = new(CommentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SkeleRULE_comment

	return p
}

func (s *CommentContext) GetParser() antlr.Parser { return s.parser }

func (s *CommentContext) COMMENT() antlr.TerminalNode {
	return s.GetToken(SkeleCOMMENT, 0)
}

func (s *CommentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SkeleListener); ok {
		listenerT.EnterComment(s)
	}
}

func (s *CommentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SkeleListener); ok {
		listenerT.ExitComment(s)
	}
}

func (p *Skele) Comment() (localctx ICommentContext) {
	localctx = NewCommentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, SkeleRULE_comment)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(99)
		p.Match(SkeleCOMMENT)
	}

	return localctx
}
