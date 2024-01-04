package main

func reverse_rotate(stack *Stack) {
	if stack == nil || stack.head == nil || stack.head.nextNode == nil {
		return
	}

	lastNode := getLastNode(stack)

	lastNode.prevNode.nextNode = nil
	lastNode.nextNode = stack.head
	stack.head.prevNode = lastNode
	stack.head = lastNode
	stack.head.prevNode = nil
	stack.tail = lastNode.prevNode
}
