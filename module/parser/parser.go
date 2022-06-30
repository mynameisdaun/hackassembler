package parser

import (
	"bufio"
	"hackassembler/common"
	"regexp"
	"strings"
)

type Command struct {
	Line        string
	Dest        string
	Comp        string
	Jump        string
	CommandType string
	Symbol      string
	LineNumber  int
}

// 입력에 명령이 더 있는가?
func HasMoreCommands(scanner *bufio.Scanner) bool {
	return scanner.Scan()
}

/* 입력에서 다음 명령을 읽어서, 현재 명령으로 만든다. */
func Advance(line string, lineNumber int) Command {
	line = strings.Replace(line, " ", "", -1)
	command := Command{
		Line:        line,
		Dest:        dest(line),
		Comp:        comp(line),
		Jump:        jump(line),
		CommandType: commandType(line),
		Symbol:      symbol(line),
		LineNumber:  lineNumber,
	}
	return command
}

/* 주석이나 비어있는 문장인지 확인한다 */
func IsEmptyOrComment(line string) bool {
	return strings.HasPrefix(line, "//") || len(line) < 1
}

/* 명령어의 커맨드 타입을 반환한다 */
func commandType(line string) string {
	if strings.HasPrefix(line, "(") {
		return common.COMMAND_L
	}
	if strings.HasPrefix(line, "@") {
		return common.COMMAND_A
	}
	return common.COMMAND_C
}

func symbol(line string) string {
	regex := regexp.MustCompile(`\(|\)|@`)
	return regex.ReplaceAllString(line, "")
}

func dest(line string) string {
	if !strings.Contains(line, "=") {
		return "null0"
	}
	return strings.Split(line, "=")[0]
}

func jump(line string) string {
	if !strings.Contains(line, ";") {
		return "null"
	}
	return strings.Split(line, ";")[1]
}

func comp(line string) string {

	if dest(line) != "null0" {
		line = strings.Replace(line, dest(line), "", 1)
	}
	if jump(line) != "null" {
		line = strings.Replace(line, jump(line), "", 1)
	}
	regex := regexp.MustCompile(`=|;`)
	return regex.ReplaceAllString(line, "")
}
