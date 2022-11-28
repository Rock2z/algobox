package _239

func maxSlidingWindow(nums []int, k int) []int {
	ret := make([]int, 0, len(nums)-k)
	window := make([]int, 0, k)
	for i := 0; i < len(nums); i++ {
		if i >= k && nums[i-k] == window[0] {
			window = window[1:]
		}
		for len(window) > 0 && window[len(window)-1] < nums[i] {
			window = window[:len(window)-1]
		}
		window = append(window, nums[i])

		if i >= k-1 {
			ret = append(ret, window[0])
		}
	}
	return ret
}
