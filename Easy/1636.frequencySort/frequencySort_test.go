/**
 * @desc
 * @date 2022/10/16
 * @user yangshuo
 */

package _636_frequencySort

import (
	"reflect"
	"testing"
)

func Test_frequencySort(t *testing.T)  {
	tests := []struct{
		name string
		arg0  []int
		want []int
	}{
		{
			name: "test1",
			arg0: []int{1,1,2,2,2,3},
			want: []int{3,1,1,2,2,2},
		},
		{
			name: "test2",
			arg0: []int{2,3,1,3,2},
			want: []int{1,3,3,2,2},
		},
		{
			name: "test3",
			arg0: []int{-1,1,-6,4,5,-6,1,4,1},
			want: []int{5,-1,4,4,-6,-6,1,1,1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := frequencySort(tt.arg0)
			if !reflect.DeepEqual(res, tt.want) {
				t.Errorf("name:%v, res:%v, want %v", tt.name, res, tt.want)
			}
		})
	}
}
