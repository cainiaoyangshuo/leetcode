/**
 * @desc
 * @date 2022/10/31
 * @user yangshuo
 */

package _defer

import "fmt"

//defer和return 的顺序关系，先defer再return
func DeferFunc1(i int) (t int) {
	t = i
	defer func() {
		t += 3
	}()
	return t
}

//返回值没有定义，return的是 零值变量。相当于s := t return s
func DeferFunc2(i int) int {
	t := i
	defer func() {
		t += 3
	}()
	return t
}

//先为返回值赋值，然后执行defer，再return
func DeferFunc3(i int) (t int) {
	defer func() {
		t += i
	}()
	return 2
}
func main() {
	fmt.Println(DeferFunc1(1))
	fmt.Println(DeferFunc2(1))
	fmt.Println(DeferFunc3(1))
}
