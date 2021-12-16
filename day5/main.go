package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type HydrothermalVent = [1000][1000]int

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalln("Couldn't open the file")
	}

	// Problem identifier
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Choose problem to solve - part1 or part2")

	userChoice, _ := reader.ReadString('\n')
	userChoice = strings.Trim(userChoice, "\n")

	if userChoice == "part1" {
		fmt.Println("hole")
	}

	// Initialize scanner on input.txt file
	scanner := bufio.NewScanner(file)
	// split scanner by new lines
	scanner.Split(bufio.ScanLines)

	var ventsRadar HydrothermalVent
	for scanner.Scan() {

		var points []int = translateInput(scanner.Text())
		x1, y1, x2, y2 := points[0], points[1], points[2], points[3]

		if userChoice == "part1" {
			// Ignoring all diagonals
			if x1 != x2 && y1 != y2 {
				continue
			}
		}
		if userChoice == "part2" {
			// to be 45 degrees x2-x1 == y2-y1
			startX, endX := minMax(x1, x2)
			startY, endY := minMax(y1, y2)

			// check if it is 45 degrees
			if (endX - startX) == (endY - startY) {
				ventsRadar = paintDiagonalVent(ventsRadar, x1, y1, x2, y2)
				continue
			}

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

	// fmt.Println(ventsRadar)

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

func paintDiagonalVent(ventsRadar [1000][1000]int, x1 int, y1 int, x2 int, y2 int) [1000][1000]int {
	// fmt.Printf("[%d, %d] - [%d, %d]\n", x1, y1, x2, y2)
	for x1 != x2 && y1 != y2 {
		switch true {
		case x1 < x2 && y1 < y2:
			ventsRadar[y1][x1]++
			x1++
			y1++

		case x1 > x2 && y1 > y2:
			ventsRadar[y1][x1]++
			x1--
			y1--

		case x1 < x2 && y1 > y2:
			ventsRadar[y1][x1]++
			x1++
			y1--

		case x1 > x2 && y1 < y2:
			ventsRadar[y1][x1]++
			x1--
			y1++
		}
	}
	ventsRadar[y1][x1]++
	return ventsRadar
}
