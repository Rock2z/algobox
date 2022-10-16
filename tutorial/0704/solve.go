package _704

func search(nums []int, target int) int {
	left, right, mid := 0, len(nums), len(nums)/2
	for left < right {
		mid = (left + right) / 2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid
		} else {
			return mid
		}
	}
	return -1
}
