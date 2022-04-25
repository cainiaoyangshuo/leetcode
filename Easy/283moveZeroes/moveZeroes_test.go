/**
 * @desc
 * @date 2022/4/23
 * @user yangshuo
 */
package _83moveZeroes

import "testing"

func Test_moveZeroes(t *testing.T)  {
	tests := []struct{
		name string
		arg  []int
		want []int
	}{
		{
			name: "test1",
			arg: []int{1,2,3,1},
			want: []int{1,2,3,1},
		},
		{
			name: "test2",
			arg: []int{1,0,0},
			want: []int{1,0,0},
		},
		{
			name: "test3",
			arg: []int{0,1,0,3,12},
			want: []int{1,3,12,0,0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			moveZeroes(tt.arg)
		})
	}
}
