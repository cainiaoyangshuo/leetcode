/**
 * @desc 扑克牌大小
 * @date 2022/11/17
 * @user yangshuo
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	str,_ := reader.ReadString('\n')
	str = strings.TrimSuffix(str,"\n")


	p := strings.Split(str,"-")
	player1 := strings.Split(p[0]," ")
	player2 := strings.Split(p[1]," ")

	if len(player1)==2 && strings.Contains(p[0],"joker"){
		fmt.Println(p[0])
		return
	}else if len(player2)==2 && strings.Contains(p[1],"joker"){
		fmt.Println(p[1])
		return
	}

	if len(player1)==1||len(player1)==2||len(player1)==3||len(player1)==5{// 1不是炸
		if len(player2)==1||len(player2)==2||len(player2)==3||len(player2)==5{ //2不是炸
			if len(player1) == len(player2){ //可以比较大小
				if biger(player1[0],player2[0])==0{ //第一个赢了
					fmt.Println(p[0])
				}else{ //第二个赢了
					fmt.Println(p[1])
				}
			}else{ //输出Error
				fmt.Println("ERROR")
			}
		}
	}


	if len(player1)==4{
		if len(player2)==2||len(player2)==3||len(player2)==5{ //2不是炸,1赢了
			fmt.Println(p[0])
		}else if len(player2)==4{ //比较大小，大的赢
			if biger(player1[0],player2[0])==0{ //第一个赢了
				fmt.Println(p[0])
			}else{ //第二个赢了
				fmt.Println(p[1])
			}
		}
	}

	if len(player2)==4{
		if len(player1)==2||len(player1)==3||len(player1)==5{ //1不是炸,2赢了
			fmt.Println(p[1])
		}
	}

}


func biger(a,b string)int{ //0代表第一个大，1代表第二个大
	m := make(map[string]int)
	m["3"] = 1
	m["4"] = 2
	m["5"] = 3
	m["6"] = 4
	m["7"] = 5
	m["8"] = 6
	m["9"] = 7
	m["10"] = 8
	m["J"] = 9
	m["Q"] = 10
	m["K"] = 11
	m["A"] = 12
	m["2"] = 13
	m["joker"] = 14
	m["JOKER"] = 15

	if m[a]>m[b]{
		return 0
	}else{
		return 1
	}
}