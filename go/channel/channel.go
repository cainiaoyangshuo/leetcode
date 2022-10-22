/**
 * @desc
 * @date 2022/10/13
 * @user yangshuo
 */

package channel

//单协程通信
func Main()  {
	//无缓冲channel
	c := make(chan string)

	go count(5, "黑🐂", c)

	for msg := range c {
		println(msg)
	}
}

func count(n int, name string, c chan string) {
	for i := 0; i < n; i++ {
		c <- name
	}
	close(c)
}

// ===================

//多子任务管理
func Main1()  {
	//无缓冲channel
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for  {
			c1 <- "红"
		}
	}()
	
	go func() {
		for  {
			c2 <- "黑"
		}
	}()

	for  {
		println(<- c1)
		println(<- c2)
	}
}