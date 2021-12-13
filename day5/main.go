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
	var ventsRadar HydrothermalVent
	var overlapingVents int = 0

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalln("Couldn't open the file")
	}

	// Initialize scanner on input.txt file
	scanner := bufio.NewScanner(file)
	// split scanner by new lines
	scanner.Split(bufio.ScanLines)

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
				if ventsRadar[y1][start] >= 2 {
					overlapingVents++
				}
				start++
			}
		}

		// It moves verticaly
		if y1 != y2 {
			start, end := minMax(y1, y2)

			for start <= end {
				ventsRadar[start][x1]++
				if ventsRadar[start][x1] >= 2 {
					overlapingVents++
				}
				start++
			}
		}

	}

	sum := 0

	for _, i := range ventsRadar {
		for _, j := range i {
			if j >= 2 {
				sum++
			}
		}
	}

	// this works
	fmt.Println(sum)

	// but this doesn't
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
