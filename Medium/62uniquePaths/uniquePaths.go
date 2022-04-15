package _2uniquePaths

/**
 不同路径
 https://leetcode-cn.com/problems/unique-paths/
 dp
 */

func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	// 因为第一行和第一列只有一种可能，先初始化为1
	// 初始化第一列
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}

	//初始化第一行
	for j := range dp[0] {
		dp[0][j] = 1
	}

	// 因为第一行第一列已经初始化过，所以从第二个开始遍历
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	// 最后一位即为组合总数
	return dp[m-1][n-1]
}
