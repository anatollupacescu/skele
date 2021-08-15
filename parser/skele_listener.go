// Code generated from Skele.g4 by ANTLR 4.9.1. DO NOT EDIT.

package parser // Skele

import "github.com/antlr/antlr4/runtime/Go/antlr"

// SkeleListener is a complete listener for a parse tree produced by Skele.
type SkeleListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterSpec is called when entering the spec production.
	EnterSpec(c *SpecContext)

	// EnterFol is called when entering the fol production.
	EnterFol(c *FolContext)

	// EnterPkg is called when entering the pkg production.
	EnterPkg(c *PkgContext)

	// EnterDoc is called when entering the doc production.
	EnterDoc(c *DocContext)

	// EnterFile is called when entering the file production.
	EnterFile(c *FileContext)

	// EnterFun is called when entering the fun production.
	EnterFun(c *FunContext)

	// EnterPre is called when entering the pre production.
	EnterPre(c *PreContext)

	// EnterPos is called when entering the pos production.
	EnterPos(c *PosContext)

	// EnterLn is called when entering the ln production.
	EnterLn(c *LnContext)

	// EnterNewln is called when entering the newln production.
	EnterNewln(c *NewlnContext)

	// EnterComment is called when entering the comment production.
	EnterComment(c *CommentContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitSpec is called when exiting the spec production.
	ExitSpec(c *SpecContext)

	// ExitFol is called when exiting the fol production.
	ExitFol(c *FolContext)

	// ExitPkg is called when exiting the pkg production.
	ExitPkg(c *PkgContext)

	// ExitDoc is called when exiting the doc production.
	ExitDoc(c *DocContext)

	// ExitFile is called when exiting the file production.
	ExitFile(c *FileContext)

	// ExitFun is called when exiting the fun production.
	ExitFun(c *FunContext)

	// ExitPre is called when exiting the pre production.
	ExitPre(c *PreContext)

	// ExitPos is called when exiting the pos production.
	ExitPos(c *PosContext)

	// ExitLn is called when exiting the ln production.
	ExitLn(c *LnContext)

	// ExitNewln is called when exiting the newln production.
	ExitNewln(c *NewlnContext)

	// ExitComment is called when exiting the comment production.
	ExitComment(c *CommentContext)
}
