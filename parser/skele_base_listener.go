// Code generated from Skele.g4 by ANTLR 4.9.1. DO NOT EDIT.

package parser // Skele

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseSkeleListener is a complete listener for a parse tree produced by Skele.
type BaseSkeleListener struct{}

var _ SkeleListener = &BaseSkeleListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSkeleListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSkeleListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSkeleListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSkeleListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStart is called when production start is entered.
func (s *BaseSkeleListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BaseSkeleListener) ExitStart(ctx *StartContext) {}

// EnterSpec is called when production spec is entered.
func (s *BaseSkeleListener) EnterSpec(ctx *SpecContext) {}

// ExitSpec is called when production spec is exited.
func (s *BaseSkeleListener) ExitSpec(ctx *SpecContext) {}

// EnterFol is called when production fol is entered.
func (s *BaseSkeleListener) EnterFol(ctx *FolContext) {}

// ExitFol is called when production fol is exited.
func (s *BaseSkeleListener) ExitFol(ctx *FolContext) {}

// EnterPkg is called when production pkg is entered.
func (s *BaseSkeleListener) EnterPkg(ctx *PkgContext) {}

// ExitPkg is called when production pkg is exited.
func (s *BaseSkeleListener) ExitPkg(ctx *PkgContext) {}

// EnterDoc is called when production doc is entered.
func (s *BaseSkeleListener) EnterDoc(ctx *DocContext) {}

// ExitDoc is called when production doc is exited.
func (s *BaseSkeleListener) ExitDoc(ctx *DocContext) {}

// EnterFile is called when production file is entered.
func (s *BaseSkeleListener) EnterFile(ctx *FileContext) {}

// ExitFile is called when production file is exited.
func (s *BaseSkeleListener) ExitFile(ctx *FileContext) {}

// EnterFun is called when production fun is entered.
func (s *BaseSkeleListener) EnterFun(ctx *FunContext) {}

// ExitFun is called when production fun is exited.
func (s *BaseSkeleListener) ExitFun(ctx *FunContext) {}

// EnterPre is called when production pre is entered.
func (s *BaseSkeleListener) EnterPre(ctx *PreContext) {}

// ExitPre is called when production pre is exited.
func (s *BaseSkeleListener) ExitPre(ctx *PreContext) {}

// EnterPos is called when production pos is entered.
func (s *BaseSkeleListener) EnterPos(ctx *PosContext) {}

// ExitPos is called when production pos is exited.
func (s *BaseSkeleListener) ExitPos(ctx *PosContext) {}

// EnterLn is called when production ln is entered.
func (s *BaseSkeleListener) EnterLn(ctx *LnContext) {}

// ExitLn is called when production ln is exited.
func (s *BaseSkeleListener) ExitLn(ctx *LnContext) {}

// EnterNewln is called when production newln is entered.
func (s *BaseSkeleListener) EnterNewln(ctx *NewlnContext) {}

// ExitNewln is called when production newln is exited.
func (s *BaseSkeleListener) ExitNewln(ctx *NewlnContext) {}

// EnterComment is called when production comment is entered.
func (s *BaseSkeleListener) EnterComment(ctx *CommentContext) {}

// ExitComment is called when production comment is exited.
func (s *BaseSkeleListener) ExitComment(ctx *CommentContext) {}
