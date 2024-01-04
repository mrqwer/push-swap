package main

import "fmt"

func sortSmallThreeNumbers(stackA *Stack) {
	biggestNode := getBiggest(stackA)

	if stackA.head == biggestNode {
		rotate(stackA)
		fmt.Println("ra")
	} else if stackA.head.nextNode == biggestNode {
		reverseRotate(stackA)
		fmt.Println("rra")
	}

	if stackA.head.value > stackA.head.nextNode.value {
		swap(stackA)
		fmt.Println("sa")
	}
}

func reverseSortSmallThreeNumbers(stackA *Stack) {
	smallestNode := getSmallest(stackA)

	if stackA.head == smallestNode {
		reverseRotate(stackA)
		fmt.Println("rra")
	} else if stackA.head.nextNode == smallestNode {
		rotate(stackA)
		fmt.Println("ra")
	}

	if stackA.head.value < stackA.head.nextNode.value {
		swap(stackA)
		fmt.Println("sa")
	}
}

func sortSmallFiveNumbers(stackA *Stack, stackB *Stack) {
	for stackA.length > 3 {
		initConfigurations(stackA, stackB)
		makeRotations(stackA, getSmallest(stackA))
		push(stackA, stackB)
		fmt.Println("pb")
	}
}
