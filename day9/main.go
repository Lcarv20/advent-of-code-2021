package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatalln("there was a problem while opening the file")
	}

	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	var heightMap [][]int

	for scanner.Scan() {
		ipt := scanner.Text()
		heightMap = append(heightMap, processInput(ipt))
	}
	input.Close()
	// fmt.Printf("%#v\n", heightMap)

	// Sum of all low points
	lowerPoints := sumLowerPoints(heightMap)
	// _ = lowerPoints
	fmt.Println("Lower points sum:", lowerPoints)

}

func processInput(strLine string) []int {
	strArr := strings.Split(strLine, "")
	var intArr []int
	for _, char := range strArr {
		height, err := strconv.Atoi(char)
		if err != nil {
			log.Fatalln("An error ocurred while parsin string to int ->", err)
		}
		intArr = append(intArr, height)
	}
	return intArr
}

func sumLowerPoints(heightMap [][]int) int {
	var lowerLocations []int = findLowerLocations(heightMap)
	sum := 0
	for _, location := range lowerLocations {
		// The result of the first part of the problem consists on the sum of the height on lower locations + 1
		sum += location + 1
	}
	return sum
}

func findLowerLocations(heightMap [][]int) []int {
	var lowerLocations []int
	maxX := len(heightMap[0]) - 1
	maxY := len(heightMap) - 1

	for x, row := range heightMap {
		for y, cell := range row {
			neighbors := findNeighbors(heightMap, x, y, maxX, maxY)
			lowerLocations = append(lowerLocations, verifyIfLower(cell, neighbors))
		}
	}

	return lowerLocations
}

func findNeighbors(heightMap [][]int, x, y, maxX, maxY int) []int {
	var neighbors []int

	// Find all tge possible neighbors
	switch x {
	case 0:
		neighbors = append(neighbors, heightMap[x+1][y])
	case maxX:
		neighbors = append(neighbors, heightMap[x-1][y])
	default:
		neighbors = append(neighbors, heightMap[x-1][y], heightMap[x+1][y])
	}

	switch y {
	case 0:
		neighbors = append(neighbors, heightMap[x][y+1])
	case maxY:
		neighbors = append(neighbors, heightMap[x][y-1])
	default:
		neighbors = append(neighbors, heightMap[x][y-1], heightMap[x][y+1])
	}

	return neighbors
}

func verifyIfLower(currCell int, neighbors []int) int {
	for _, neighbor := range neighbors {
		if neighbor <= currCell {
			// The end result is gonna be the sum of all lower locations +1, so I return -1 to not affect end result
			return -1
		}
	}
	return currCell
}
