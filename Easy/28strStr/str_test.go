package _8strStr

import (
	"fmt"
	"testing"
)

func Test_strStr(t *testing.T)  {
	test := []struct{
		name string
		args []string
		want int
	}{
		{
			name: "1",
			args: []string{"mississippi", "issip"},
			want: 4,
		},
		{
			name: "2",
			args: []string{"mississippi", "issipi"},
			want: -1,
		},
	}

	for _, c := range test {

		t.Run(c.name, func(t *testing.T) {
			res := strStr(c.args[0], c.args[1])
			fmt.Printf(" %d\n ", res)
			if res != c.want {
				t.Errorf("strStr = %d want %d", res, c.want)
			}
		})

	}
}
