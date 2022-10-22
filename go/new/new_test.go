/**
 * @desc
 * @date 2022/10/19
 * @user yangshuo
 */

package new

import (
	"fmt"
	"testing"
)

func Test_init(t *testing.T)  {
	var p int
	var v *int

	p = 1
	v = &p

	println(p)
	println(v)
}

func Test_a(t *testing.T)  {
	var a *int
	*a = 9
	println(*a)
}

func Test_b(t *testing.T)  {
	var a *int
	a = new(int)
	*a = 1
	fmt.Println(*a)
	fmt.Println(a)
}

func Test_c(t *testing.T)  {
	type Name struct {

		P string

	}

	var av *[5]int

	var iv *int

	var sv *string

	var tv *Name

	var m *map[int]int


	av = new([5]int)

	fmt.Println(*av) //[0 0 0 0 0 0]

	iv = new(int)

	fmt.Println(*iv) // 0

	sv = new(string)

	fmt.Println(*sv) //

	tv = new(Name)

	fmt.Println(*tv) //{}

	m = new(map[int]int)

	fmt.Println(*m) //map[]
}