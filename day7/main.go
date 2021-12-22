package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalln("There was an error opening the file -> ", err)
	}

	scanner := bufio.NewScanner(file)
	// scanner.Split(bufio.ScanLines)
	scanner.Scan()

	crabsMap := processInput(scanner.Text())

	var fuelRecords []int

	for destination := range crabsMap {
		fuelSpent := 0
		for origin, multiplier := range crabsMap {
			fuelSpent +=
				// calculate distance between origin and destination
				// pkg math only works with floats, so I have to convert ints to floats,
				// operate and then convert back to ints
				int(
					math.Abs(
						float64(destination-(origin)),
					),
				) * multiplier
		}
		fuelRecords = append(fuelRecords, fuelSpent)
	}

	sort.Ints(fuelRecords)

	fmt.Println(fuelRecords[0])
}

func processInput(input string) map[int]int {
	strArr := strings.Split(input, ",")
	crabMap := make(map[int]int)

	for _, num := range strArr {
		n, err := strconv.Atoi(num)
		if err != nil {
			log.Fatalln("Error parsing the string to number ->", err)
		}

		crabMap[n]++
	}
	return crabMap
}
