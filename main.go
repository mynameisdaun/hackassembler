package main

import (
	"bufio"
	"fmt"
	"hackassembler/common"
	"hackassembler/module/code"
	parser "hackassembler/module/parser"
	table "hackassembler/module/symboltable"
	"os"
	"strings"
)

func main() {
	/* Initializer */
	var fileName string
	fmt.Print("Please enter the Assembly file name which is in the 'asm' directory...\nex) Max\n")
	fmt.Scanf("%s", &fileName)
	pass1, err := os.Open("asm/" + fileName + ".asm")
	pass2, err := os.Open("asm/" + fileName + ".asm")
	common.HandleErr(err)
	wFile, err := os.Create("hack/" + fileName + ".hack")
	common.HandleErr(err)
	defer pass1.Close()
	defer pass2.Close()
	defer wFile.Close()

	symbolTable := table.SymbolTable()
	scanner := bufio.NewScanner(pass1)
	lineNumber := 0

	for parser.HasMoreCommands(scanner) {
		line := deleteInlineComment(scanner.Text())
		if parser.IsEmptyOrComment(line) {
			continue
		}
		command := parser.Advance(line, lineNumber)
		if command.CommandType == common.COMMAND_L {
			symbolTable.AddCommand(command.Symbol, command.LineNumber)
		} else {
			lineNumber++
		}
	}
	scanner = bufio.NewScanner(pass2)
	for parser.HasMoreCommands(scanner) {
		line := deleteInlineComment(scanner.Text())
		if parser.IsEmptyOrComment(line) {
			continue
		}
		command := parser.Advance(line, lineNumber)
		if command.CommandType != common.COMMAND_L {
			binary := code.FromCommand(command)
			fmt.Fprintf(wFile, fmt.Sprintln(binary.Line))
		}
	}
	common.HandleErr(scanner.Err())
}

func deleteInlineComment(line string) string {
	hasComment := strings.Index(line, "//")
	if hasComment > -1 {
		line = strings.Split(line, "//")[0]
	}
	return line
}
