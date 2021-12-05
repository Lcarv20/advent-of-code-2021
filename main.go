package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	// Open connection to the file
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Failed to open the file")
	}

	// Create NewScanner object
	scanner := bufio.NewScanner(file)

	// The bufio.ScanLines is used as an
	// input to the method bufio.Scanner.Split()
	// and then the scanning forwards to each
	// new line using the bufio.Scanner.Scan()
	// method.
	scanner.Split(bufio.ScanLines)
	var depth []int

	for scanner.Scan() {
		value, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal("Failed to parse the string to int")
		}
		depth = append(depth, value)
	}

	// Cut the connection to the file
	file.Close()

	// *** Part 1 of the problem ***
	var largerMeasurements int = 0
	for measurement := 1; measurement < len(depth); measurement++ {
		if depth[measurement] > depth[measurement-1] {
			largerMeasurements++
		}
	}
	fmt.Println(largerMeasurements)

	// ### Part 2 of the problem ###
	var largerMeasurementsWindow int = 0

	for measurement := 0; measurement < len(depth)-3; measurement++ {
		var window1 int = depth[measurement] + depth[measurement+1] + depth[measurement+2]
		var window2 int = depth[measurement+1] + depth[measurement+2] + depth[measurement+3]
		if window2 > window1 {
			largerMeasurementsWindow++
		}
	}

	fmt.Println(largerMeasurementsWindow)
}
