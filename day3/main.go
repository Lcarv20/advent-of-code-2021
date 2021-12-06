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

	// *** Part 1 ***
	var bits []string

	for scanner.Scan() {
		bits = append(bits, scanner.Text())
	}

	for _, bit := range bits {
		for idx, char := range bit {
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

	// ------------------------------------------------------------------------------------------------------------
	// ### Part 2 ###

	var oxygenGenerator []string = bits
	x := 0
	for {
		if len(oxygenGenerator) <= 1 {
			break
		}

		zerosCounter, onesCounter := countDigits(oxygenGenerator, x)

		// filter only most common
		// Note.: If zeros and ones are present equaly, select just the ones
		if zerosCounter > 0 && onesCounter > 0 {
			if onesCounter >= zerosCounter {
				oxygenGenerator = filterBits("1", x, oxygenGenerator)
			} else {
				oxygenGenerator = filterBits("0", x, oxygenGenerator)
			}
		}

		// This shouldn't happen
		if x == len(bits[0])-1 {
			break
		}
		x++
	}

	var co2Scrubber []string = bits
	x2 := 0
	for {
		if len(co2Scrubber) <= 1 {
			break
		}

		zerosCounter, onesCounter := countDigits(co2Scrubber, x2)

		// filter most from less common into respective holders
		// Note.: If zeros and ones are present equaly, select just the zeros
		if zerosCounter > 0 && onesCounter > 0 {
			if zerosCounter <= onesCounter {
				co2Scrubber = filterBits("0", x2, co2Scrubber)
			} else {
				co2Scrubber = filterBits("1", x2, co2Scrubber)
			}
		}

		// This shouldn't happen either
		if x2 == len(bits[0])-1 {
			break
		}
		x2++
	}

	fmt.Println("Oxygen Generator rate: ", oxygenGenerator)
	fmt.Println("CO2 Scrubber rate: ", co2Scrubber)

	// Lets not care about parsing errors for now shall we?
	oxygenDecimal, _ := strconv.ParseInt(oxygenGenerator[0], 2, 64)
	co2Decimal, _ := strconv.ParseInt(co2Scrubber[0], 2, 64)

	lifeSupportRate := oxygenDecimal * co2Decimal

	fmt.Println("Life support Rate :", lifeSupportRate)

}

// Return total 0's and 1's
func countDigits(bits []string, searchIndex int) (zeros int, ones int) {
	zeros = 0
	ones = 0

	for _, bit := range bits {
		if bit[searchIndex] == '1' {
			ones++
		} else {
			zeros++
		}
	}
	return
}

func filterBits(digit string, searchIndex int, bitsToFilter []string) []string {
	// If there are only 2 remaining elements and the digit at the searchIndex is the same, return the first
	// This also shouldn't happen
	if len(bitsToFilter) == 2 && (bitsToFilter[0][searchIndex] == bitsToFilter[1][searchIndex]) {
		return []string{bitsToFilter[0]}
	}

	var accumulator []string
	for _, bit := range bitsToFilter {
		if string(bit[searchIndex]) == digit {
			accumulator = append(accumulator, bit)
		}
	}

	return accumulator
}
