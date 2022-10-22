/**
 * @desc
 * @date 2022/8/29
 * @user yangshuo
 */

package goroutine

import (
	"fmt"
	"testing"
)

func TestPrintabc(t *testing.T)  {
	printabc()
}

func Test_Print_abc(t *testing.T)  {
	print_abc()
}

func Test_Slice(t *testing.T)  {
	var a  = [3]int{1, 2, 3}
	fmt.Println(a, a[1:2])
}