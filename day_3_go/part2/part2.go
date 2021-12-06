package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	allInputs := make([]string, 0)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter binary:")

	for {
		scanner.Scan()
		input := scanner.Text()

		if len(input) == 0 {
			break
		}

		allInputs = append(allInputs, input)
	}

	sort.Strings(allInputs)

	oxygen, _ := strconv.ParseInt(splitMost(allInputs, 0), 2, 64)
	co2, _ := strconv.ParseInt(splitLeast(allInputs, 0), 2, 64)

	result := fmt.Sprintf("oxygen %[1]d, co2 %[2]d, life support %[3]d", oxygen, co2, oxygen*co2)
	fmt.Println(result)
}

func splitMost(inputs []string, startingIndex int) string {
	if len(inputs) == 1 {
		return inputs[0]
	}

	index := sort.Search(len(inputs), func(i int) bool {
		return inputs[i][startingIndex] == '1'
	})

	if index > (len(inputs) / 2) {
		// more 0
		return splitMost(inputs[0:index], startingIndex+1)
	} else {
		// more 1
		return splitMost(inputs[index:], startingIndex+1)
	}
}

func splitLeast(inputs []string, startingIndex int) string {
	if len(inputs) == 1 {
		return inputs[0]
	}

	index := sort.Search(len(inputs), func(i int) bool {
		return inputs[i][startingIndex] == '1'
	})

	if index > (len(inputs) / 2) {
		// more 0
		return splitLeast(inputs[index:], startingIndex+1)
	} else {
		// more 1
		return splitLeast(inputs[0:index], startingIndex+1)
	}
}
