package day23

import (
	"strconv"
	"strings"
	"text/scanner"
)

func parse(input string) *machine {
	machine := &machine{}

	reader := strings.NewReader(input)
	var s scanner.Scanner
	s.Init(reader)

	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		op := s.TokenText()
		switch op {
		case "inc", "tpl", "hlf":
			s.Scan()
			reg := s.TokenText()
			machine.program = append(machine.program, instruction{op: op, reg: reg})
		case "jmp":
			s.Scan()
			sgn := s.TokenText()
			s.Scan()
			offset, err := strconv.Atoi(s.TokenText())
			if err != nil {
				panic(err)
			}
			if sgn == "-" {
				offset = -offset
			}
			machine.program = append(machine.program, instruction{op: op, offset: offset})
		case "jio", "jie":
			s.Scan()
			reg := s.TokenText()
			_ = s.Scan() // skip comma
			_ = s.Scan()
			sgn := s.TokenText()
			_ = s.Scan()
			offset, err := strconv.Atoi(s.TokenText())
			if err != nil {
				panic(err)
			}
			if sgn == "-" {
				offset = -offset
			}
			machine.program = append(machine.program, instruction{op: op, reg: reg, offset: offset})
		}
	}
	return machine
}

type instruction struct {
	op     string
	reg    string
	offset int
}

type machine struct {
	pc, a, b int
	program  []instruction
}

func (m *machine) getReg(reg string) *int {
	if reg == "a" {
		return &m.a
	}
	return &m.b
}

func (m *machine) step() bool {
	op := m.program[m.pc]
	switch op.op {
	case "hlf": // half register
		reg := m.getReg(op.reg)
		*reg /= 2
		m.pc++
	case "tpl": // triple register
		reg := m.getReg(op.reg)
		*reg *= 3
		m.pc++
	case "inc": // increment register
		reg := m.getReg(op.reg)
		*reg++
		m.pc++
	case "jmp": // jump to offset
		m.pc += op.offset
	case "jie": // jump if even
		reg := m.getReg(op.reg)
		if *reg%2 == 0 {
			m.pc += op.offset
		} else {
			m.pc++
		}
	case "jio": // jump if one
		reg := m.getReg(op.reg)
		if *reg == 1 {
			m.pc += op.offset
		} else {
			m.pc++
		}
	}
	return m.pc < len(m.program)
}

func (m *machine) execute() {
	for m.step() {
	}
}
