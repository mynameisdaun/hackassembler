package module

import (
	"bufio"
	"log"
	"strconv"
	"strings"
)

var arithmeticCommands = [9]string{"add", "sub", "neg", "eq", "gt", "lt", "and", "or", "not"}

func HasMoreCommands(scanner *bufio.Scanner) bool {
	return scanner.Scan()
}

func DeleteAllWhiteSpaceAndComment(line string) string {
	if strings.Index(line, "//") > -1 {
		line = strings.Split(line, "//")[0]
	}
	return strings.Replace(line, " ", "", -1)
}

func CommandType(line string) string {
	for _, operator := range arithmeticCommands {
		if strings.HasPrefix(line, operator) {
			return "C_ARITHMETIC"
		}
	}
	if strings.HasPrefix(line, "push") {
		return "C_PUSH"
	}
	if strings.HasPrefix(line, "pop") {
		return "C_POP"
	}
	if strings.HasPrefix(line, "") {
		return "C_LABEL"
	}
	if strings.HasPrefix(line, "push") {
		return "C_GOTO"
	}
	if strings.HasPrefix(line, "push") {
		return "C_IF"
	}
	if strings.HasPrefix(line, "push") {
		return "C_FUNCTION"
	}
	if strings.HasPrefix(line, "push") {
		return "C_RETURN"
	}
	if strings.HasPrefix(line, "push") {
		return "C_CALL"
	}
	return ""
}

func Arg1(line string) string {
	if CommandType(line) == "C_ARITHMETIC" {
		return line
	}
	return strings.Split(line, " ")[0]
}

func Arg2(line string) int {
	cType := CommandType(line)
	if cType != "C_PUSH" && cType != "C_POP" && cType != "C_FUNCTION" && cType != "C_CALL" {
		panic("올바르지 않은 커맨드 타입에서 두번째 인수 호출하지 마라 ...")
	}
	value, err := strconv.Atoi(strings.Split(line, " ")[1])
	if err != nil {
		log.Panic(err)
	}
	return value
}
