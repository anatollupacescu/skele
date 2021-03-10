package main

import (
	"strings"

	"github.com/anatollupacescu/skele/parser"
)

type machine struct {
	specs []spec
}

type spec struct {
	pkg, fol string
	doc      []string
	file     []file
}

type file struct {
	name string
	fun  []fun
}

type fun struct {
	name     string
	pre, pos []prepos
}

type prepos struct {
	domain, impl string
}

type skeleListener struct {
	*machine
	*parser.BaseSkeleListener
	sink func(string)
}

func (c *skeleListener) EnterPkg(ctx *parser.PkgContext) {
	c.machine.specs = append(c.machine.specs, spec{
		pkg: ctx.WORD().GetText(),
	})
}

func (c *skeleListener) EnterFol(ctx *parser.FolContext) {
	specIndex := len(c.machine.specs) - 1
	cs := &c.machine.specs[specIndex]
	if word := ctx.WORD(); word != nil {
		cs.fol = word.GetText()
	}
}

func (c *skeleListener) EnterDoc(ctx *parser.DocContext) {
	specIndex := len(c.machine.specs) - 1
	cs := &c.machine.specs[specIndex]
	c.sink = func(in string) {
		in = strings.TrimLeft(in, "\\")
		in = strings.Trim(in, " ")
		cs.doc = append(cs.doc, in)
	}
}

func (c *skeleListener) EnterLn(ctx *parser.LnContext) {
	line := ctx.LINE().GetText()
	c.sink(line)
}

func (c *skeleListener) EnterFile(ctx *parser.FileContext) {
	specIndex := len(c.machine.specs) - 1
	cs := &c.machine.specs[specIndex]
	f := file{
		name: ctx.FILENAME().GetText(),
	}
	cs.file = append(cs.file, f)
}

func (c *skeleListener) EnterFun(*parser.FunContext) {
	c.sink = func(in string) {
		in = strings.TrimLeft(in, "\\")
		in = strings.Trim(in, " ")
		f := fun{
			name: in,
		}

		specIndex := len(c.machine.specs) - 1
		cs := &c.machine.specs[specIndex]

		fileIndex := len(cs.file) - 1
		cf := &cs.file[fileIndex]

		cf.fun = append(cf.fun, f)
	}
}

func (c *skeleListener) EnterPre(ctx *parser.PreContext) {
	adapter := func(given, assert string) {
		var pp prepos
		pp.domain = given
		pp.impl = assert

		specIndex := len(c.machine.specs) - 1
		cs := &c.machine.specs[specIndex]

		fileIndex := len(cs.file) - 1
		cf := &cs.file[fileIndex]

		funIndex := len(cf.fun) - 1
		fn := &cf.fun[funIndex]

		fn.pre = append(fn.pre, pp)
	}

	c.sink = func(s string) {
		parsePreposLine(s, adapter)
	}
}

func (c *skeleListener) EnterPos(ctx *parser.PosContext) {
	adapter := func(given, assert string) {
		var pp = prepos{
			domain: given,
			impl:   assert,
		}

		specIndex := len(c.machine.specs) - 1
		cs := &c.machine.specs[specIndex]

		fileIndex := len(cs.file) - 1
		cf := &cs.file[fileIndex]

		funIndex := len(cf.fun) - 1
		fn := &cf.fun[funIndex]

		fn.pos = append(fn.pos, pp)
	}

	c.sink = func(s string) {
		parsePreposLine(s, adapter)
	}
}

func parsePreposLine(in string, sink func(_, _ string)) {
	seg := strings.Split(in, "\\")

	var lines []string
	for _, s := range seg {
		if s != "" {
			lines = append(lines, strings.Trim(s, " "))
		}
	}

	var domain, impl string

	domain = lines[0]

	if len(lines) > 1 {
		impl = lines[1]
	}

	sink(domain, impl)
}
