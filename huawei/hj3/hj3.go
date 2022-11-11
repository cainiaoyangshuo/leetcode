/**
 * @desc
 * @date 2022/11/9
 * @user yangshuo
 */

package main

import "fmt"

func main()  {
	var inputA,inputB int
	arr := []int{}
	n, _ := fmt.Scanln(&inputA)
	if n == 0 {

	}
	for {
		n, _ := fmt.Scan(&inputB)
		if n == 0 {
			break
		} else {
			arr = append(arr, inputB)
		}
	}

	bucket := make([]int, 500)
	for _, v := range arr {
		bucket[v] = v
	}

	for _, v := range bucket {
		if v > 0 {
			fmt.Println(v)
		}
	}
}
