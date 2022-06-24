package main

import (
	"errors"
	"fmt"
)

type mFSMS map[string]mFSM

type mFSM struct {
	pkg    string
	states []string
	tt     []transition
}

type transition struct {
	fun      string
	from, to string
}

var (
	ErrDuplicateFSM      = errors.New("duplicate fsm")
	ErrStatesSize        = errors.New("not enough states")
	ErrInvalidTransition = errors.New("invalid transition")
)

func (m mFSMS) add(pkg string, f fsm) error {
	if len(f.states) < 2 {
		return fmt.Errorf("name %s: %w", f.name, ErrStatesSize)
	}

	if dpkg, ok := m[f.name]; ok {
		return fmt.Errorf("previous declaration of %s in package %s: %w", f.name, dpkg, ErrDuplicateFSM)
	}

	m[f.name] = mFSM{
		pkg:    pkg,
		states: f.states,
	}

	return nil
}

func (mFSMS) checkCirclular() []error {
	return nil
}

func (mFSMS) checkUnused() []error {
	return nil
}

// Transitions enriches the given FSM with its transitions
func (m mFSMS) fun(fun fun) (errs []error) {
	t := transition{from: "init"}

	if fun.fsm != nil {
		t.from = fun.fsm.state
	}

	// for each pre / check valid reference

	// complete transition
	for _, pos := range fun.pos {
		if pos.fsm.name == "" {
			continue
		}

		cf := m[pos.fsm.name]
		if len(cf.states) == 0 {
			errs = append(errs, fmt.Errorf("state machine not found: %s: %w", pos.fsm.name, ErrInvalidTransition))
			return
		}

		init := cf.states[0]

		if fun.fsm != nil && pos.fsm.name != fun.fsm.name && pos.fsm.state != init {
			errs = append(errs, fmt.Errorf("invalid combination of start/end states: \nfun '%s'\nstart: %s=%s\nend: %s->%s: %w", fun.name, fun.fsm.name, fun.fsm.state, pos.fsm.name, pos.fsm.state, ErrInvalidTransition))
			return
		}

		t.to = pos.fsm.state
		t.fun = fun.name

		cf.tt = append(m[pos.fsm.name].tt, t)
		m[pos.fsm.name] = cf
		return
	}

	return
}
