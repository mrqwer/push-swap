package main

func push(src *Stack, dest *Stack) {
	if src == nil || src.head == nil || dest == nil {
		return
	}

	nodeToPush := src.head
	src.head = src.head.nextNode

	if src.head != nil {
		src.head.prevNode = nil
	}
	nodeToPush.nextNode = nil
	nodeToPush.prevNode = nil

	if dest.head == nil {
		dest.head = nodeToPush
		dest.tail = nodeToPush
	} else {
		nodeToPush.nextNode = dest.head
		dest.head.prevNode = nodeToPush
		dest.head = nodeToPush
	}

	src.length--
	dest.length++

}
