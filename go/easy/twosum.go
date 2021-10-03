package main

import "fmt"

func Twosum(nums []int, target int) []int {

	var vMap = make(map[int]int)
	res := []int{}
	for i, v := range nums{
		if _,ok := vMap[v]; ok  {
			res = []int{vMap[v], i}
			return res
		} else {
			vMap[target-v] = i
		}
	}

	return res
}

func main()  {
	arr := []int{1,2,5,7}
	target := 4
	res := Twosum(arr, target)
	fmt.Printf("res : %+v\n", res)
}