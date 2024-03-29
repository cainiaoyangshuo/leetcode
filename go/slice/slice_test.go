/**
 * @desc
 * @date 2022/10/17
 * @user yangshuo
 */

package slice

import (
	"fmt"
	"testing"
)

func Test_slice(t *testing.T)  {
	_go()
}

func TestSliceAsParams(t *testing.T) {
	a := []int{1,2}
	SliceAsParams(a)
	fmt.Print(a)
}

func TestChangeSlice(t *testing.T) {
	a := []int{1, 2}
	fmt.Printf("%p: ", &a)
	fmt.Println(a)
	//修改slice
	ChangeSlice(a)
	fmt.Printf("%p: ", &a)
	fmt.Println(a)
}
