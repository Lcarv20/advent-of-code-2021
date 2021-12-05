package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Position struct {
	horizontalPos int
	depthPos      int
	aim           int
}

func (p *Position) finalVal() {
	fmt.Println(p.depthPos * p.horizontalPos)
}

func main() {
	var pos = Position{
		horizontalPos: 0,
		depthPos:      0,
	}

	var pos2 = Position{
		horizontalPos: 0,
		depthPos:      0,
		aim:           0,
	}

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Could not open the file")
	}

	// initialize scanner
	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		var line []string = strings.Split(scanner.Text(), " ")
		operation := line[0]
		value, err := strconv.Atoi(line[1])
		if err != nil {
			log.Fatal("Error parsing the value from string to Int")
		}

		// *** Part 1 of the problem ***
		switch operation {
		// depth
		case "up":
			pos.depthPos -= value

		case "down":
			pos.depthPos += value

		// horizontal
		case "forward":
			pos.horizontalPos += value

		default:
			log.Fatal("Something went wrong")
		}

		// ### Part 2 of the problem ###
		switch operation {
		case "forward":
			pos2.horizontalPos += value
			pos2.depthPos += value * pos2.aim

		case "down":
			pos2.aim += value

		case "up":
			pos2.aim -= value
		}
	}

	pos.finalVal()
	pos2.finalVal()

	file.Close()

}
