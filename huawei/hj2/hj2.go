/**
 * @desc
 * @date 2022/11/8
 * @user yangshuo
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main()  {
	inputReader := bufio.NewReader(os.Stdin)
	inputA, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}

	inputB, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}

	inputA = strings.Trim(inputA, "\n")
	inputB = strings.Trim(inputB, "\n")
	if len(inputB) > 1 {
		panic(fmt.Sprintf("inputB: %s is wrong", inputB))
	}

	m := make(map[string]int)

	for _, c := range inputA {
		if string(c) == " " {
			continue
		}
		if unicode.IsNumber(c) {
			m[string(c)]++
			continue
		}
		low := strings.ToLower(string(c))
		m[low]++
	}

	low := strings.ToLower(inputB)
	fmt.Println(m[low])
}
