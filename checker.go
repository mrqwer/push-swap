//go:build checker
// +build checker

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var INSTRUCTIONS = map[string]func(stacks ...*Stack){
	"pa": func(stacks ...*Stack) {
		b := stacks[0]
		a := stacks[1]
		push(b, a)
	},
	"pb": func(stacks ...*Stack) {
		a := stacks[0]
		b := stacks[1]
		push(a, b)
	},
	"sa": func(stacks ...*Stack) {
		a := stacks[0]
		swap(a)
	},
	"sb": func(stacks ...*Stack) {
		b := stacks[0]
		swap(b)
	},
	"ra": func(stacks ...*Stack) {
		a := stacks[0]
		rotate(a)
	},
	"rb": func(stacks ...*Stack) {
		b := stacks[0]
		rotate(b)
	},
	"rr": func(stacks ...*Stack) {
		a := stacks[0]
		b := stacks[1]
		rotate(a)
		rotate(b)
	},
	"rra": func(stacks ...*Stack) {
		a := stacks[0]
		reverseRotate(a)
	},
	"rrb": func(stacks ...*Stack) {
		b := stacks[0]
		reverseRotate(b)
	},
	"rrr": func(stacks ...*Stack) {
		a := stacks[0]
		b := stacks[1]
		reverseRotate(a)
		reverseRotate(b)
	},
}

func push_swap(stackA *Stack) {

	scanner := bufio.NewScanner(os.Stdin)
	operations := []string{}
	for scanner.Scan() {
		input := strings.TrimSpace(scanner.Text())
		input = strings.ReplaceAll(input, "\n", "")
		if input == "" {
			continue
		}
		if strings.ToLower(input) == "stop" {
			break
		}
		if _, ok := INSTRUCTIONS[input]; !ok {
			fmt.Println(input)
			fmt.Fprintf(os.Stderr, "Error\n")
			return
		}
		operations = append(operations, input)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading standard input: %s\n", err)
		return
	}

	stackB := &Stack{}

	for _, operation := range operations {
		switch operation {
		case "pa":
			INSTRUCTIONS[operation](stackB, stackA)
		case "pb", "rr", "rrr":
			INSTRUCTIONS[operation](stackA, stackB)
		case "sa", "ra", "rra":
			INSTRUCTIONS[operation](stackA)
		case "sb", "rb", "rrb":
			INSTRUCTIONS[operation](stackB)
		}
	}

	if stackB.head == nil && isStackSorted(stackA) {
		fmt.Println("Ok")
		return
	}
	fmt.Println("Ko")
}
