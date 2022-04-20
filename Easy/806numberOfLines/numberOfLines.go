package _06numberOfLines

/**
 * @desc 806 写字符串所需的行数
 * @date 2022/4/20
 * @user yangshuo
 */

func numberOfLines(widths []int, s string) []int {
	// 每行最大单位
	maxWidth := 100
	// 需要的行数
	lines := 1
	// 每行宽度，满一行重新计算
	width := 0

	for _, c := range s {
		// need 当前字母需要的单位
		need := widths[c-'a']
		// 行宽累加
		width += need

		// 如果当前字母超过行最大宽度，则行数+1，宽度为当前字母所需宽度
		if width > maxWidth {
			lines++
			width = need
		}
	}

	return []int{lines, width}
}
