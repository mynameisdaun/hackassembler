package code

import (
	common2 "hackassembler/hackassembler/common"
	"hackassembler/hackassembler/module/parser"
	"hackassembler/hackassembler/module/symboltable"
	"strconv"
)

type Binary struct {
	Line string
	Dest string
	Comp string
	Jump string
}

func FromCommand(command parser.Command) Binary {
	symbolTable := symboltable.SymbolTable()
	binary := Binary{}
	if command.CommandType == common2.COMMAND_C {
		binary.Dest = dest(command.Dest)
		binary.Comp = comp(command.Comp)
		binary.Jump = jump(command.Jump)
		binary.Line = "111" + binary.Comp + binary.Dest + binary.Jump
	}
	if command.CommandType == common2.COMMAND_A {
		if _, err := strconv.Atoi(command.Symbol); err == nil {
			binary.Line = common2.DecimalToBinary(command.Symbol)
		} else {
			if !symboltable.Contains(symbolTable, command.Symbol) {
				symbolTable.AddVariable(command.Symbol)
			}
			address := symboltable.GetAddress(symbolTable, command.Symbol)
			binary.Line = common2.DecimalToBinary(strconv.Itoa(address))
		}
	}
	return binary
}

var destMap = map[string]string{
	"null0": "000",
	"M":     "001",
	"D":     "010",
	"MD":    "011",
	"A":     "100",
	"AM":    "101",
	"AD":    "110",
	"AMD":   "111",
}

var jumpMap = map[string]string{
	"null": "000",
	"JGT":  "001",
	"JEQ":  "010",
	"JGE":  "011",
	"JLT":  "100",
	"JNE":  "101",
	"JLE":  "110",
	"JMP":  "111",
}

var compMap = map[string]string{
	"0":   "0101010",
	"1":   "0111111",
	"-1":  "0111010",
	"D":   "0001100",
	"A":   "0110000",
	"M":   "1110000",
	"!D":  "0001101",
	"!A":  "0110001",
	"!M":  "1110001",
	"-D":  "0001111",
	"-A":  "0110011",
	"-M":  "1110011",
	"D+1": "0011111",
	"A+1": "0110111",
	"M+1": "1110111",
	"D-1": "0001110",
	"A-1": "0110010",
	"M-1": "1110010",
	"D+A": "0000010",
	"D+M": "1000010",
	"D-A": "0010011",
	"D-M": "1010011",
	"A-D": "0000111",
	"M-D": "1000111",
	"D&A": "0000000",
	"D&M": "1000000",
	"D|A": "0010101",
	"D|M": "1010101",
}

func dest(dest string) string {
	return destMap[dest]
}

func comp(comp string) string {
	return compMap[comp]
}

func jump(jump string) string {
	return jumpMap[jump]
}
