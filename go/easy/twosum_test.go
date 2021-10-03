package main

import (
	"fmt"
	"testing"
)

func TestTwosum(t *testing.T) {
	arr := []int{1,2,5,7}
	target := 7
	res := Twosum(arr, target)
	fmt.Printf("res : %+v\n", res)
}