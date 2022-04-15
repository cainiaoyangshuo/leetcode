package _28_selfDividingNumbers

func SelfDividingNumbers(left int, right int) []int {
	res := []int{}
	for i := left; i <= right; i++ {
		if isDividing(i) {
			res = append(res, i)
		}
	}
	return res
}

func isDividing(num int) bool {

	for i := num; i > 0; i /= 10 {
		// 取个位数
		d := i%10
		// 0不能做被除数
		if d == 0 || num % d != 0 {
			return false
		}
	}
	return true
}