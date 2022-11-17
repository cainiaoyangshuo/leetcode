/**
 * @desc 24点运算
 * @date 2022/11/17
 * @user yangshuo
 */

/**
特别注意：不要考虑运算优先级，不要考虑运算优先级，不要考虑运算优先级。
 */
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var value = map[string]int{
	"A":  1,
	"2":  2,
	"3":  3,
	"4":  4,
	"5":  5,
	"6":  6,
	"7":  7,
	"8":  8,
	"9":  9,
	"10": 10,
	"J":  11,
	"Q":  12,
	"K":  13,
}

var symbols = []string{"+", "-", "*", "/"}

// 最终解
var solution = ""

// 输出算式的运算顺序从左至右，不包含括号，如1+2+3*4的结果为24，2 A 9 A不能变为(2+1)*(9-1)=24
// 注意：没有优先级，没有优先级，没有优先级
func eval(exp []string) int {
	res := value[exp[0]]
	for i := 1; i < len(exp); i += 2 {
		switch exp[i] {
		case "+":
			res += value[exp[i+1]]
		case "-":
			res -= value[exp[i+1]]
		case "*":
			res *= value[exp[i+1]]
		case "/":
			res /= value[exp[i+1]]
		}
	}
	return res
}

func search(pokers []string, visited []bool, ans []string) {
	// 已有解，直接返回
	if len(solution) != 0 {
		return
	}

	// 检查最终解
	if len(ans) == 7 && eval(ans) == 24 {
		solution = strings.Join(ans, "")
		return
	}

	// 构造数字
	if len(ans)&1 == 0 {
		for i := 0; i < len(pokers); i++ {
			if visited[i] {
				continue
			}
			visited[i] = true
			ans = append(ans, pokers[i])
			search(pokers, visited, ans)
			ans = ans[:len(ans)-1]
			visited[i] = false
		}
	} else {
		// 构造符号
		for _, symbol := range symbols {
			ans = append(ans, symbol)
			search(pokers, visited, ans)
			ans = ans[:len(ans)-1]
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := strings.ToUpper(scanner.Text())

	// 检查大小王
	if strings.Contains(input, "JOKER") {
		fmt.Println("ERROR")
		return
	}

	ans := []string{}
	solution = ""
	visited := make([]bool, 4)
	search(strings.Split(input, " "), visited, ans)
	if len(solution) == 0 {
		solution = "NONE"
	}
	fmt.Println(solution)
}