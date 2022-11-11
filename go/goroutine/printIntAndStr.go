/**
 * @desc
 * @date 2022/10/21
 * @user yangshuo
 */

package goroutine

import (
	"fmt"
	"sync"
)

func Print()  {
	str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	NumChannel := make(chan struct{})
	StrChannel := make(chan struct{})

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		i := 1
		for {
			select {
			case <-NumChannel:
				fmt.Println(i)
				i++
				//通知字符串通道
				StrChannel <- struct{}{}
			default:
				break
			}
		}
	}()
	go func() {
		i := 0
		for {
			select {
			case <-StrChannel:
				//打印字母结束条件
				if i >= len(str)-2 {
					fmt.Println(string(str[i]))
					i++
					fmt.Println(string(str[i]))
					wg.Done()
					return
				}
				fmt.Println(string(str[i]))
				i++
				//通知数字通道
				NumChannel <- struct{}{}
			default:
				break
			}
		}
	}()

	NumChannel <- struct{}{}
	wg.Wait()
}
