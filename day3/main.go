package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalln("Something went wrong while reading the file")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var bits []string

	for scanner.Scan() {
		bits = append(bits, scanner.Text())
	}
	fmt.Println(len(bits))
}
