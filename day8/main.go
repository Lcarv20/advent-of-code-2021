package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatalln("Couldn't read input from file -> ", err)
	}

	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	res1, res2 := 0, 0
	for scanner.Scan() {
		ipt := scanner.Text()

		numsDic, numsOutput := filterInput(ipt)
		// *** PART 1 ***
		// see how many times numbers 1, 7, 4 and 8 appear
		for _, val := range numsOutput {
			length := len(val)
			if length == 2 || length == 3 || length == 4 || length == 7 {
				res1++
			}
		}

		// ### PART 2 ###
		// calculate sum of output (numbers after "|" )
		numsDictionary := sortNumbers(numsDic)
		res2 += translateAndCalculate(numsDictionary, numsOutput)
	}
	fmt.Println("Part I ->", res1, "\nPart II ->", res2)
}

func filterInput(str string) ([]string, []string) {
	strArr := strings.Split(str, " | ")
	nDic, nOutput := strings.Fields(strArr[0]), strings.Fields(strArr[1])

	// sort each number representation for easier comparison later
	for i := 0; i < len(nDic); i++ {
		nDic[i] = sortString(nDic[i])
	}
	for i := 0; i < len(nOutput); i++ {
		nOutput[i] = sortString(nOutput[i])
	}
	return nDic, nOutput
}

func sortString(s string) string {
	sArr := strings.Split(s, "")
	sort.Strings(sArr)
	s = strings.Join(sArr, "")
	return s
}

// This function maps the chars into the corresponding numbers.
// There is an order which numbers have to be found! starting with 1, 4, 7 and 8
// and then 0, 6 and 9. Lastly 2, 3 and 5
func sortNumbers(nDic []string) []string {
	sortedArr := make([]string, 10)

	// Filter numbers 1, 4, 7 and 8
	for _, val := range nDic {
		switch len(val) {
		case 2:
			sortedArr[1] = val

		case 3:
			sortedArr[7] = val

		case 4:
			sortedArr[4] = val

		case 7:
			sortedArr[8] = val
		}
	}

	// Find numbers 0, 6 and 9
	for _, val := range nDic {
		if len(val) == 6 {
			sortedArr[0], sortedArr[6], sortedArr[9] = sortSixLength(nDic, sortedArr[1], sortedArr[4])
		}
	}

	// Find remaining numbers (2, 3 and 5)
	for _, val := range nDic {
		if len(val) == 5 {
			sortedArr[2], sortedArr[3], sortedArr[5] = sortFiveLength(nDic, sortedArr[1], sortedArr[9])
		}
	}

	return sortedArr
}

// This function uses the existing numbers 1 and 4 to see the diferences from 0, 6 and 9.
// Number 6 shouldn't have a character from 1.
// Number 0 shouldn't have a character from 4.
// Number 9 is found by exclusion.
func sortSixLength(nDic []string, num1, num4 string) (string, string, string) {

	var num0 string
	var num6 string
	numbers := filterByLen(6, nDic)

	// // find 6 using number 1
	num6, numbers = findAndFilter(num6, num1, numbers)
	// Find 0 using number 4
	num0, numbers = findAndFilter(num0, num4, numbers)

	return num0, num6, numbers[0]
}

func findAndFilter(numToFind, num string, numbers []string) (string, []string) {
	for _, char := range num {
		for idx, num := range numbers {
			if !strings.Contains(num, string(char)) {
				numToFind = num
				// remove number from further search
				numbers = append(numbers[:idx], numbers[idx+1:]...)
			}
		}
	}
	return numToFind, numbers
}

// This function will use the numbers 1 and 9 to differentiate 2, 3 and 5
// 3 should have all chars as number 1
// 5 should have all chars of number 9 except for 1
// 2 is found by exclusion
func sortFiveLength(nDic []string, num1, num9 string) (string, string, string) {

	var num2 string
	var num3 string
	var num5 string
	numbers := filterByLen(5, nDic)

	// Numbers 2 and 5
	var unmappedNums []string = make([]string, 0)

	// Find 3 using number 1
	// Number 3 has to have all characters from number 1
	for _, char := range num1 {
		for idx, num := range numbers {
			if !strings.Contains(num, string(char)) {
				unmappedNums = append(unmappedNums, num)
				numbers = append(numbers[:idx], numbers[idx+1:]...)
			}
		}
	}
	num3 = numbers[0]

	// All chars from nummber 9 should be in 5, so if there is an unmatching char, it's number 2
	num2 = unmappedNums[1]
	num5 = unmappedNums[0]
	for _, char := range unmappedNums[0] {
		if !strings.Contains(num9, string(char)) {
			num2 = unmappedNums[0]
			num5 = unmappedNums[1]
			break
		}
	}

	return num2, num3, num5
}

func filterByLen(length int, nDic []string) []string {
	var numbers []string = make([]string, 0)
	for _, num := range nDic {
		if len(num) == length {
			numbers = append(numbers, num)
		}
	}
	return numbers
}

func translateAndCalculate(nDic []string, nOutput []string) int {
	numStr := ""

	for _, num1 := range nOutput {
		for idx, num2 := range nDic {
			if num1 == num2 {
				numStr += strconv.Itoa(idx)
			}
		}
	}

	num, err := strconv.Atoi(numStr)
	if err != nil {
		log.Fatalln("Error parsing the number -> ", err)
		return -1
	}
	return num
}
