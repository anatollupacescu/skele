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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 18, 86, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 3, 2, 6, 2, 18, 10, 2, 13, 2, 14, 2, 19, 3, 2, 3, 2, 3, 3, 3,
	3, 6, 3, 26, 10, 3, 13, 3, 14, 3, 27, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 6, 5, 42, 10, 5, 13, 5, 14, 5, 43,
	3, 5, 7, 5, 47, 10, 5, 12, 5, 14, 5, 50, 11, 5, 3, 5, 3, 5, 3, 6, 3, 6,
	3, 6, 3, 6, 3, 6, 3, 6, 6, 6, 60, 10, 6, 13, 6, 14, 6, 61, 5, 6, 64, 10,
	6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 6, 7, 74, 10, 7, 13,
	7, 14, 7, 75, 5, 7, 78, 10, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8,
	2, 2, 9, 2, 4, 6, 8, 10, 12, 14, 2, 2, 2, 86, 2, 17, 3, 2, 2, 2, 4, 23,
	3, 2, 2, 2, 6, 29, 3, 2, 2, 2, 8, 35, 3, 2, 2, 2, 10, 63, 3, 2, 2, 2, 12,
	77, 3, 2, 2, 2, 14, 81, 3, 2, 2, 2, 16, 18, 5, 4, 3, 2, 17, 16, 3, 2, 2,
	2, 18, 19, 3, 2, 2, 2, 19, 17, 3, 2, 2, 2, 19, 20, 3, 2, 2, 2, 20, 21,
	3, 2, 2, 2, 21, 22, 7, 2, 2, 3, 22, 3, 3, 2, 2, 2, 23, 25, 5, 6, 4, 2,
	24, 26, 5, 8, 5, 2, 25, 24, 3, 2, 2, 2, 26, 27, 3, 2, 2, 2, 27, 25, 3,
	2, 2, 2, 27, 28, 3, 2, 2, 2, 28, 5, 3, 2, 2, 2, 29, 30, 7, 4, 2, 2, 30,
	31, 7, 11, 2, 2, 31, 32, 7, 14, 2, 2, 32, 33, 7, 10, 2, 2, 33, 34, 7, 12,
	2, 2, 34, 7, 3, 2, 2, 2, 35, 36, 7, 3, 2, 2, 36, 37, 7, 11, 2, 2, 37, 38,
	7, 9, 2, 2, 38, 39, 7, 13, 2, 2, 39, 41, 7, 8, 2, 2, 40, 42, 7, 12, 2,
	2, 41, 40, 3, 2, 2, 2, 42, 43, 3, 2, 2, 2, 43, 41, 3, 2, 2, 2, 43, 44,
	3, 2, 2, 2, 44, 48, 3, 2, 2, 2, 45, 47, 5, 10, 6, 2, 46, 45, 3, 2, 2, 2,
	47, 50, 3, 2, 2, 2, 48, 46, 3, 2, 2, 2, 48, 49, 3, 2, 2, 2, 49, 51, 3,
	2, 2, 2, 50, 48, 3, 2, 2, 2, 51, 52, 5, 12, 7, 2, 52, 9, 3, 2, 2, 2, 53,
	54, 7, 6, 2, 2, 54, 55, 7, 11, 2, 2, 55, 64, 7, 8, 2, 2, 56, 59, 7, 6,
	2, 2, 57, 58, 7, 12, 2, 2, 58, 60, 5, 14, 8, 2, 59, 57, 3, 2, 2, 2, 60,
	61, 3, 2, 2, 2, 61, 59, 3, 2, 2, 2, 61, 62, 3, 2, 2, 2, 62, 64, 3, 2, 2,
	2, 63, 53, 3, 2, 2, 2, 63, 56, 3, 2, 2, 2, 64, 65, 3, 2, 2, 2, 65, 66,
	7, 12, 2, 2, 66, 11, 3, 2, 2, 2, 67, 68, 7, 7, 2, 2, 68, 69, 7, 11, 2,
	2, 69, 78, 7, 8, 2, 2, 70, 73, 7, 7, 2, 2, 71, 72, 7, 12, 2, 2, 72, 74,
	5, 14, 8, 2, 73, 71, 3, 2, 2, 2, 74, 75, 3, 2, 2, 2, 75, 73, 3, 2, 2, 2,
	75, 76, 3, 2, 2, 2, 76, 78, 3, 2, 2, 2, 77, 67, 3, 2, 2, 2, 77, 70, 3,
	2, 2, 2, 78, 79, 3, 2, 2, 2, 79, 80, 7, 12, 2, 2, 80, 13, 3, 2, 2, 2, 81,
	82, 7, 16, 2, 2, 82, 83, 7, 11, 2, 2, 83, 84, 7, 8, 2, 2, 84, 15, 3, 2,
	2, 2, 10, 19, 27, 43, 48, 61, 63, 75, 77,
}
var literalNames = []string{
	"", "'fun'", "'file'", "'item'", "'pre'", "'pos'", "", "", "", "", "",
	"')'", "'/'", "'.'", "'\\'",
}
var symbolicNames = []string{
	"", "FUN", "FILE", "ITEM", "PRE", "POS", "COMMENT", "ARG_NAME", "FILENAME",
	"WORD", "NEWLINE", "CP", "FS", "DT", "ML", "WS", "ID",
}

