package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	validInputs := make([]int, 0)
	lessers := 0
	greaters := 0
	var intPointer *int
	var last = 0

	windowLimit := 3

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter depths:")

	for {
		scanner.Scan()
		depth := scanner.Text()

		if len(depth) != 0 {
			value, err := strconv.Atoi(depth)

			if err != nil {
				continue
			}

			validInputs = append(validInputs, value)

			if len(validInputs) == windowLimit {
				var currentSum = sum(validInputs)
				validInputs = validInputs[1:]

				if intPointer == nil {
					intPointer = &last
					last = currentSum
					continue
				}

				if currentSum > last {
					greaters++
				}

				if currentSum < last {
					lessers++
				}

				last = currentSum
			}

		} else {
			break
		}
	}

	result := fmt.Sprintf("lessers: %[1]d, greaters: %[2]d", lessers, greaters)
	fmt.Println(result)
}

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}
