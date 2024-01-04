package main

func swap(stack *Stack) {
	if stack == nil || stack.head == nil || stack.head.nextNode == nil {
		return
	}

	first := stack.head
	second := first.nextNode

	first.nextNode = second.nextNode
	first.prevNode = second
	second.prevNode = nil
	second.nextNode = first

	stack.head = second

	if first.nextNode != nil {
		first.nextNode.prevNode = first
	} else {
		stack.tail = first
	}
}
