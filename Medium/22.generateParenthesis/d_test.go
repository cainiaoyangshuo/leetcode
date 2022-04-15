package _2_generateParenthesis

import (
	"fmt"
	"testing"
)

func TestGenerateParenthesis(t *testing.T) {
	n := 3
	res := GenerateParenthesis(n)
	fmt.Printf("%v\n", res)
}
