/**
 * @desc
 * @date 2022/4/22
 * @user yangshuo
 */
package _5threeSum

import (
	"fmt"
	"testing"
)

func Test_threeSum(t *testing.T)  {
	tests := []struct{
		name string
		args []int
		want [][]int
	}{
		{
			name: "test1",
			args: []int{-1,0,1,2,-1,-4},
			want: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
	}


	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := threeSum(tt.args)
			fmt.Printf(" %v\n ", res)
			//if res != tt.want {
			//	t.Errorf("name:%v, arg:%v, res:%v, want %v", tt.name, tt.args0, res, tt.want)
			//}
		})
	}
}
