package _5searchInsert

import (
	"fmt"
	"testing"
)

func Test_searchInsert(t *testing.T)  {
	tests := []struct{
		name string
		args []int
		args1 int
		want int
	}{
		{
			name: "test1",
			args: []int{1,3,5,6},
			args1: 2,
			want: 1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := searchInsert(test.args, test.args1)
			fmt.Printf(" %v\n ", res)
			if res != test.want {
				t.Errorf("name:%v, arg:%v, res:%v, want %v", test.name, test.args, res, test.want)
			}
		})
	}
}
