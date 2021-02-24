package main

import (
	"assembler/code"
	"assembler/parser"
	"bufio"
	"log"
	"os"
	s "strings"
)

var ADD string = "./examples/Add.asm"
var MAX string = "./examples/Max.asm"
var MAXL string = "./examples/MaxL.asm"

func main() {
	fname := ADD
	if len(os.Args) > 1 {
		fname = os.Args[1]
	}

	infile, err := os.Open(fname)
	check(err)
	defer infile.Close()

	outfile, err := os.Create(s.TrimSuffix(fname, ".asm") + ".hack")
	check(err)
	defer outfile.Close()

	writer := bufio.NewWriter(outfile)
	defer writer.Flush()

	p := parser.NewParser(infile)
	var cmd string
	for p.HasMoreCommands() {
		p.Advance()
		switch p.CommandType() {
		case parser.A_COMMAND:
			cmd = code.TranslateA(p.Symbol())
		case parser.C_COMMAND:
			cmd = code.TranslateC(p.Dest(), p.Comp(), p.Jump())
		}
		writer.WriteString(cmd + "\n")
	}
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
