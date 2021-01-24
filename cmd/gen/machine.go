package main

import (
	"strings"

	"github.com/spirit/parser"
)

type machine struct {
	specs []spec
}

type spec struct {
	pkg, filename string
	fun           []fun
}

type fun struct {
	name, arg, comm string
	pre, pos        []prepos
}

type prepos struct {
	name, comm string
}

type skeleListener struct {
	*machine
	adapter func(n, c string)
	*parser.BaseSkeleListener
}

func (c *skeleListener) EnterFile(ctx *parser.FileContext) {
	var cs spec
	cs.pkg = ctx.WORD().GetText()
	cs.filename = ctx.FILENAME().GetText()
	c.machine.specs = append(c.machine.specs, cs)
}

func (c *skeleListener) EnterFun(ctx *parser.FunContext) {
	arg := strings.Trim(ctx.ARG_NAME().GetText()[1:], " ")
	comm := ctx.COMMENT().GetText()[1:]
	comm = strings.Trim(comm, " ")

	var newFun fun

	newFun.name = ctx.WORD().GetText()
	newFun.arg = arg
	newFun.comm = comm

	specIndex := len(c.machine.specs) - 1
	cs := &c.machine.specs[specIndex]
	cs.fun = append(cs.fun, newFun)
}

func (c *skeleListener) EnterPre(ctx *parser.PreContext) {
	adapter := func(name, comm string) {
		var pp prepos
		pp.name = name
		pp.comm = comm
		specIndex := len(c.machine.specs) - 1
		funIndex := len(c.machine.specs[specIndex].fun) - 1
		fn := &c.machine.specs[specIndex].fun[funIndex]
		fn.pre = append(fn.pre, pp)
	}
	if ctx.WORD() == nil {
		c.adapter = adapter
		return // multiline
	}
	name := ctx.WORD().GetText()
	comm := ctx.COMMENT().GetText()

	adapter(name, comm)
}

func (c *skeleListener) EnterPos(ctx *parser.PosContext) {
	adapter := func(name, comm string) {
		var pp prepos
		pp.name = name
		pp.comm = comm
		specIndex := len(c.machine.specs) - 1
		funIndex := len(c.machine.specs[specIndex].fun) - 1
		fn := &c.machine.specs[specIndex].fun[funIndex]
		fn.pos = append(fn.pos, pp)
	}
	if ctx.WORD() == nil {
		c.adapter = adapter
		return // multiline
	}
	name := ctx.WORD().GetText()
	comm := ctx.COMMENT().GetText()

	adapter(name, comm)
}

func (c *skeleListener) EnterMl(ctx *parser.MlContext) {
	name := ctx.WORD().GetText()
	comm := strings.TrimLeft(ctx.COMMENT().GetText(), "#")
	comm = strings.Trim(comm, " ")

	c.adapter(name, comm)
}
