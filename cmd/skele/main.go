package main

import (
	"log"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		println("Usage: skele FILE")
		return
	}

	m := new(machine)

	var errs errs
	for _, file := range os.Args[1:] {
		errs = append(errs, m.read(file)...)
	}

	for _, err := range m.ValidateFSM() {
		errs = append(errs, err)
	}

	if len(errs) > 0 {
		errs.print()
		return
	}

	m.write()
}

type errs []error

func (e errs) print() {
	for _, err := range e {
		log.Println(err)
	}
}
