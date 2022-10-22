/**
 * @desc
 * @date 2022/5/18
 * @user yangshuo
 */

package __findMedianSortedArrays

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_findMedianSortedArrays(t *testing.T)  {
	tests := []struct{
		name  string
		args0 []int
		args1 []int
		want  float64
	}{
		{
			name:  "test1",
			args0: []int{1,3},
			args1: []int{2},
			want: 2,
		},
		{
			name:  "test2",
			args0: []int{1,2},
			args1: []int{3,4},
			want: 2.5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := findMedianSortedArrays(tt.args0, tt.args1)
			fmt.Printf(" %v\n ", res)
			if reflect.DeepEqual(res, tt.want) {
				t.Errorf("name:%v, arg:%v, want %v", tt.name, tt.args0, tt.want)
			}
		})
	}
}
