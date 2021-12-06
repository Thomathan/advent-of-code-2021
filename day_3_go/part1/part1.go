package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	totalInputs := 0
	onCounts := make([]int, 0)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter binary:")

	for {
		scanner.Scan()
		input := scanner.Text()

		if len(input) == 0 {
			break
		}

		totalInputs++

		for i := 0; i < len(input); i++ {
			if input[i] == '1' {
				if len(onCounts) <= i {
					onCounts = append(onCounts, 1)
				} else {
					onCounts[i]++
				}
			}
		}

	}

	gamma, _ := strconv.ParseInt(buildCommon(totalInputs, onCounts, false), 2, 64)
	epsilon, _ := strconv.ParseInt(buildCommon(totalInputs, onCounts, true), 2, 64)

	power := gamma * epsilon

	result := fmt.Sprintf("gamma: %[1]d, epsilon %[2]d, power %[3]d", gamma, epsilon, power)
	fmt.Println(result)
}

func buildCommon(total int, onCounts []int, invert bool) string {
	result := ""
	multiplier := 1

	if invert {
		multiplier = -1
	}

	half := total / 2 * multiplier

	for i := 0; i < len(onCounts); i++ {
		counter := onCounts[i] * multiplier
		if counter > half {

			result = result + "1"
		} else {

			result = result + "0"
		}
	}

	return result
}
