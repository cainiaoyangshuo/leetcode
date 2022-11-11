/**
 * @desc
 * @date 2022/10/31
 * @user yangshuo
 */

package _defer

import "testing"

func TestDeferFunc1(t *testing.T) {
	res := DeferFunc1(1)
	println(res)
}

func TestDeferFunc2(t *testing.T) {
	res := DeferFunc2(1)
	println(res)
}

func TestDeferFunc3(t *testing.T) {
	res := DeferFunc3(1)
	println(res)
}
