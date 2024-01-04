package main

import "fmt"

func isStackSorted(stack *Stack) bool {
	current := stack.head

	if current == nil || current.nextNode == nil {
		return true
	}

	for current.nextNode != nil {
		if current.value > current.nextNode.value {
			return false
		}
		current = current.nextNode
	}

	return true
}

func isStackReverseSorted(stack *Stack) bool {
	current := stack.head

	if current == nil || current.nextNode == nil {
		return true
	}

	for current.nextNode != nil {
		if current.value < current.nextNode.value {
			return false
		}
		current = current.nextNode
	}

	return true
}

func getBiggest(stack *Stack) *Node {
	if stack == nil || stack.head == nil {
		return nil
	}

	maxValue := stack.head.value
	maxNode := stack.head
	current := stack.head.nextNode

	for current != nil {
		if current.value > maxValue {
			maxValue = current.value
			maxNode = current
		}
		current = current.nextNode
	}

	return maxNode
}

func getLastNode(stack *Stack) *Node {
	if stack == nil || stack.head == nil {
		return nil
	}
	return stack.tail
}

func getSmallest(stack *Stack) *Node {
	minNode := stack.head
	current := stack.head

	for current != nil {
		if minNode.value > current.value {
			minNode = current
		}
		current = current.nextNode
	}

	return minNode
}

func makeRotations(stack *Stack, topNode *Node) {
	for stack.head != topNode {
		switch stack.name {
		case 'a':
			if topNode.above_median {
				rotate(stack)
				fmt.Println("ra")
			} else {
				reverseRotate(stack)
				fmt.Println("rra")
			}
		case 'b':
			if topNode.above_median {
				rotate(stack)
				fmt.Println("rb")
			} else {
				reverseRotate(stack)
				fmt.Println("rrb")
			}
		}
	}
}

func getCheapest(stack *Stack) *Node {
	if stack == nil || stack.head == nil {
		return nil
	}

	current := stack.head

	for current != nil {
		if current.cheapest {
			return current
		}
		current = current.nextNode
	}
	return nil
}

func reverseRotateBoth(stackA *Stack, stackB *Stack, cheapestNode *Node) {
	for stackA.head != cheapestNode.targetNode && stackB.head != cheapestNode {
		reverseRotate(stackA)
		reverseRotate(stackB)
		fmt.Println("rrr")
	}
	setPositions(stackA)
	setPositions(stackB)
}

func rotateBoth(stackA *Stack, stackB *Stack, cheapestNode *Node) {
	for stackA.head != cheapestNode.targetNode && stackB.head != cheapestNode {
		rotate(stackA)
		rotate(stackB)
		fmt.Println("rr")
	}
	setPositions(stackA)
	setPositions(stackB)
}

func makeMovements(stackA *Stack, stackB *Stack) {
	var cheapestNode *Node = getCheapest(stackB)

	if cheapestNode.above_median && cheapestNode.targetNode.above_median {
		rotateBoth(stackA, stackB, cheapestNode)
	} else if !(cheapestNode.above_median) && !(cheapestNode.targetNode.above_median) {
		reverseRotateBoth(stackA, stackB, cheapestNode)
	}

	makeRotations(stackB, cheapestNode)
	makeRotations(stackA, cheapestNode.targetNode)
	push(stackB, stackA)
	fmt.Println("pa")
}
