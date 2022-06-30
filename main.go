package main

import (
	"bufio"
	"fmt"
	"hackassembler/module/code"
	parser "hackassembler/module/parser"
	"hackassembler/utils"
	"os"
)

const fileName = "PongL"

func main() {
	/* Initializer */
	rFile, err := os.Open("asm/" + fileName + ".asm")
	utils.HandleErr(err)
	wFile, err := os.Create("hack/" + fileName + ".hack")
	utils.HandleErr(err)
	defer rFile.Close()
	defer wFile.Close()
	scanner := bufio.NewScanner(rFile)
	fmt.Println("******************** ASSEMBLE START ********************")
	for parser.HasMoreCommands(scanner) {
		line := scanner.Text()
		fmt.Printf("########## current line : %s ##########\n", line)
		if parser.IsEmptyOrComment(line) {
			fmt.Println("########## skip ##########")
			continue
		}
		command := parser.Advance(line)
		binary := code.FromCommand(command)
		fmt.Fprintf(wFile, fmt.Sprintln(binary.Line))
		fmt.Printf("########## binary line : %s ##########\n", binary.Line)
	}
	utils.HandleErr(scanner.Err())
	fmt.Println("******************** ASSEMBLE END ********************")
}