var ruleNames = []string{
	"start", "spec", "file", "fun", "pre", "pos", "ml",
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
	SkeleFUN      = 1
	SkeleFILE     = 2
	SkeleITEM     = 3
	SkelePRE      = 4
	SkelePOS      = 5
	SkeleCOMMENT  = 6
	SkeleARG_NAME = 7
	SkeleFILENAME = 8
	SkeleWORD     = 9
	SkeleNEWLINE  = 10
	SkeleCP       = 11
	SkeleFS       = 12
	SkeleDT       = 13
	SkeleML       = 14
	SkeleWS       = 15
	SkeleID       = 16
)

// Skele rules.
const (
	SkeleRULE_start = 0
	SkeleRULE_spec  = 1
	SkeleRULE_file  = 2
	SkeleRULE_fun   = 3
	SkeleRULE_pre   = 4
	SkeleRULE_pos   = 5
	SkeleRULE_ml    = 6
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
	p.SetState(15)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == SkeleFILE {
		{
			p.SetState(14)
			p.Spec()
		}

		p.SetState(17)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(19)
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

func (s *SpecContext) File() IFileContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFileContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFileContext)
}

func (s *SpecContext) AllFun() []IFunContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFunContext)(nil)).Elem())
	var tst = make([]IFunContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFunContext)
		}
	}

	return tst
}

func (s *SpecContext) Fun(i int) IFunContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFunContext)
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
		p.SetState(21)
		p.File()
	}
	p.SetState(23)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == SkeleFUN {
		{
			p.SetState(22)
			p.Fun()
		}

		p.SetState(25)
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

func (s *FileContext) WORD() antlr.TerminalNode {
	return s.GetToken(SkeleWORD, 0)
}

func (s *FileContext) FS() antlr.TerminalNode {
	return s.GetToken(SkeleFS, 0)
}

func (s *FileContext) FILENAME() antlr.TerminalNode {
	return s.GetToken(SkeleFILENAME, 0)
}

func (s *FileContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(SkeleNEWLINE, 0)
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
	p.EnterRule(localctx, 4, SkeleRULE_file)

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
		p.SetState(27)
		p.Match(SkeleFILE)
	}
	{
		p.SetState(28)
		p.Match(SkeleWORD)
	}
	{
		p.SetState(29)
		p.Match(SkeleFS)
	}
	{
		p.SetState(30)
		p.Match(SkeleFILENAME)
	}
	{
		p.SetState(31)
		p.Match(SkeleNEWLINE)
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

func (s *FunContext) WORD() antlr.TerminalNode {
	return s.GetToken(SkeleWORD, 0)
}

func (s *FunContext) ARG_NAME() antlr.TerminalNode {
	return s.GetToken(SkeleARG_NAME, 0)
}

func (s *FunContext) CP() antlr.TerminalNode {
	return s.GetToken(SkeleCP, 0)
}

func (s *FunContext) COMMENT() antlr.TerminalNode {
	return s.GetToken(SkeleCOMMENT, 0)
}

func (s *FunContext) Pos() IPosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPosContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPosContext)
}

func (s *FunContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(SkeleNEWLINE)
}

func (s *FunContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(SkeleNEWLINE, i)
}

func (s *FunContext) AllPre() []IPreContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPreContext)(nil)).Elem())
	var tst = make([]IPreContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPreContext)
		}
	}

	return tst
}

