/**
 * @desc
 * @date 2022/10/19
 * @user yangshuo
 */

package make

import (
	"fmt"
	"testing"
)

func Test_a(t *testing.T)  {
	arr := make([]int, 5)
	fmt.Println(arr)

	m := make(map[int]int)
	println(m)

	cha := make(chan int)
	println(cha)
}
