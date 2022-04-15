package _02removeKdigits

import (
	"fmt"
	"testing"
)

func Test_RemoveK(t *testing.T)  {
	tests := []struct{
		name  string
		args0 string
		args1 int
		want  string
	}{
		{
			name: "test1",
			args0: "10200",
			args1: 1,
			want: "200",
		},
		{
			name: "test2",
			args0: "112",
			args1: 1,
			want: "11",
		},
		{
			name: "test3",
			args0: "1234567890",
			args1: 9,
			want: "0",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := removeKdigits(tt.args0, tt.args1)
			fmt.Printf(" %v\n ", res)
			if res != tt.want {
				t.Errorf("removeKdigits = %v want %v", res, tt.want)
			}
		})
	}
}
