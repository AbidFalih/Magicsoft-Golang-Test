package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	input := scanner.Text()
	numStrings := strings.Split(input, ", ")

	numbers := []int{}
	i := 0
	max := 0
	for i < len(numStrings) {
		if i == 0 {
			dummy := strings.Split(numStrings[i], "[")
			numStrings[i] = dummy[1]
		}

		if i == len(numStrings)-1 {
			dummy := strings.Split(numStrings[i], "]")
			numStrings[i] = dummy[0]
		}

		dummy, _ := strconv.Atoi(numStrings[i])
		numbers = append(numbers, dummy)
		// fmt.Println(numbers[i])

		if dummy > max {
			max = dummy
		}
		i++
	}

	// insertionsort(numbers, max)
	reverseInsertionsort(numbers, max)
}

func visualization(data []int, max int) {
	for i := 0; i <= max; i++ {
		for j := 0; j < len(data); j++ {
			if i != max {
				if data[j] >= max-i {
					fmt.Print("|")
				} else {
					fmt.Print(" ")
				}
			} else {
				fmt.Print(data[j])
			}

			fmt.Print(" ")
		}

		fmt.Println()
	}

	fmt.Println()
}

func insertionsort(data []int, max int) {
	visualization(data, max)

	var n = len(data)
	for i := 1; i < n; i++ {
		for j := i; j > 0; j-- {
			if data[j-1] > data[j] {
				data[j-1], data[j] = data[j], data[j-1]
			}
		}

		visualization(data, max)
	}
}

func reverseInsertionsort(data []int, max int) {
	visualization(data, max)

	var n = len(data)
	for i := 1; i < n; i++ {
		for j := i; j > 0; j-- {
			if data[j-1] < data[j] {
				data[j-1], data[j] = data[j], data[j-1]
			}
		}

		visualization(data, max)
	}
}
