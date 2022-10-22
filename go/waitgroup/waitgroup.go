/**
 * @desc
 * @date 2022/10/21
 * @user yangshuo
 */

package waitgroup

import (
	"sync"
	"time"
)

func Main()  {

	var wg sync.WaitGroup
	//添加2个计数
	wg.Add(2)

	go func() {
		count(5, "黑")
		wg.Done()
	}()
	go func() {
		count(5, "白")
		wg.Done()
	}()

	wg.Wait()
}

func count(n int, name string) {
	for i := 0; i < n; i++ {
		println(name)
		time.Sleep(time.Millisecond * 200)
	}

}
