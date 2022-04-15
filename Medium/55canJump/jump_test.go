package _5canJump

import (
	"fmt"
	"testing"
)

func Test_Jump(t *testing.T)  {
	tests := []struct{
		name string
		args []int
		want bool
	}{
		{
			name: "test1",
			args: []int{3,2,1,0,4},
			want: false,
		},
		{
			name: "test2",
			args: []int{0,2,3},
			want: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := canJumpI(test.args)
			fmt.Printf(" %v\n ", res)
			if res != test.want {
				t.Errorf("canJumpI = %v want %v", res, test.want)
			}
		})
	}
}
