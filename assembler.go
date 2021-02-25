package main

import (
	"assembler/code"
	"assembler/parser"
	"assembler/symbol"
	"bufio"
	"log"
	"os"
	s "strings"
)

var ADD string = "./examples/Add.asm"
var MAX string = "./examples/Max.asm"
var MAXL string = "./examples/MaxL.asm"

func main() {
	fname := MAX // Default input

	// Get input file and open
	if len(os.Args) > 1 {
		fname = os.Args[1]
	}
	infile, err := os.Open(fname)
	check(err)
	defer infile.Close()

	// Create output file and writer
	outfile, err := os.Create(s.TrimSuffix(fname, ".asm") + ".hack")
	check(err)
	defer outfile.Close()
	writer := bufio.NewWriter(outfile)
	defer writer.Flush()

	// Create symbol table
	symbolTable := symbol.NewSymbolTable()
	
	// First pass: find labels and add them to symbol table
	p := parser.NewParser(infile)
	for p.HasMoreCommands() {
		p.Advance()
		if p.CommandType() == parser.L_COMMAND {
			symbolTable.AddEntry(p.Symbol(), p.LineNo)
		}
	}

	// Parse input and write to output
	p.Reset()
	var cmd string
	for p.HasMoreCommands() {
		p.Advance()
		switch p.CommandType() {
		case parser.A_COMMAND:
			cmd = code.TranslateA(p.Address(symbolTable))
			writer.WriteString(cmd + "\n")
		case parser.C_COMMAND:
			cmd = code.TranslateC(p.Dest(), p.Comp(), p.Jump())
			writer.WriteString(cmd + "\n")
		}
		
	}
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
