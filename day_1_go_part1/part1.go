package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	lessers := 0
	greaters := 0
	var intPointer *int
	var last = 0

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

			if intPointer == nil {
				intPointer = &last
				last = value
				continue
			}

			if value > *intPointer {
				greaters++
			}
			if value < *intPointer {
				lessers++
			}

			last = value
		} else {
			break
		}
	}

	result := fmt.Sprintf("lessers: %[1]d, greaters: %[2]d", lessers, greaters)
	fmt.Println(result)
}
