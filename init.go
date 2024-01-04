package main

import "math"

func initStackA(numbers []int) *Stack {
	stack := &Stack{
		length: len(numbers),
		name:   'a',
	}

	for _, num := range numbers {
		node := &Node{
			value: num,
		}
		if stack.head == nil {
			stack.head = node
			stack.tail = node
		} else {
			node.prevNode = stack.tail
			stack.tail.nextNode = node
			stack.tail = node
		}
	}

	return stack
}

func setPositions(stack *Stack) {
	if stack == nil {
		return
	}

	centerLine := stack.length / 2
	current := stack.head
	i := 0
	for current != nil {
		current.current_position = i
		if i <= centerLine {
			current.above_median = true
		} else {
			current.above_median = false
		}
		current = current.nextNode
		i++
	}
}

func setTargetNodes(stackA *Stack, stackB *Stack) {
	var currentA *Node
	var targetNode *Node
	var bestMatch int
	currentB := stackB.head

	for currentB != nil {
		bestMatch = math.MaxInt
		currentA = stackA.head

		for currentA != nil {
			if (currentA.value > currentB.value) && (currentA.value < bestMatch) {
				bestMatch = currentA.value
				targetNode = currentA
			}
			currentA = currentA.nextNode
		}
		if math.MaxInt == bestMatch {
			currentB.targetNode = getSmallest(stackA)
		} else {
			currentB.targetNode = targetNode
		}
		currentB = currentB.nextNode
	}
}

func setPrices(stackA *Stack, stackB *Stack) {
	currentB := stackB.head

	for currentB != nil {
		currentB.push_price = currentB.current_position

		if !currentB.above_median {
			currentB.push_price = stackB.length - currentB.current_position
		}

		if currentB.targetNode.above_median {
			currentB.push_price += currentB.targetNode.current_position
		} else {
			currentB.push_price += stackA.length - currentB.targetNode.current_position
		}

		currentB = currentB.nextNode
	}
}

func setCheapestNode(stack *Stack) {
	var bestMatchValue int = math.MaxInt
	var bestMatchNode *Node = &Node{}

	if stack == nil {
		return
	}

	current := stack.head

	for current != nil {
		if current.push_price < bestMatchValue {
			bestMatchValue = current.push_price
			bestMatchNode = current
		}

		current = current.nextNode
	}
	bestMatchNode.cheapest = true
}

func initConfigurations(stackA *Stack, stackB *Stack) {
	setPositions(stackA)
	setPositions(stackB)
	setTargetNodes(stackA, stackB)
	setPrices(stackA, stackB)
	setCheapestNode(stackB)
}
