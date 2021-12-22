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

	// *** PART 1 ***
	var fuelRecords []int
	for destination := range crabsMap {
		fuelSpent := 0
		for origin, multiplier := range crabsMap {
			fuelSpent += (calculateDistance(origin, destination) * multiplier)
		}
		fuelRecords = append(fuelRecords, fuelSpent)
	}

	// ### PART 2 ###
	var fuelRecords2 []int
	for destination := range crabsMap {
		fuelSpent := 0
		for origin, multiplier := range crabsMap {
			fuelSpent += (triangularDistance(origin, destination) * multiplier)
		}
		fuelRecords2 = append(fuelRecords2, fuelSpent)
	}
	//

	sort.Ints(fuelRecords)
	sort.Ints(fuelRecords2)

	fmt.Println(fuelRecords[0])
	fmt.Println(fuelRecords2[0])
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

func calculateDistance(origin, destination int) int {
	// pkg math only works with floats, so I have to convert ints to floats,
	// operate and then convert back to ints
	return int(
		math.Abs(
			float64(destination - (origin)),
		),
	)
}

func triangularDistance(origin, destination int) int {
	linearDist := calculateDistance(origin, destination)

	return ((linearDist * (linearDist + 1)) / 2)
}
