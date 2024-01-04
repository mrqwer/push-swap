package main

import "fmt"

func sortComplex(stackA *Stack, stackB *Stack) {
	switch stackA.length {
	case 5:
		sortSmallFiveNumbers(stackA, stackB)
	default:
		for stackA.length > 3 {
			push(stackA, stackB)
			fmt.Println("pb")
		}
	}
	sortSmallThreeNumbers(stackA)
	for stackB.head != nil {
		if isStackSorted(stackA) && isStackSorted(stackB) && stackB.length == 2 && stackA.head.value > stackB.head.nextNode.value {
			swap(stackB)
			fmt.Println("sb")
			for stackB.head != nil {
				push(stackB, stackA)
				fmt.Println("pa")
			}
			return
		}
		initConfigurations(stackA, stackB)
		makeMovements(stackA, stackB)
	}
	fmt.Println("hello")
	setPositions(stackA)
	smallest := getSmallest(stackA)
	if smallest.above_median {
		for stackA.head != smallest {
			rotate(stackA)
			fmt.Println("ra")
		}
	} else {
		for stackA.head != smallest {
			reverse_rotate(stackA)
			fmt.Println("rra")
		}
	}
}
