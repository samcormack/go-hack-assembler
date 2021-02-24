package parser

import (
	"bufio"
	"fmt"
	"os"
	s "strings"
)

const (
	A_COMMAND = "A"
	C_COMMAND = "C"
	L_COMMAND = "L"
)

type Parser struct {
	file *os.File
	scanner *bufio.Scanner
	Command string
}

func NewParser(file *os.File) *Parser {
	scanner := bufio.NewScanner(file)
	p := Parser{file: file, scanner: scanner}
	return &p
}

func (p *Parser) Print() {
	for p.scanner.Scan() {
		fmt.Println(p.scanner.Text())
	}
}

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

func (p *Parser) Advance() {
	subs := s.SplitN(p.scanner.Text(), "//", 2)
	p.Command = s.TrimSpace(subs[0])
}

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


func (p *Parser) Dest() string {
	subs := s.SplitN(p.Command, "=", 2)
	if len(subs) == 1 {
		return ""
	}
	return s.TrimSpace(subs[0])
	
}

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

func (p *Parser) Jump() string {
	subs := s.SplitN(p.Command, ";", 2)
	if len(subs) == 1 {
		return ""
	}
	return s.TrimSpace(subs[len(subs)-1])
}