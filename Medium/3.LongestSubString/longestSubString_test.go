/**
 * @desc
 * @date 2022/11/4
 * @user yangshuo
 */

package __LongestSubString

import (
	"fmt"
	"testing"
)

func TestLongestSubString(t *testing.T) {
	case1 := "abcabcde"
	res := LongestSubString(case1)
	fmt.Println(res)
}
