/**
 * @desc
 * @date 2022/5/22
 * @user yangshuo
 */

package _50_intersect

import (
	"reflect"
	"testing"
)

func Test_Intersect(t *testing.T)  {
	tests := []struct{
		name string
		arg0  []int
		arg1  []int
		want []int
	}{
		{
			name: "test1",
			arg0: []int{1,2,2,1},
			arg1: []int{2,2},
			want: []int{2,2},
		},
		{
			name: "test2",
			arg0: []int{4,9,5},
			arg1: []int{9,4,9,8,4},
			want: []int{4,9},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := intersect(tt.arg0, tt.arg1)
			if !reflect.DeepEqual(res, tt.want) {
				t.Errorf("name:%v, arg:%v, want %v", tt.name, res, tt.want)
			}
		})
	}
}
