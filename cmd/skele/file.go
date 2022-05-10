package main

import (
	"bytes"
	"log"
	"os"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/anatollupacescu/skele/parser"
)

func (m *machine) read(name string) []error {
	str := readFile(name)
	is := antlr.NewInputStream(str)
	lexer := parser.NewSkeleLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewSkele(stream)
	obj := new(skeleListener)
	obj.machine = m
	antlr.ParseTreeWalkerDefault.Walk(obj, p.Start())
	return obj.errors
}

func readFile(name string) string {
	file, err := os.Open(name)
	if err != nil {
		log.Fatalf("open file: %s: %v\n", name, err)
	}
	defer file.Close()
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(file)
	if err != nil {
		log.Fatal("read file\n", err)
	}
	return buf.String()
}
