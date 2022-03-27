package main

import (
	"log"
	"regexp"
	"strings"

	"github.com/anatollupacescu/skele/parser"
)

type machine struct {
	pkgs []pkg
}

func (m *machine) currentPackage() *pkg {
	lastPkgIndex := len(m.pkgs) - 1
	return &m.pkgs[lastPkgIndex]
}

type pkg struct {
	name, fol string
	doc       []string
	file      []file
	fsm       []fsm
}

func (p *pkg) currentFile() *file {
	lastFileIndex := len(p.file) - 1
	return &p.file[lastFileIndex]
}

type fsm struct {
	name   string
	states []string
}

type file struct {
	name string
	fun  []fun
}

func (f *file) currentFun() *fun {
	lastFunIndex := len(f.fun) - 1
	return &f.fun[lastFunIndex]
}

type edge struct {
	name, state string
}

type fun struct {
	name     string
	fsm      []edge
	pre, pos []prepos
}

func (f *fun) currentPre() *prepos {
	lastPreIndex := len(f.pre) - 1
	return &f.pre[lastPreIndex]
}

func (f *fun) currentPos() *prepos {
	lastPosIndex := len(f.pos) - 1
	return &f.pos[lastPosIndex]
}

func (f *fun) condition(from, to string) {
	f.fsm = append(f.fsm, edge{
		name:  from,
		state: to,
	})
}

type prepos struct {
	fsm        []edge
	tcE        []string
	tcS        []string
	succ, fail string
}

func (p *prepos) condition(from, to string) { // it is `condition` for `pre`
	p.transition(from, to) // but `transition` for `pos`
}

func (p *prepos) errTestCase(text string) {
	p.tcE = append(p.tcE, text)
}

func (p *prepos) okTestCase(text string) {
	p.tcS = append(p.tcS, text)
}

func (p *prepos) transition(from, to string) {
	p.fsm = append(p.fsm, edge{
		name:  from,
		state: to,
	})
}

type trimmable string

func (ts trimmable) trim() string {
	out := strings.TrimLeft(string(ts), "\\")
	out = strings.Trim(out, " ")
	return out
}

type skeleListener struct {
	*machine
	*parser.BaseSkeleListener
	sink    func(trimmable)
	fsmSink func(string, string)
	tcSink  func(trimmable)
}

func (c *skeleListener) EnterPkg(ctx *parser.PkgContext) {
	c.machine.pkgs = append(c.machine.pkgs, pkg{
		name: ctx.WORD().GetText(),
	})
}

func (c *skeleListener) EnterFol(ctx *parser.FolContext) {
	cs := c.currentPackage()
	if word := ctx.WORD(); word != nil {
		cs.fol = word.GetText()
	}
}

func (c *skeleListener) EnterDoc(ctx *parser.DocContext) {
	cs := c.currentPackage()
	c.sink = func(in trimmable) {
		cs.doc = append(cs.doc, in.trim())
	}
}

func (c *skeleListener) EnterFsm(ctx *parser.FsmContext) {
	cs := c.currentPackage()
	c.sink = func(in trimmable) {
		cs.fsm = append(cs.fsm, fsm{name: in.trim()})
	}
}

func (c *skeleListener) EnterSts(ctx *parser.StsContext) {
	currentPkg := c.machine.currentPackage()
	fsmIndex := len(currentPkg.fsm) - 1
	states := &currentPkg.fsm[fsmIndex].states
	c.sink = func(in trimmable) {
		*states = append(*states, in.trim())
	}
}

func (c *skeleListener) EnterFile(ctx *parser.FileContext) {
	cs := c.currentPackage()
	f := file{
		name: ctx.FILENAME().GetText(),
	}

	if containsFile(cs.file, f) {
		log.Fatal("duplicate file: ", f.name)
	}

	cs.file = append(cs.file, f)
}

func containsFile(ff []file, f file) bool {
	for _, ex := range ff {
		if ex.name == f.name {
			return true
		}
	}
	return false
}

