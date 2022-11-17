/**
 * @desc 成绩排名
 * @date 2022/11/17
 * @user yangshuo
 */

package main

import (
	"sort"
	"fmt"
)

func main() {
	var stus,tag int
	fmt.Scanln(&stus)
	fmt.Scanln(&tag)

	gradesAndStu := make(map[int][]string, 0)
	grades := make([]int,0)
	var (
		stu string
		g int
	)
	for i := 0; i < stus; i++ {
		fmt.Scanf("%s %d\n", &stu, &g)
		_,ok := gradesAndStu[g]
		if ok {
			gradesAndStu[g] = append(gradesAndStu[g], stu)
		} else {
			gradesAndStu[g] = []string{stu}
			grades = append(grades,g)
		}
	}

	sort.Ints(grades)

	if tag == 0 {
		for i := len(grades) -1; i >= 0; i-- {
			for _,v := range gradesAndStu[grades[i]] {
				fmt.Printf("%s %d\n",v, grades[i])
			}
		}
	} else {
		for i := 0; i < len(grades); i++ {
			for _,v := range gradesAndStu[grades[i]] {
				fmt.Printf("%s %d\n",v, grades[i])
			}
		}
	}

}
