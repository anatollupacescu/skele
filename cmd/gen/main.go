package main

import "os"

func main() {
	var m = new(machine)

	args := os.Args

	if len(args) == 1 {
		println("Usage: gen FILE")
		return
	}

	for i, arg := range os.Args {
		if i == 0 {
			continue
		}
		m.read(arg)
	}

	m.write()
}
