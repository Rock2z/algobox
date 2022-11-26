package _001

func twoSum(nums []int, target int) []int {
	for i, x := range nums {
		for j := i + 1; j < len(nums); j++ {
			if nums[j]+x == target {
				return []int{i, j}
			}
		}
	}
	return nil
}
