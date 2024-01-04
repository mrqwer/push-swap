package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isClientInputValid(numberStrings []string) (bool, []int) {
	numbers := []int{}
	checkerHashMap := make(map[int]bool)
	for _, numberString := range numberStrings {
		number, err := strconv.Atoi(numberString)
		if err != nil || checkerHashMap[number] {
			return false, []int{}
		}
		checkerHashMap[number] = true
		numbers = append(numbers, number)
	}
	return true, numbers
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 || (len(args) == 1 && len(args[0]) == 0) {
		return
	}

	if len(args) > 1 {
		fmt.Fprintf(os.Stderr, "Error\n")
		return
	}

	numberStrings := strings.Split(args[0], " ")
	ok, numbers := isClientInputValid(numberStrings)
	if !ok {
		fmt.Fprintf(os.Stderr, "Error\n")
		return
	}
	stack := initStackA(numbers)
	if !isStackSorted(stack) {
		push_swap(stack)
	}
}
