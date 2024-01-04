package main

func rotate(stack *Stack) {
	if stack == nil || stack.head == nil || stack.head.nextNode == nil {
		return
	}

	lastNode := getLastNode(stack)

	stack.head, lastNode.nextNode, stack.tail = stack.head.nextNode, stack.head, stack.head
	stack.head.prevNode, stack.tail.nextNode, lastNode.nextNode.prevNode = nil, nil, lastNode
}
