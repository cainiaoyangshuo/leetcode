/**
 * @desc 学英语。
Jessi初学英语，为了快速读出一串数字，编写程序将数字转换成英文：
 * @date 2022/11/17
 * @user yangshuo
 */

package main

import (
	"fmt"
	"strings"
)

func main(){
	var n int
	fmt.Scan(&n)

	fmt.Println( EnglishNum(n) )
}

func EnglishNum (n int ) string {
	var (
		count int //记录当前计数单位的位置
		temp string  //记录当前数字转换成英文
	)
	res:=make([]string,0) //记录各个单词

	for n>0 {
		if (n%100)/10==1 {
			switch n%10 {
			case 0: res=append(res,"ten")
			case 1: res=append(res,"eleven")
			case 2: res=append(res,"twelve")
			case 3: res=append(res,"thirthteen")
			case 4: res=append(res,"fourteen")
			case 5: res=append(res,"fifteen")
			case 6: res=append(res,"sixteen")
			case 7: res=append(res,"seventeen")
			case 8: res=append(res,"eighteen")
			case 9: res=append(res,"nineteen")
			}
			n/=100
			if n==0 {
				break
			} else {
				goto flag1
			}

		}

		if n%10==0 && n%100==0 {
			n/=100
			res=append(res,[]string{"hundred",OneNum(n%10)}...)
			goto flag2
		}

		temp = OneNum(n%10)
		if temp != ""{
			res=append(res,temp)
		}
		n/=10 //100100
		if n==0 {
			break
		}
		switch n%10 {
		case 2: res=append(res,"twenty")
		case 3: res=append(res,"thirty")
		case 4: res=append(res,"forty")
		case 5: res=append(res,"fifty")
		case 6: res=append(res,"sixty")
		case 7: res=append(res,"seventy")
		case 8: res=append(res,"eighty")
		case 9: res=append(res,"ninety")
		}
		n/=10 //10010
		if n==0 {
			break
		}

	flag1:
		if n%10 > 0 {
			//fmt.Println("百位数字：",n%10)
			res=append(res,[]string{"and","hundred",OneNum(n%10)}...)
		}

	flag2:
		n/=10 //1001
		if n==0 {
			break
		} else {
			count++
			//fmt.Println("count=",count)
			//fmt.Println(res)
			switch count {
			case 1: res=append(res,"thousand")
			case 2: res=append(res,"million")
			case 3: res=append(res,"billion")
			}
		}

	}

	for i:=0;i<len(res)/2;i++{
		res[i], res[len(res)-1-i] = res[len(res)-1-i], res[i]
	}
	str := strings.Join(res," ")
	return str
}

func OneNum (a int) string {
	var res string
	switch a {
	case 1: res="one"
	case 2: res="two"
	case 3: res="three"
	case 4: res="four"
	case 5: res="five"
	case 6: res="six"
	case 7: res="seven"
	case 8: res="eight"
	case 9: res="nine"
	}
	return res
}
