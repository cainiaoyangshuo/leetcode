package _55_minStack

import (
	"fmt"
	"testing"
)

func TestMinStack_GetMin(t *testing.T) {
	stack := Constructor()
	stack.Push(3)
	stack.Push(-2)
	stack.Push(10)
	stack.Push(-5)
	fmt.Printf("top : %v\n", stack.Top())
	fmt.Printf("min : %v\n", stack.GetMin())
	stack.Pop()
	fmt.Printf("top : %v\n", stack.Top())
	fmt.Printf("min : %v\n", stack.GetMin())
}
