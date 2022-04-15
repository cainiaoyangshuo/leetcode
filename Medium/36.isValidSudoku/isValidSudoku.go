package _6_isValidSudoku

import "fmt"

func isValidSudoku(board [][]string) bool {
	mapRow, mapColumn, mapArea := map[int]map[string]int{}, map[int]map[string]int{}, map[string]map[string]int{}
	for i, row := range board {
		for j, val := range row {
			if "." == string(val) {
				continue
			}
			// 判断行
			if value, ok := mapRow[i]; ok {
				if _, ok = value[val]; ok {
					return false
				} else {
					mapRow[i][val] = 1
				}
			} else {
				mapRow[i] = map[string]int{val:1}
			}

			// 判断列
			if colVal, ok := mapColumn[j]; ok {
				if _, ok = colVal[val]; ok {
					return false
				} else {
					mapColumn[j][val] = 1
				}
			} else {
				mapColumn[j] = map[string]int{val:1}
			}

			// 判断九宫格
			x := j/3
			y := i/3
			areaKey := fmt.Sprint("%D%D", x, y)
			if areaVal, ok := mapArea[areaKey]; ok {
				if _, ok = areaVal[val]; ok {
					return false
				} else {
					mapArea[areaKey][val] = 1
				}
			} else {
				mapArea[areaKey] = map[string]int{val:1}
			}
		}
	}

	return true
}
