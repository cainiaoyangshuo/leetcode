package _101

func isUnique(astr string) bool {
	mapStr := make(map[rune]int, 0)

	for _, s := range astr {
		if _, ok := mapStr[s]; ok {
			return false
		}
		mapStr[s] += 1
	}

	return true
}
