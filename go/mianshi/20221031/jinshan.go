/**
 * @desc 有编号为1-100 的共100个任务，每个任务打印一次 自己的编号即可，请实现如何让100个任务并发执行，但是同一时刻最多运行10个任务
 * @date 2022/10/31
 * @user yangshuo
 */

package _0221031

import "sync"

func Print()  {
	count := 100
	max := 10

	var wg sync.WaitGroup
	//缓冲区大小控制并发数量
	NumChannel := make(chan struct{}, max)

	for i := 0; i < count; i++ {
		wg.Add(1)
		NumChannel <- struct{}{}

		go func(j int) {
			println(j)
			<- NumChannel
			defer wg.Done()
		}(i)
	}

	defer close(NumChannel)
	wg.Wait()
}
