/**
 * @desc
 * @date 2022/11/10
 * @user yangshuo
 */

package leetcode

import (
	"fmt"
	"sync"
)

func main()  {
	chanA := make(chan int)
	var wg sync.WaitGroup
	wg.Add(11)
	n := 10
	for i := 0; i < n; i++ {
		go func() {
			chanA <- 3
			wg.Done()
		}()
	}

	go func() {
		for  {
			select {
			case val := <- chanA :
				fmt.Println(val)
				break
			}
		}
		wg.Done()
	}()
	defer close(chanA)

	wg.Wait()

}
