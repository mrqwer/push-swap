package main

type Node struct {
	value            int
	current_position int
	push_price       int
	above_median     bool
	cheapest         bool
	targetNode       *Node
	nextNode         *Node
	prevNode         *Node
}

type Stack struct {
	head   *Node
	tail   *Node
	length int
	name   byte
}

func (s *Stack) GetLength() int {
	return s.length
}
