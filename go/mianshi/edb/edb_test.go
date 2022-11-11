/**
 * @desc
 * @date 2022/11/5
 * @user yangshuo
 */

package edb

import (
	"fmt"
	"testing"
)

func TestPathfinding(t *testing.T) {
	maps := [][]string{
		{"0", "x", "0", "0"},
		{"0", "1", "0", "0"},
		{"1", "0", "y", "0"},
		{"0", "0", "0", "0"},
	}
	//maps := [][]string{
	//	{"0", "x", "0"},
	//	{"0", "1", "0"},
	//	{"1", "0", "y"},
	//}
	fmt.Println(maps)
	res := Pathfinding(maps)

	for _, oc := range res {
		fmt.Println(oc)
	}
}

func TestPathFinding(t *testing.T) {
	//grid3 := [][]string{
	//	{"0", "x", "0", "0"},
	//	{"0", "1", "0", "0"},
	//	{"1", "0", "0", "y"},
	//	{"0", "0", "1", "0"},
	//}
	grid3 := [][]string{
		{"0", "x", "0"},
		{"0", "1", "0"},
		{"1", "0", "y"},
	}
	fmt.Println(grid3)
	res := PathFinding(grid3)

	for _, oc := range res {
		fmt.Println(oc)
	}
}