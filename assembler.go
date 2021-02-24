package main

import (
	// "bufio"
	"os"
	"log"
	"fmt"
	"assembler/parser"
)

var ADD string = "/home/sam/Dropbox/Sam/CS/Core CS/10. Nand2Tetris/nand2tetris/projects/06/add/Add.asm"
var MAX string = "/home/sam/Dropbox/Sam/CS/Core CS/10. Nand2Tetris/nand2tetris/projects/06/max/Max.asm"
var MAXL string = "/home/sam/Dropbox/Sam/CS/Core CS/10. Nand2Tetris/nand2tetris/projects/06/max/MaxL.asm"

func main() {
	infile, err := os.Open(MAXL)
	check(err)
	defer infile.Close()

	p := parser.NewParser(infile)
	for p.HasMoreCommands() {
		p.Advance()
		switch p.CommandType() {
		case parser.A_COMMAND:
			fmt.Println(p.CommandType(), ": ",p.Symbol())
		case parser.C_COMMAND:
			fmt.Println(p.CommandType(), ": ",p.Dest(), ", ", p.Comp(), ", ", p.Jump())
		}
		
	}
	
	
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}