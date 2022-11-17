/**
 * @desc
 * @date 2022/11/12
 * @user yangshuo
 */

package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main()  {
	for {
		n:=0
		_,err:=fmt.Scan(&n)
		if err !=nil{
			break
		}

		li:=make([]int ,n)
		//赋值
		for i:=0;i<n;i++{
			num:=0
			fmt.Scan(&num)
			li[i]=num
		}
		fmt.Scan(&n)
		lr:=make(map[int]int)
		//赋值
		for i:=0;i<n;i++{
			num:=0
			fmt.Scan(&num)
			lr[num]=0
		}
		lr1:=make([]int,0)
		for v,_:=range lr{
			lr1=append(lr1,v)
		}

		sort.Ints(lr1)

		sumnum:=0
		str:=""

		for i:=0;i<len(lr1);i++{
			num:=0
			str1 :=""
			for j:=0;j<len(li);j++{

				n1,n2:=strconv.Itoa(lr1[i]),strconv.Itoa(li[j])

				if strings.Contains(n2,n1){
					num++
					sumnum+=2
					str1+= strconv.Itoa(j) +" "+n2+" "
				}
			}
			if num!=0{
				sumnum+=2
				str1= strconv.Itoa(lr1[i])+" "+strconv.Itoa(num) +" "+str1
				str+= str1
			}

		}
		fmt.Printf("%d %s\n",sumnum,str)
	}
}
