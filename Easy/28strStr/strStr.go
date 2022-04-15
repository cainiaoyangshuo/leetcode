package _8strStr

/**
 https://leetcode-cn.com/problems/implement-strstr/
 暴力枚举/kmp

 */
func strStr(haystack string, needle string) int {
	lenHay, lenNee := len(haystack), len(needle)
	if lenNee == 0 {
		return 0
	}

	p1 :=  0
	for ;p1 <= lenHay-lenNee;p1++ {
		isMatch := true
		for p2:=0;p2 < lenNee;p2++ {
			if haystack[p1+p2] != needle[p2] {
				isMatch = false
				break
			}
		}
		if isMatch {
			return p1
		}
	}

	// 没找到
	return -1
}


//todo kmp
