package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type HydrothermalVent = [999][999]int

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalln("Couldn't open the file")
	}

	// Initialize scanner on input.txt file
	scanner := bufio.NewScanner(file)
	// split scanner by new lines
	scanner.Split(bufio.ScanLines)

	var ventsRadar HydrothermalVent
	for scanner.Scan() {

		var points []int = translateInput(scanner.Text())
		x1, y1, x2, y2 := points[0], points[1], points[2], points[3]

		// Ignoring diagonals
		if x1 != x2 && y1 != y2 {
			continue
		}

		// It moves horizontaly
		if x1 != x2 {
			start, end := minMax(x1, x2)
			for start <= end {
				ventsRadar[y1][start]++
				start++
			}
		}

		// It moves verticaly
		if y1 != y2 {
			start, end := minMax(y1, y2)
			for start <= end {
				ventsRadar[start][x1]++
				start++
			}
		}

	}

	var overlapingVents int = 0
	for _, y := range ventsRadar {
		for _, x := range y {
			if x >= 2 {
				overlapingVents++
			}
		}
	}

	fmt.Println(overlapingVents)
}

// Get input and convert it into an array of ints
// It should look like this : [x1, y1, x2, y2]
func translateInput(inputStr string) []int {
	// 1 - scanmn the text, 2 - replace " -> " with ",", 3 - split all by comas
	var stringInput []string = strings.Split(strings.ReplaceAll(inputStr, " -> ", ","), ",")

	// Convert []string into []int
	var points []int
	for _, val := range stringInput {
		strToNum, err := strconv.Atoi(val)
		if err != nil {
			log.Fatalln("Not possible to convert this string into number")
		}
		points = append(points, strToNum)
	}
	return points
}

func minMax(a, b int) (int, int) {
	if a < b {
		return a, b
	}
	return b, a
}
