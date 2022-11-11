/**
 * @desc
 * @date 2022/10/17
 * @user yangshuo
 */

package slice

import "fmt"

func _go() {
	v := []int{1,2,3}
	for i := range v {//i为下标
		fmt.Println(i)
		v = append(v, i)
	}

	fmt.Printf("%+v\n", v)//1 2 3 0 1 2 原因是range切片只取最初的值，后面增加的不会遍历。
}

func SliceAsParams(s []int) {
	s = append(s, 1)
	fmt.Print(s)

	s[1] = 0
	fmt.Print(s)
}