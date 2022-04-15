package _3romanToInt

import (
	"fmt"
	"testing"
)

func TestRomanToInt(t *testing.T) {
	//s := []string{"III", "LVIII", "MCMXCIV"}
	test := []struct{
		name string
		args string
		want int
	}{
		{
			name: "test1",
			args: "III",
			want: 3,
		},
		{
			name: "test2",
			args: "LVIII",
			want: 58,
		},
		{
			name: "test3",
			args: "MCMXCIV",
			want: 1994,
		},
	}
	for _, c := range test {

		t.Run(c.name, func(t *testing.T) {
			res := RomanToInt(c.args)
			fmt.Printf(" %d\n ", res)
			if res != c.want {
				t.Errorf("RomanToInt = %d want %d", res, c.want)
			}
		})

	}
}
