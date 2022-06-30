package utils

import (
	"fmt"
	"log"
	"strconv"
)

const COMMAND_A = "COMMAND_A"
const COMMAND_C = "COMMAND_C"
const COMMAND_L = "COMMAND_L"

/* HandleErr */
func HandleErr(err error) {
	if err != nil {
		log.Fatalf("Error when opening file: %s", err)
	}
}

func DecimalToBinary(str string) string {
	fmt.Printf("decimal is : %s\n", str)
	num, err := strconv.Atoi(str)
	HandleErr(err)
	n := int64(num)
	return fmt.Sprintf("%016b", n)
}
