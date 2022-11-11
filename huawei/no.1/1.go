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
)

func main() {

	inputReader := bufio.NewReader(os.Stdin)
	input, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	arr := strings.Split(input, " ")
	last := strings.Trim(arr[len(arr)-1], "\n")
	fmt.Println(len(last))
}
