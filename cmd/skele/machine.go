package main

import (
	"errors"
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/anatollupacescu/skele/parser"
)

type machine struct {
	pkgs []pkg
}

/*
considerations
- first declared state is the initial state
- one initial state
- one transition per edge
- one end state

build phase
- fsm names should be unique globally
- fsm in function name and precondition should be equality check
- fsm in postcondition should be a transition

validation phase
- state without connection to the initial state
- missing transition function for a certain edge
- two functions for the same edge (postcondition)
- error if fsm name not mentioned in function name and state is not initial
*/
func (m *machine) ValidateFSM() (errs []error) {
	global := make(mFSMS)

	for _, pkg := range m.pkgs {
		for _, f := range pkg.fsm {
			if err := global.add(pkg.name, f); err != nil {
				errs = append(errs, err)
			}
		}
	}

	if len(errs) > 0 {
		return
	}

	// register transitions for every function
	for _, pkg := range m.pkgs {
		for _, f := range pkg.file {
			for _, fun := range f.fun {
				errs = append(errs, global.fun(fun)...)
			}
		}
	}

	if len(errs) > 0 {
		return
	}

	// validate
	errs = append(errs, global.checkUnused()...)
	// errs = append(errs, global.checkDuplicateTransition()...)
	// errs = append(errs, global.checkCirclular()...)

	return
}

func (m *machine) currentPackage() *pkg {
	return &m.pkgs[len(m.pkgs)-1]
}

type pkg struct {
	name, fol string
	doc       []string
	file      []file
	fsm       []fsm
}

func (p *pkg) currentFile() *file {
	return &p.file[len(p.file)-1]
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
	return &f.fun[len(f.fun)-1]
}

type edge struct {
	name, state string
}

type fun struct {
	name     string
	fsm      *edge //one or more?
	pre, pos []prepos
}

func (f *fun) currentPre() *prepos {
	return &f.pre[len(f.pre)-1]
}

func (f *fun) currentPos() *prepos {
	return &f.pos[len(f.pos)-1]
}

func (f *fun) from(name, to string) {
	f.fsm = &edge{name: name, state: to}
}

type prepos struct {
	fsm        edge
	tcF, tcS   []string
	arnF, arnS []string
	succ, fail string
}

func (p *prepos) condition(from, to string) {
	p.edge(from, to)
}

func (p *prepos) transition(from, to string) {
	p.edge(from, to)
}

func (p *prepos) edge(from, to string) {
	p.fsm = edge{
		name:  from,
		state: to,
	}
}

func (p *prepos) failTestCase(text string) {
	p.tcF = append(p.tcF, text)
}

func (p *prepos) okTestCase(text string) {
	p.tcS = append(p.tcS, text)
}

func (p *prepos) failAssertion(text string) {
	p.arnF = append(p.arnF, text)
}

func (p *prepos) okAssertion(text string) {
	p.arnS = append(p.arnS, text)
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
	errors []error

	lnSink func(trimmable)

	fsmSink func(string, string, string)
	arnSink func(string)
	tcsSink func(string)
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
	c.lnSink = func(in trimmable) {
		cs.doc = append(cs.doc, in.trim())
	}
}

func (c *skeleListener) EnterFsm(ctx *parser.FsmContext) {
	cs := c.currentPackage()
	c.lnSink = func(in trimmable) {
		cs.fsm = append(cs.fsm, fsm{name: in.trim()})
	}
}

