package common

import (
	"fmt"
	"log"
	"strconv"
)

/* HandleErr */
func HandleErr(err error) {
	if err != nil {
		log.Fatalf("Error when opening file: %s", err)
	}
}

func DecimalToBinary(str string) string {
	num, err := strconv.Atoi(str)
	HandleErr(err)
	n := int64(num)
	return fmt.Sprintf("%016b", n)
}
