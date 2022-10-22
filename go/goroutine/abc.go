/**
 * @desc 协程打印abc。
 * @date 2022/8/26
 * @user yangshuo
 */

package goroutine

import "sync"

func printabc() {
	var ch1, ch2, ch3 = make(chan struct{}), make(chan struct{}), make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(3)

	go func(s string) {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			<- ch1
			println(s)
			ch2 <- struct{}{}
		}
		<- ch1
	}("A")
	go func(s string) {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			<- ch2
			println(s)
			ch3 <- struct{}{}
		}
	}("B")
	go func(s string) {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			<- ch3
			println(s)
			ch1 <- struct{}{}
		}
	}("C")
	ch1 <- struct{}{}

	wg.Wait()
}

func print_abc()  {
	//var ch1, ch2, ch3 = make(chan struct{}), make(chan struct{}), make(chan struct{})
	go func(s string) {
		println(s)
	}("a")
	go func(s string) {
		println(s)
	}("b")
	go func(s string) {
		println(s)
	}("c")
}