func (c *skeleListener) EnterFun(*parser.FunContext) {
	cf := c.machine.currentPackage().currentFile()

	c.sink = func(in trimmable) {
		f := fun{
			name: in.trim(),
		}

		if containsFun(cf.fun, f) {
			log.Fatal("duplicate fun:", in)
		}

		cf.fun = append(cf.fun, f)
	}

	c.fsmSink = func(name, state string) {
		fun := cf.currentFun()
		fun.condition(name, state)
	}
}

func containsFun(list []fun, f fun) bool {
	for _, ex := range list {
		if ex.name == f.name {
			return true
		}
	}
	return false
}

func (c *skeleListener) EnterPre(ctx *parser.PreContext) {
	cf := c.machine.currentPackage().currentFile().currentFun()
	adapter := func(given, assert string) {
		pp := prepos{
			succ: given,
			fail: assert,
		}

		cf.pre = append(cf.pre, pp)
	}

	c.sink = func(in trimmable) {
		parsePreposLine(in.trim(), adapter)
	}

	c.fsmSink = func(name, state string) {
		pre := cf.currentPre()
		pre.condition(name, state)
	}

	c.tcSink = func(s trimmable) {
		ts := s.trim()
		pre := cf.currentPre()
		if !strings.Contains(ts, "assert") {
			ts += ", assert error"
		}
		pre.errTestCase(ts)
	}
}

func (c *skeleListener) EnterPos(ctx *parser.PosContext) {
	adapter := func(given, assert string) {
		pp := prepos{
			succ: given,
			fail: assert,
		}

		fn := c.machine.currentPackage().currentFile().currentFun()
		fn.pos = append(fn.pos, pp)
	}

	c.sink = func(in trimmable) {
		parsePreposLine(in.trim(), adapter)
	}

	c.tcSink = func(s trimmable) {
		ts := s.trim()
		pos := c.machine.currentPackage().currentFile().currentFun().currentPos()
		if isSuccess(pos.succ, ts) {
			pos.okTestCase(ts)
			return
		}

		pos.errTestCase(ts)
	}
}

func isSuccess(pos, s string) bool {
	if strings.HasSuffix(s, pos) ||
		strings.HasSuffix(s, "assert success") ||
		strings.HasSuffix(s, "assert ok") {
		return true
	}
	return false
}

func (c *skeleListener) EnterLn(ctx *parser.LnContext) {
	line := ctx.LINE().GetText()
	c.sink(trimmable(line))
}

var (
	fsmReg   = regexp.MustCompile(`fsm\{(?P<name>\w+)((?P<op>(=))|(?P<arr>(->)))(?P<state>\w+)\}`)
	tcReg    = regexp.MustCompile(`tc\{(?P<text>(\w+(,* +\w+)*))\}`)
	fsmNames = fsmReg.SubexpNames()
	tcNames  = tcReg.SubexpNames()
)

func (c *skeleListener) EnterComment(ctx *parser.CommentContext) {
	commentText := ctx.COMMENT().GetText()

	for _, submatches := range tcReg.FindAllStringSubmatch(commentText, -1) {
		results := make(map[string]string, len(submatches))
		for i, name := range submatches {
			results[tcNames[i]] = name
		}
		txt := results["text"]
		c.tcSink(trimmable(txt))
	}

	// search for `fsm` declaration in the comment text
	for _, submatches := range fsmReg.FindAllStringSubmatch(commentText, -1) {
		results := make(map[string]string, len(submatches))
		for i, name := range submatches {
			results[fsmNames[i]] = name
		}

		name, state := results["name"], results["state"]

		if results["op"] == "=" {
			c.fsmSink(name, state)
			continue
		}

		// if op is `->` means it's a post effect
		pos := c.machine.currentPackage().currentFile().currentFun().currentPos()
		pos.transition(name, state)
	}
}

func parsePreposLine(in string, sink func(string, string)) {
	seg := strings.Split(in, "\\")

	var lines []string
	for _, s := range seg {
		if s == "" {
			continue
		}
		lines = append(lines, strings.Trim(s, " "))
	}

	var success, fail string

	success = lines[0]

	if len(lines) > 1 { // the technical failure precondition is optional
		fail = lines[1]
	}

	sink(success, fail)
}
