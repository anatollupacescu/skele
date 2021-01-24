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

	// EnterFile is called when entering the file production.
	EnterFile(c *FileContext)

	// EnterFun is called when entering the fun production.
	EnterFun(c *FunContext)

	// EnterPre is called when entering the pre production.
	EnterPre(c *PreContext)

	// EnterPos is called when entering the pos production.
	EnterPos(c *PosContext)

	// EnterMl is called when entering the ml production.
	EnterMl(c *MlContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitSpec is called when exiting the spec production.
	ExitSpec(c *SpecContext)

	// ExitFile is called when exiting the file production.
	ExitFile(c *FileContext)

	// ExitFun is called when exiting the fun production.
	ExitFun(c *FunContext)

	// ExitPre is called when exiting the pre production.
	ExitPre(c *PreContext)

	// ExitPos is called when exiting the pos production.
	ExitPos(c *PosContext)

	// ExitMl is called when exiting the ml production.
	ExitMl(c *MlContext)
}
