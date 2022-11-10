/**
 * @desc m producer
 * @date 2022/10/25
 * @user yangshuo
 */

package leetcode

import (
	"sync"
)

func Main()  {
	a(2,3)
}
func a(m, n int)  {
	var wg = sync.WaitGroup{}
	ch1 := make(chan int)
	wg.Add(m+n)

	for i := 0; i < m; i++ {
		go func() {

			for j := 0; j < 5; j++ {
				ch1 <- j
				println(j)
				wg.Done()
			}
		}()
	}

	for k := 0; k < n; k++ {
		go func() {
			for {
				v := <- ch1
				println(v)
				wg.Done()
			}
		}()
	}

	close(ch1)
	wg.Wait()
}