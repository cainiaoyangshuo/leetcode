/**
 * @desc
 * @date 2022/11/10
 * @user yangshuo
 */

package channel

import (
	"fmt"
	"time"
)

func MultiProducer()  {
	//producerNum := 2
	//consumerNum := 1
	chanA := make(chan int)
	//var wg sync.WaitGroup
	//wg.Add(producerNum)

	//c := 0
	//for i := 0; i < n; i++ {
	//	go func() {
	//		chanA <- c
	//		c++
	//		wg.Done()
	//	}()
	//}
	go func() {
		for i := 0; i < 4; i++{
			fmt.Printf("producer1: %d\n", i)
			chanA <- i
		}

	}()
	go func() {
		for i := 0; i < 4; i++{
			fmt.Printf("producer2: %d\n", i)
			chanA <- i
		}
	}()

	go func() {
		for v := range chanA {
			fmt.Printf("consumer: %d\n", v)

		}
		//select {
		//case val := <- chanA :
		//	fmt.Println(val)
		//	break
		//default:
		//	fmt.Println("default")
		//}
	}()

	time.Sleep(time.Second*2)
	//wg.Wait()
	//fmt.Println(c)
}

func producer(str string, ch chan int)  {
	for i := 0; i < 4; i++ {
		fmt.Printf("%s: %d\n", str, i)
		ch <- i
	}
	//time.Sleep(time.Second*1)
}

func consumer(str string, ch chan int)  {
	for v := range ch {
		fmt.Printf("%s: %d\n", str, v)
	}
	fmt.Println("channel closed")
	//time.Sleep(time.Second*1)
}

func MultiProducerOneConsumer()  {
	chanInt := make(chan int)

	go producer("生产者1", chanInt)
	go producer("生产者2", chanInt)

	go consumer("consumer1", chanInt)

	time.Sleep(time.Second*2)
}