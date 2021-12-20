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
		log.Fatalln("An error ocurred while opening the file -> ", err)
	}

	scanner := bufio.NewScanner(input)
	scanner.Scan()
	stringArr := strings.Split(scanner.Text(), ",")
	var initialState []int
	for _, str := range stringArr {
		val, err := strconv.Atoi(str)
		if err != nil {
			log.Fatalln("An error ocurred while parsing the strings to ints")
		}
		initialState = append(initialState, val)
	}

	// *** PART 1 ***
	// lifecylce over 80 days

	var lanternFish []int
	// copy of original state
	lanternFish = append(lanternFish, initialState...)
	for i := 1; i <= 80; i++ {
		lanternFish = dayLog(lanternFish)
	}
	fmt.Println(len(lanternFish))

	// ### PART 2 ###
	var lanternMap map[int]int = map[int]int{0: 0, 6: 0, 7: 0, 8: 0}

	// convert into a map key = fish day | val = number of fish in that day
	for _, fish := range initialState {
		lanternMap[fish]++
	}

	for day := 1; day <= 256; day++ {
		lanternMap = dayLog2(lanternMap)
	}

	fishAccumulator := 0

	for _, numFish := range lanternMap {
		fishAccumulator += numFish
	}

	fmt.Println(fishAccumulator)

}

func dayLog(lanternFish []int) []int {
	var lanterFry []int
	for idx, fish := range lanternFish {
		if fish > 0 {
			lanternFish[idx]--
		} else {
			lanternFish[idx] = 6
			lanterFry = append(lanterFry, 8)
		}
	}

	lanternFish = append(lanternFish, lanterFry...)

	return lanternFish
}

func dayLog2(lanternMap map[int]int) map[int]int {
	lanternFry := 0

	for lFish := 0; lFish < len(lanternMap)-1; lFish++ {
		if lFish == 0 {
			lanternFry = lanternMap[0]
		}
		lanternMap[lFish] = lanternMap[lFish+1]
	}

	// the fish that gave birth to new fry go to 6 and the fry to 8, ratio 1:1
	lanternMap[8] = lanternFry
	lanternMap[6] += lanternFry
	return lanternMap
}
