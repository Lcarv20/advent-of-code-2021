package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Rates struct {
	zeros int
	ones  int
}

func main() {
	var powerConsumption []Rates
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalln("Something went wrong while reading the file")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		bits := scanner.Text()

		for idx, char := range bits {
			if len(powerConsumption) <= idx {
				powerConsumption = append(powerConsumption, Rates{})
			}
			if char == '0' {
				powerConsumption[idx].zeros += 1

			} else {
				powerConsumption[idx].ones += 1
			}
		}
	}

	gamaRate, epsilonRate := "", ""

	for _, powerData := range powerConsumption {
		if powerData.zeros > powerData.ones {
			gamaRate += "0"
			epsilonRate += "1"
		} else {
			gamaRate += "1"
			epsilonRate += "0"
		}
	}

	gama, _ := strconv.ParseInt(gamaRate, 2, 64)
	epsilon, _ := strconv.ParseInt(epsilonRate, 2, 64)

	fmt.Println(gama * epsilon)
}
