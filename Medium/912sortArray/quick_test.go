package _12sortArray

import (
	"fmt"
	"testing"
)

/**
 * @desc
 * @date 2022/4/18
 * @user yangshuo
 */

func Test_quickSort(t *testing.T)  {
	tests := []struct{
		name string
		args []int
		want []int
	}{
		{
			name: "test1",
			args: []int{5,1,1,2,0,0},
			want: []int{0,0,1,1,2,5},
		},
	}


	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := QuickSort(tt.args)
			fmt.Printf(" %v\n ", res)
			//if res != tt.want {
			//	t.Errorf("name:%v, arg:%v, res:%v, want %v", tt.name, tt.args0, res, tt.want)
			//}
		})
	}
}

func TestQuickSort2(t *testing.T) {
	tests := []struct{
		name string
		args []int
		want []int
	}{
		{
			name: "test1",
			args: []int{5,1,1,2,0,0},
			want: []int{0,0,1,1,2,5},
		},
		{
			name: "test2",
			args: []int{5,2,3,1},
			want: []int{1,2,3,5},
		},
	}


	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := QuickSort2(tt.args)
			fmt.Printf(" %v\n ", res)
			//if res != tt.want {
			//	t.Errorf("name:%v, arg:%v, res:%v, want %v", tt.name, tt.args0, res, tt.want)
			//}
		})
	}
}