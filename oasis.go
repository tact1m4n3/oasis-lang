package main

import (
	"fmt"
	"oasis/lexer"
	"oasis/parser"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "no input file specified")
		os.Exit(1)
	}

	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
		os.Exit(1)
	}

	l := lexer.New(string(data))
	p := parser.New(l)
	p.Error()

	program := p.ParseProgram()
	if program == nil {
		fmt.Fprintf(os.Stderr, "%s", p.Error())
		os.Exit(1)
	}

	fmt.Println(program)
}
