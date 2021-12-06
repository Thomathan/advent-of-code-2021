package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	validDirections := []string{"up", "down", "forward"}
	horizontal := 0
	vertical := 0
	aim := 0

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter direction:")

	for {
		scanner.Scan()
		direction := scanner.Text()

		if len(direction) == 0 {
			break
		}

		split := strings.Split(direction, " ")

		if len(split) != 2 {
			continue
		}

		if !isValidDirection(split[0], validDirections) {
			continue
		}

		value, err := strconv.Atoi(split[1])

		if err != nil {
			continue
		}

		if split[0] == "forward" {
			horizontal += value
			vertical += value * aim
		} else {
			verticalMultiplier := getVerticalMultiplier(split[0])

			aim += (verticalMultiplier * value)
		}
	}

	area := horizontal * vertical
	result := fmt.Sprintf("Final Position: %[1]d, %[2]d. Final Area %[3]d", horizontal, vertical, area)
	fmt.Println(result)
}

func isValidDirection(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func getVerticalMultiplier(direction string) int {
	if direction == "up" {
		return -1
	}
	return 1
}
