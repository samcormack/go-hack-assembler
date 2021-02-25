package parser

import (
	"assembler/symbol"
	"bufio"
	"fmt"
	"os"
	s "strings"
	"strconv"
	"log"
)

// Command type constants
const (
	A_COMMAND = "A"
	C_COMMAND = "C"
	L_COMMAND = "L"
)

// Object that parses a .asm file
type Parser struct {
	file *os.File
	scanner *bufio.Scanner
	Command string
	LineNo int64
	nextAddress int64
}

func NewParser(file *os.File) *Parser {
	scanner := bufio.NewScanner(file)
	p := Parser{file: file, scanner: scanner, LineNo: 0, nextAddress: 16}
	return &p
}

// Reset parser to start of file
func (p *Parser) Reset() {
	_,err := p.file.Seek(0,0)
	if err != nil {
		log.Fatal(err)
	}
	p.scanner = bufio.NewScanner(p.file)
	p.Command = ""
	p.LineNo = 0
	p.nextAddress = 16
}
// Print contents of input file
func (p *Parser) Print() {
	for p.scanner.Scan() {
		fmt.Println(p.scanner.Text())
	}
}

// Move to next command line in input and return true. Return false if no more commands
func (p *Parser) HasMoreCommands() bool {
	for p.scanner.Scan() {
		line := p.scanner.Text()
		if s.HasPrefix(line, "//") || s.TrimSpace(line) == "" {
			continue
		}
		return true
	} 
	return false
}

// Set Command of parse to current line in input
func (p *Parser) Advance() {
	subs := s.SplitN(p.scanner.Text(), "//", 2)
	p.Command = s.TrimSpace(subs[0])
	if p.CommandType() != L_COMMAND {
		p.LineNo++
	}
}

//Return command type of parser's Command
func (p *Parser) CommandType() string {
	switch {
	case s.HasPrefix(p.Command, "@"):
		return A_COMMAND
	case s.HasPrefix(p.Command, "("):
		return L_COMMAND
	default:
		return C_COMMAND
	}
}

// Return symbol of Parser's Command
func (p *Parser) Symbol() string {
	switch p.CommandType() {
	case A_COMMAND:
		return s.TrimPrefix(p.Command, "@")
	case L_COMMAND:
		return s.TrimPrefix(s.TrimSuffix(p.Command, ")"),"(")
	case C_COMMAND:
		panic("Tried to access symbol of C command")
	}
	return ""
}

// Return address of an A command
func (p *Parser) Address(st *symbol.SymbolTable) int64 {
	sym := p.Symbol()
	address, err := strconv.ParseInt(sym, 10, 64)
	if err == nil {
		return address
	} 
	if !st.Contains(sym) {
		st.AddEntry(sym, p.nextAddress)
		p.nextAddress++
	}
	return st.GetAddress(sym)
}

// Return current dest mnemonic
func (p *Parser) Dest() string {
	subs := s.SplitN(p.Command, "=", 2)
	if len(subs) == 1 {
		return ""
	}
	return s.TrimSpace(subs[0])
	
}

// Return current comp mnemonic
func (p *Parser) Comp() string {
	var right string
	subs := s.SplitN(p.Command, "=", 2)
	if len(subs) == 1 {
		right = subs[0]
	} else {
		right = subs[1]
	}
	subs = s.SplitN(right, ";", 2)
	return s.TrimSpace(subs[0])
}

// Return current jump mnemonic
func (p *Parser) Jump() string {
	subs := s.SplitN(p.Command, ";", 2)
	if len(subs) == 1 {
		return ""
	}
	return s.TrimSpace(subs[len(subs)-1])
}