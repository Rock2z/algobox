package _977

func sortedSquares(nums []int) []int {
	ret := make([]int, len(nums))
	left, right := 0, len(nums)-1
	for i := range ret {
		l, r := nums[left], nums[right]
		lv, rv := l*l, r*r
		if lv > rv {
			ret[len(ret)-1-i] = lv
			left++
		} else {
			ret[len(ret)-1-i] = rv
			right--
		}
	}
	return ret
}
