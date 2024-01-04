//go:build push_swap
// +build push_swap

package main

func push_swap(stackA *Stack) {
	switch stackA.length {
	case 2:
		swap(stackA)
	case 3:
		sortSmallThreeNumbers(stackA)
	default:
		stackB := &Stack{name: 'b'}
		sortComplex(stackA, stackB)
	}
}
