package _28_selfDividingNumbers

import (
	"fmt"
	"testing"
)

func TestSelfDividingNumbers(t *testing.T) {
	left , right := 21, 22
	res := SelfDividingNumbers(left, right)
	fmt.Printf(": %v\n", res)
}
