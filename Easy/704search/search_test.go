/**
 * @desc
 * @date 2022/4/20
 * @user yangshuo
 */
package _04search

import (
	"fmt"
	"testing"
)

func Test_search(t *testing.T)  {
	tests := []struct{
		name string
		args []int
		args1 int
		want int
	}{
		{
			name: "test1",
			args: []int{-1,0,3,5,9,12},
			args1: 9,
			want: 4,
		},
	}


	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := search(tt.args, tt.args1)
			fmt.Printf(" %v\n ", res)
			if res != tt.want {
				t.Errorf("name:%v, arg:%v, res:%v, want %v", tt.name, tt.args, res, tt.want)
			}
		})
	}

}
