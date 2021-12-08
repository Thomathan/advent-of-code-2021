package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	calledNumbers := make([]int, 0)
	boards := make([][]int, 0)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter Inputs:")

	for {
		scanner.Scan()
		input := scanner.Text()

		if len(input) == 0 {
			continue
		}

		if input == "q" {
			break
		}

		if len(calledNumbers) == 0 {
			calledStrings := strings.Split(input, ",")
			calledNumbers = append(convertToInts(calledStrings))
			continue
		}

		board := len(boards)

		if board == 0 || len(boards[board-1]) == 25 {
			boards = append(boards, make([]int, 0))
		} else {
			board = board - 1
		}

		inputs := strings.Split(input, " ")
		inputInts := convertToInts(inputs)

		boards[board] = append(boards[board], inputInts[:]...)
	}

	winner, number := getWinner(calledNumbers, boards)
	score := processWinner(winner, number)

	fmt.Println(score)
}

func convertToInts(values []string) []int {
	ints := make([]int, 0)
	for _, s := range values {
		if s == "" {
			continue
		}
		v, _ := strconv.Atoi(s)
		ints = append(ints, v)
	}
	return ints
}

func processWinner(winner []int, lastNumber int) int {
	score := 0
	for _, r := range winner {
		if r != -1 {
			score += r
		}
	}

	return score * lastNumber
}

func getWinner(numbers []int, boards [][]int) ([]int, int) {
	winner := make([]int, 0)
	lastNumber := -1
out:
	for _, n := range numbers {
		for _, board := range boards {
			for i, v := range board {
				if n == v {
					board[i] = -1
					// check for winner
					if checkForWinner(board) == true {
						winner = append(winner, board...)
						lastNumber = n
						break out
					}
				}
			}
		}
	}
	return winner, lastNumber
}

func checkForWinner(board []int) bool {
	isWinner := checkRows(board)

	if isWinner == true {
		return isWinner
	}

	return checkColumns(board)
}

func checkRows(board []int) bool {
	isWinner := false

	for i := 0; i < 25; i += 5 {
		if board[i] == -1 {
			marksInRow := 1
			for j := i + 1; j < i+5; j++ {
				if board[j] == -1 {
					marksInRow++
				} else {
					break
				}
			}

			if marksInRow == 5 {
				isWinner = true
			}
		}
	}

	return isWinner
}

func checkColumns(board []int) bool {
	numberOfColumns := 5
	columnsToCheck := make([]int, 0)

	for i := 0; i < numberOfColumns; i++ {
		if board[i] == -1 {
			columnsToCheck = append(columnsToCheck, i)
		}
	}

	if len(columnsToCheck) == 0 {
		return false
	}

	bingo := false

	for i := len(columnsToCheck) - 1; i >= 0; i-- {
		for j := 5; j <= 20; j += 5 {
			nextRow := columnsToCheck[i] + j

			if board[nextRow] != -1 {
				columnsToCheck = append(columnsToCheck[:i], columnsToCheck[i+1:]...)
				break
			}
		}
	}

	if len(columnsToCheck) > 0 {
		bingo = true
	}

	return bingo
}
