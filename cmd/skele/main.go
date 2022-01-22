package main

import "os"

func main() {
	if len(os.Args) == 1 {
		println("Usage: skele FILE")
		return
	}

	m := new(machine)

	for _, file := range os.Args[1:] {
		m.read(file)
	}

	m.write()
}
