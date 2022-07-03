package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	rFile, err := os.Open("vmware/vm/StackArithmetic/SimpleAdd/SimpleAdd.vm")
	HandleErr(err)
	scanner := bufio.NewScanner(rFile)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func HandleErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}
