/**
 * @desc
 * @date 2022/5/12
 * @user yangshuo
 */

package _8_merge

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_merge(t *testing.T)  {
	tests := []struct{
		name  string
		args0 []int
		args1 int
		args2 []int
		args3 int
		want  []int
	}{
		{
			name:  "test1",
			args0: []int{1,0},
			args1: 1,
			args2: []int{2},
			args3: 1,
			want: []int{1,2},
		},
		{
			name:  "test2",
			args0: []int{1},
			args1: 1,
			args2: []int{},
			args3: 0,
			want: []int{1},
		},
		{
			name:  "test3",
			args0: []int{0},
			args1: 0,
			args2: []int{1},
			args3: 1,
			want: []int{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			merge(tt.args0, tt.args1, tt.args2, tt.args3)
			fmt.Printf(" %v\n ", tt.args0)
			if reflect.DeepEqual(tt.args0, tt.want) {
				t.Errorf("name:%v, arg:%v, want %v", tt.name, tt.args0, tt.want)
			}
		})
	}
}
