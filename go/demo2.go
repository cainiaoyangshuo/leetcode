package _go

import "fmt"

type People2 interface {
	Show()
}

type Student struct{}

func (stu *Student) Show() {

}

func live() People2 {
	var stu *Student
	return stu
}

func main4() {
	res := live()
	fmt.Printf("res: %v\n", res)
	if res == nil {//!!！判断结构体是否为空不能这么判断
		fmt.Println("AAAAAAA")
	} else {
		fmt.Println("BBBBBBB")
	}
}