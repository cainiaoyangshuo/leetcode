/**
 * @desc
 * @date 2022/9/19
 * @user yangshuo
 */

package goroutine

import (
	"fmt"
	"sync"
)

var testMap  = map[string]string{}
func main() {
	go func() {
		for{
			_ = testMap["bar"]
		}
	}()
	go func() {
		for  {
			testMap["bar"] = "foo"
		}
	}()
	//select{}
}

func syncMap()  {
	var m sync.Map
	// 写
	m.Store("test", 1)
	m.Store(1, true)

	// 读
	val1, _ := m.Load("test")
	val2, _ := m.Load(1)
	fmt.Println(val1.(int))
	fmt.Println(val2.(bool))

	//遍历
	m.Range(func(key, value interface{}) bool {
		//....
		return true
	})

	//删除
	m.Delete("test")

}