package _454

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	m1 := make(map[int]int, len(nums1)+len(nums2))
	for _, x := range nums1 {
		for _, y := range nums2 {
			m1[x+y] += 1
		}
	}
	m2 := make(map[int]int, len(nums3)+len(nums4))
	for _, x := range nums3 {
		for _, y := range nums4 {
			m2[x+y] += 1
		}
	}
	ret := 0
	for k, v := range m1 {
		if m2v, exist := m2[0-k]; exist {
			ret += v * m2v
		}
	}
	return ret
}
