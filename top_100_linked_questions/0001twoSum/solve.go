package _001twoSum

func twoSum(nums []int, target int) []int {
	return solve2(nums, target)
}

func solve(nums []int, target int) []int {
	for i, num := range nums {
		complement := target - num
		loc := find(complement, nums, complement == num)
		if loc > 0 {
			return []int{i, loc}
		}
	}
	return []int{}
}

func find(a int, slice []int, ignoreTheFirst bool) int {
	for i, e := range slice {
		if e == a {
			if ignoreTheFirst {
				ignoreTheFirst = false
				continue
			}
			return i
		}
	}
	return -1
}

func solve2(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for i, num := range nums {
		complement := target - num
		if index, ok := m[num]; ok {
			return []int{index, i}
		}
		m[complement] = i
	}
	return []int{}
}
