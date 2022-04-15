package _29majorityElement

//先要知道，在任何数组中，出现次数大于该数组长度1/3的值最多只有两个。
func MajorityElement(nums []int) []int {
	res := []int{}

	a, b := 0, 0
	countA, countB := 0, 0
	for _, n := range nums {
		// 必须先判断count
		if countA > 0 && a == n {
			countA++
		} else if countB > 0 && b == n {
			countB++
		} else if countB == 0 {
			countB++
			b = n
		} else if countA == 0 {
			countA++
			a = n
		}  else {
			countA--
			countB--
		}
	}

	cA, cB := 0,0
	for _, n := range nums {
		if countA > 0 && n == a {
			cA++
			continue
		}

		if countB > 0 && n == b {
			cB++
		}
	}

	if cA > len(nums)/3 {
		res = append(res, a)
	}
	if cB > len(nums)/3 {
		res = append(res, b)
	}

	return res
}
