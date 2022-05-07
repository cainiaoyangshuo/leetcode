/**
 * @desc
 * @date 2022/5/3
 * @user yangshuo
 */
package _4spiralOrder

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_spiralOrder(t *testing.T)  {
	tests := []struct{
		name string
		args [][]int
		want []int
	}{
		{
			name: "test1",
			args: [][]int{{2,5},{8,4},{0,-1}},
			want: []int{2,5,4,-1,0,8},
		},
		{
			name: "test2",
			args: [][]int{{1,2,3},{4,5,6},{7,8,9}},
			want: []int{1,2,3,6,9,8,7,4,5},
		},
	}


	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := spiralOrder(tt.args)
			fmt.Printf(" %v\n ", res)
			if !reflect.DeepEqual(res, tt.want) {
				t.Errorf("name:%v, arg:%v, res:%v, want %v", tt.name, tt.args, res, tt.want)
			}
		})
	}
}
