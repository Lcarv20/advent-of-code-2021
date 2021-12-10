package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Board = [][]string

func main() {

	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatalln("something went wrong while opening the file")
	}
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	// get the numbers for the sorting
	scanner.Scan()
	bingo := strings.Split(scanner.Text(), ",")
	scanner.Scan()
	// this will let me start exactly at the boards skipping the empty lines
	// --

	var playersBoard map[int]Board = make(map[int]Board)
	var playerNum int = 0
	// Boards, each consisting of a 5x5 grid
	for scanner.Scan() {
		var board Board

		// Player Board
		for i := 0; i < 5; i++ {
			line := strings.Fields(scanner.Text())
			board = append(board, line)
			scanner.Scan()
		}
		playersBoard[playerNum] = board
		playerNum++
	}

	// playing the actual bingo
	// Rules.: The numbers are drawn from the bingo pocket, and the player that completes a row or a collumn
	// 		   first wins.

	// will hold a snapshot of a player board when that won
	var scoreBoard = make(map[int]Board)
	var winnerPlace = 0
	// holds the numbers that each player drawn when they won
	var lastBingoDrawn []string
	func() {
		for i := 0; i < len(bingo); i++ {
			bingoNum := bingo[i]
			winnerId, won := markBoards(playersBoard, bingoNum)
			if won {
				lastBingoDrawn = append(lastBingoDrawn, bingoNum)
				scoreBoard[winnerPlace] = playersBoard[winnerId]
				delete(playersBoard, winnerId)
				winnerPlace++
				// we check gain if any other board wins with that number
				i--
			}
		}
	}()

	// *** 1ST PROBLEM ***
	firstWinner := 0

	// ### 2ND PROBLEM ###
	lastWinner := len(scoreBoard) - 1

	fmt.Println(scoreBoard)
	fmt.Println(len(lastBingoDrawn))

	result1stPart := getResult(firstWinner, scoreBoard, lastBingoDrawn)
	result2ndPart := getResult(lastWinner, scoreBoard, lastBingoDrawn)

	fmt.Printf("First winner scored: %d\nLast winner scored: %d\n", result1stPart, result2ndPart)

}

// Check every board and mark (x) the number drawn on the bingo
func markBoards(playersBoard map[int]Board, bingoNum string) (int, bool) {
	var didWin bool
	for playerId, board := range playersBoard {
		for row, rowNums := range board {
			col, didScore := contains(rowNums, bingoNum)
			if didScore {
				playersBoard[playerId][row][col] = "x"

				didWin = checkAxis(board, row, col)
				if didWin {
					return playerId, true
				}
			}
		}
	}
	return -1, false
}

// Check if row and collumn are fulfilled
func checkAxis(playersBoard Board, row int, col int) bool {

	// check row
	rowMarks := 0
	for _, x := range playersBoard[row] {
		if x == "x" {
			rowMarks++
		}
	}
	if rowMarks == 5 {
		return true
	}

	// check collum
	colMarks := 0
	for _, y := range playersBoard {
		if y[col] == "x" {
			colMarks++
		}
	}
	return colMarks == 5
}

func contains(strArr []string, str string) (int, bool) {
	for idx, char := range strArr {
		if str == char {
			return idx, true
		}
	}
	return -1, false
}

func getResult(playerId int, scoreBoard map[int]Board, lastBingoDrawn []string) int {
	sum := 0
	for _, row := range scoreBoard[playerId] {
		for _, cell := range row {
			/*
				Some fields are marked with "x" so parsing errors are expected, and the default
				val of 0 is exactly what we want in the sum.
			*/
			val, _ := strconv.Atoi(cell)
			sum += val
		}
	}
	lastBingo, err := strconv.Atoi(lastBingoDrawn[playerId])
	if err != nil {
		log.Fatalln("Error parsing last drawn Bingo number")
	}
	// fmt.Println("Result is: ", sum*lastBingo)

	return sum * lastBingo
}
