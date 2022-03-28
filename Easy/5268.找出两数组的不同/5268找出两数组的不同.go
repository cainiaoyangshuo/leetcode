package _268_找出两数组的不同

func findDifference(nums1 []int, nums2 []int) [][]int {
	map1 := map[int]int{}
	map2 := map[int]int{}

	r1, r2 := make([]int, 0), make([]int, 0)

	for i := range nums1 {
		map1[nums1[i]] = 1
	}

	for j := range nums2 {
		map2[nums2[j]] = 1
	}

	for k1 := range map1 {
		if _, ok := map2[k1]; !ok {
			r1 = append(r1, k1)
		}
	}

	for k2 := range map2 {
		if _, ok := map1[k2]; !ok {
			r2 = append(r2, k2)
		}
	}

	return [][]int{r1, r2}
}