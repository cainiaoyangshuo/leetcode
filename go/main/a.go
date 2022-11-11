/**
 * @desc m producer
 * @date 2022/10/25
 * @user yangshuo
 */

package main

import (
	"sync"
)

func main()  {
	a(2,3)
}
var wg = sync.WaitGroup{}
func a(m, n int)  {

	ch1 := make(chan int)
	wg.Add(10)


		go func() {

			for j := 0; j < 5; j++ {
				ch1 <- j
				println(j)
				wg.Done()
			}
		}()



		go func() {
			for {
				v := <- ch1
				println(v)
				wg.Done()
			}
		}()


	wg.Wait()
}