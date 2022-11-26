package _349

func intersection(nums1 []int, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	ret := make([]int, len(nums1))
	m := make(map[int]int, len(nums1))
	for _, num := range nums1 {
		m[num] = 1
	}
	for _, num := range nums2 {
		if m[num] > 0 {
			ret = append(ret, num)
			m[num] -= 1
		}
	}
	return ret
}