func (c *skeleListener) EnterSts(ctx *parser.StsContext) {
	currentPkg := c.machine.currentPackage()
	fsmIndex := len(currentPkg.fsm) - 1
	states := &currentPkg.fsm[fsmIndex].states
	c.lnSink = func(in trimmable) {
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

var ErrInvalidOperator = errors.New("invalid operator")

func (c *skeleListener) EnterFun(*parser.FunContext) {
	cf := c.machine.currentPackage().currentFile()

	c.lnSink = func(in trimmable) {
		f := fun{
			name: in.trim(),
		}

		if containsFun(cf.fun, f) {
			log.Fatal("duplicate fun:", in)
		}

		cf.fun = append(cf.fun, f)
	}

	c.fsmSink = func(name, state, op string) {
		if op != "=" {
			c.errors = append(c.errors, fmt.Errorf("fun: %s, fsm: %s %s, op '%s': %w", cf.name, name, state, op, ErrInvalidOperator))
			return
		}
		fun := cf.currentFun()
		fun.from(name, state)
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

const (
	failSign = "> "
	okSign   = "< "
	bothSign = "<> "
)

func (c *skeleListener) EnterPre(ctx *parser.PreContext) {
	cf := c.machine.currentPackage().currentFile().currentFun()
	adapter := func(given, assert string) {
		pp := prepos{
			succ: given,
			fail: assert,
		}

		cf.pre = append(cf.pre, pp)
	}

	c.lnSink = func(in trimmable) {
		parsePreposLine(in.trim(), adapter)
	}

	c.fsmSink = func(name, state, op string) {
		if op != "=" {
			c.errors = append(c.errors, fmt.Errorf("fun: '%s', fsm: %s %s, op '%s': %w", cf.name, name, state, op, ErrInvalidOperator))
			return
		}
		cf.currentPre().condition(name, state)
	}

	c.tcsSink = func(s string) {
		pre := cf.currentPre()
		if strings.HasPrefix(s, failSign) {
			pre.failTestCase(s[len(failSign):])
			return
		}
		s = strings.TrimPrefix(s, okSign)
		pre.okTestCase(s)
	}

	c.arnSink = func(s string) {
		pre := cf.currentPre()
		if strings.HasPrefix(s, bothSign) {
			pre.failAssertion(s[len(bothSign):])
			pre.okAssertion(s[len(bothSign):])
			return
		}
		if strings.HasPrefix(s, failSign) {
			pre.failAssertion(s[len(failSign):])
			return
		}
		s = strings.TrimPrefix(s, okSign)
		pre.okAssertion(s)
	}
}

func (c *skeleListener) EnterPos(ctx *parser.PosContext) {
	cf := c.machine.currentPackage().currentFile().currentFun()
	adapter := func(given, assert string) {
		pp := prepos{
			succ: given,
			fail: assert,
		}

		fn := cf
		fn.pos = append(fn.pos, pp)
	}

	c.lnSink = func(in trimmable) {
		parsePreposLine(in.trim(), adapter)
	}

	c.fsmSink = func(name, state, op string) {
		if op != "->" {
			c.errors = append(c.errors, fmt.Errorf("fun: '%s', fsm: %s %s, op '%s': %w", cf.name, name, state, op, ErrInvalidOperator))
			return
		}
		cf.currentPos().transition(name, state)
	}

	c.tcsSink = func(s string) {
		pos := cf.currentPos()

		if strings.HasPrefix(s, failSign) {
			pos.failTestCase(s[len(failSign):])
			return
		}
		s = strings.TrimPrefix(s, okSign)
		pos.okTestCase(s)
	}

	c.arnSink = func(s string) {
		pos := cf.currentPos()
		if strings.HasPrefix(s, bothSign) {
			pos.failAssertion(s[len(bothSign):])
			pos.okAssertion(s[len(bothSign):])
			return
		}
		if strings.HasPrefix(s, failSign) {
			pos.failAssertion(s[len(failSign):])
			return
		}
		s = strings.TrimPrefix(s, okSign)
		pos.okAssertion(s)
	}
}

func (c *skeleListener) EnterLn(ctx *parser.LnContext) {
	line := ctx.LINE().GetText()
	c.lnSink(trimmable(line))
}

var (
	fsmReg = regexp.MustCompile(`fsm\{(?P<name>\w+)((?P<op>(=|->)))(?P<state>\w+)\}`)
	arnReg = regexp.MustCompile(`arn\{(?P<text>([<>] |<> )?(\w+(,* +\w+)*))\}`)
	tcsReg = regexp.MustCompile(`tcs\{(?P<text>([<>] )?(\w+(,* +\w+)*))\}`)

	fsmNames = fsmReg.SubexpNames()
	tcsNames = tcsReg.SubexpNames()
	arnNames = arnReg.SubexpNames()
)

func (c *skeleListener) EnterComment(ctx *parser.CommentContext) {
	commentText := ctx.COMMENT().GetText()

	// search for `arn` declaration in the comment text
	for _, submatches := range arnReg.FindAllStringSubmatch(commentText, -1) {
		results := make(map[string]string, len(submatches))
		for i, name := range submatches {
			results[arnNames[i]] = name
		}
		txt := results["text"]
		c.arnSink(txt)
	}

	// search for `tcs` declaration in the comment text
	for _, submatches := range tcsReg.FindAllStringSubmatch(commentText, -1) {
		results := make(map[string]string, len(submatches))
		for i, name := range submatches {
			results[tcsNames[i]] = name
		}
		txt := results["text"]
		c.tcsSink(txt)
	}

	// search for `fsm` declaration in the comment text
	for _, submatches := range fsmReg.FindAllStringSubmatch(commentText, -1) {
		results := make(map[string]string, len(submatches))
		for i, name := range submatches {
			results[fsmNames[i]] = name
		}

		name, state, op := results["name"], results["state"], results["op"]
		c.fsmSink(name, state, op)
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