func (s *FunContext) Pre(i int) IPreContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPreContext)(nil)).Elem(), i)

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
	p.EnterRule(localctx, 6, SkeleRULE_fun)
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
		p.SetState(33)
		p.Match(SkeleFUN)
	}
	{
		p.SetState(34)
		p.Match(SkeleWORD)
	}
	{
		p.SetState(35)
		p.Match(SkeleARG_NAME)
	}
	{
		p.SetState(36)
		p.Match(SkeleCP)
	}
	{
		p.SetState(37)
		p.Match(SkeleCOMMENT)
	}
	p.SetState(39)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == SkeleNEWLINE {
		{
			p.SetState(38)
			p.Match(SkeleNEWLINE)
		}

		p.SetState(41)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(46)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SkelePRE {
		{
			p.SetState(43)
			p.Pre()
		}

		p.SetState(48)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(49)
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

func (s *PreContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(SkeleNEWLINE)
}

func (s *PreContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(SkeleNEWLINE, i)
}

func (s *PreContext) PRE() antlr.TerminalNode {
	return s.GetToken(SkelePRE, 0)
}

func (s *PreContext) WORD() antlr.TerminalNode {
	return s.GetToken(SkeleWORD, 0)
}

func (s *PreContext) COMMENT() antlr.TerminalNode {
	return s.GetToken(SkeleCOMMENT, 0)
}

func (s *PreContext) AllMl() []IMlContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMlContext)(nil)).Elem())
	var tst = make([]IMlContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMlContext)
		}
	}

	return tst
}

func (s *PreContext) Ml(i int) IMlContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMlContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMlContext)
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
	p.EnterRule(localctx, 8, SkeleRULE_pre)

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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(61)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(51)
			p.Match(SkelePRE)
		}
		{
			p.SetState(52)
			p.Match(SkeleWORD)
		}
		{
			p.SetState(53)
			p.Match(SkeleCOMMENT)
		}

	case 2:
		{
			p.SetState(54)
			p.Match(SkelePRE)
		}
		p.SetState(57)
		p.GetErrorHandler().Sync(p)
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(55)
					p.Match(SkeleNEWLINE)
				}
				{
					p.SetState(56)
					p.Ml()
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(59)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())
		}

	}
	{
		p.SetState(63)
		p.Match(SkeleNEWLINE)
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

func (s *PosContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(SkeleNEWLINE)
}

func (s *PosContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(SkeleNEWLINE, i)
}

func (s *PosContext) POS() antlr.TerminalNode {
	return s.GetToken(SkelePOS, 0)
}

func (s *PosContext) WORD() antlr.TerminalNode {
	return s.GetToken(SkeleWORD, 0)
}

func (s *PosContext) COMMENT() antlr.TerminalNode {
	return s.GetToken(SkeleCOMMENT, 0)
}

func (s *PosContext) AllMl() []IMlContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMlContext)(nil)).Elem())
	var tst = make([]IMlContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMlContext)
		}
	}

	return tst
}

func (s *PosContext) Ml(i int) IMlContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMlContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMlContext)
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
	p.EnterRule(localctx, 10, SkeleRULE_pos)

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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(75)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(65)
			p.Match(SkelePOS)
		}
		{
			p.SetState(66)
			p.Match(SkeleWORD)
		}
		{
			p.SetState(67)
			p.Match(SkeleCOMMENT)
		}

	case 2:
		{
			p.SetState(68)
			p.Match(SkelePOS)
		}
		p.SetState(71)
		p.GetErrorHandler().Sync(p)
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(69)
					p.Match(SkeleNEWLINE)
				}
				{
					p.SetState(70)
					p.Ml()
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(73)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())
		}

	}
	{
		p.SetState(77)
		p.Match(SkeleNEWLINE)
	}

	return localctx
}

// IMlContext is an interface to support dynamic dispatch.
type IMlContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMlContext differentiates from other interfaces.
	IsMlContext()
}

type MlContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMlContext() *MlContext {
	var p = new(MlContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SkeleRULE_ml
	return p
}

func (*MlContext) IsMlContext() {}

func NewMlContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MlContext {
	var p = new(MlContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SkeleRULE_ml

	return p
}

func (s *MlContext) GetParser() antlr.Parser { return s.parser }

func (s *MlContext) ML() antlr.TerminalNode {
	return s.GetToken(SkeleML, 0)
}

func (s *MlContext) WORD() antlr.TerminalNode {
	return s.GetToken(SkeleWORD, 0)
}

func (s *MlContext) COMMENT() antlr.TerminalNode {
	return s.GetToken(SkeleCOMMENT, 0)
}

func (s *MlContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MlContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MlContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SkeleListener); ok {
		listenerT.EnterMl(s)
	}
}

func (s *MlContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SkeleListener); ok {
		listenerT.ExitMl(s)
	}
}

func (p *Skele) Ml() (localctx IMlContext) {
	localctx = NewMlContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, SkeleRULE_ml)

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
		p.SetState(79)
		p.Match(SkeleML)
	}
	{
		p.SetState(80)
		p.Match(SkeleWORD)
	}
	{
		p.SetState(81)
		p.Match(SkeleCOMMENT)
	}

	return localctx
}
