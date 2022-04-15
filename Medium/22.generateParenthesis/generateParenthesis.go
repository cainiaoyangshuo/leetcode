package _2_generateParenthesis
/*
括号生成
https://leetcode-cn.com/problems/generate-parentheses/
回溯、dp
 */

func GenerateParenthesis(n int) []string {
	res := []string{}

	var dfs func(l,r int, path string)
	dfs = func(l, r int, path string) {
		// len==2n即为合法解
		if n * 2 == len(path) {
			res = append(res, path)
			return
		}

		if l > 0 {
			dfs(l-1, r, path+"(")
		}

		if l < r {
			dfs(l, r-1, path+")")
		}
	}

	dfs(n, n, "")
	return res
}
