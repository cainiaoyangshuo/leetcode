package _3romanToInt

//https://leetcode-cn.com/problems/roman-to-integer/submissions/
func RomanToInt(s string) int {
	num := 0
	preIsI, preIsX, preIsC := false, false, false

	for _, c := range s {
		n := string(c)
		switch n {
		case "I":
			num += 1
			preIsI = true
			preIsX = false
			preIsC = false
			break
		case "V":
			num += 5
			if preIsI {
				num -= 2
			}
			preIsI = false
			preIsX = false
			preIsC = false
			break
		case "X":
			num += 10
			if preIsI {
				num -= 2
			}
			preIsX = true
			preIsI = false
			preIsC = false
			break
		case "L":
			num += 50
			if preIsX {
				num -= 20
			}
			preIsI = false
			preIsX = false
			preIsC = false
			break
		case "C":
			num += 100
			if preIsX {
				num -= 20
			}
			preIsI = false
			preIsX = false
			preIsC = true
			break
		case "D":
			num += 500
			if preIsC {
				num -= 200
			}
			preIsI = false
			preIsX = false
			preIsC = false
			break
		case "M":
			num += 1000
			if preIsC {
				num -= 200
			}
			preIsI = false
			preIsX = false
			preIsC = false
			break
		}
	}

	return num
}

func RomanToIntMap(s string) int {
	romanInt := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	
	length := len(s)
	num := 0
	for i := 0; i < length; i++ {
		// 没到最后一个且比后一个小，则减去当前数值;其他情况都累加
		if i < length - 1 && romanInt[s[i]] < romanInt[s[i+1]] {
			num -= romanInt[s[i]]
		} else {
			num += romanInt[s[i]]
		}
	}
	return num
}