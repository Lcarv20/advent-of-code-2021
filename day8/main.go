package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatalln("Couldn't read input from file -> ", err)
	}

	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	var numsRecord = make(map[int]int)
	for scanner.Scan() {
		outputConsole := processOutput(scanner.Text())

		for _, val := range outputConsole {
			switch len(val) {
			case 2:
				numsRecord[1]++

			case 3:
				numsRecord[7]++

			case 4:
				numsRecord[4]++

			case 7:
				numsRecord[8]++

			default:
				continue
			}
		}
	}
	fmt.Printf("%#v\n", numsRecord)
	resPart1 := calculateRes(numsRecord)
	print(resPart1)
}

func processOutput(str string) []string {
	return strings.Fields(strings.Split(str, "|")[1])
}

func calculateRes(numsRecord map[int]int) int {
	var total int
	for _, val := range numsRecord {
		total += val
	}

	return total
}
