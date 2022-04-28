/**
 * @desc
 * @date 2022/4/28
 * @user yangshuo
 */
package _5sortcolors

import (
	"fmt"
	"testing"
)

func Test_sortColors(t *testing.T)  {
	tests := []struct{
		name string
		args []int
		want [][]int
	}{
		{
			name: "test1",
			args: []int{2,0,2,1,1,0},
			want: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
	}


	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sortColors(tt.args)
			fmt.Printf("%v\n", tt.args)
			//if res != tt.want {
			//	t.Errorf("name:%v, arg:%v, res:%v, want %v", tt.name, tt.args0, res, tt.want)
			//}
		})
	}
}
