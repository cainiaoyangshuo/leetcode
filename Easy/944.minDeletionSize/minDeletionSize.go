/**
 * @desc 944.删列造序  返回非递增的列的个数
 * @date 2022/5/12
 * @user yangshuo
 */

package _44_minDeletionSize


func minDeletionSize(strs []string) int {
	res := 0
	for i := 0; i < len(strs[0]); i++ {
		for j := 0; j < len(strs)-1; j++ {
			if strs[j][i] > strs[j+1][i] {
				res++
				break
			}
		}
	}

	return res
}
