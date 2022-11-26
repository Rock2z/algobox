package _018

import (
	"sort"
	"strconv"
)

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	ret := make([][]int, 0)

	if len(nums) < 4 {
		return ret
	}

	retMap := make(map[string]bool)
	for i := range nums {
		tari := target - nums[i]
		for j := range nums {
			if j <= i+1 || j == len(nums)-1 {
				continue
			}
			tarij := tari - nums[j]
			l, r := j-1, j+1
			for {
				cur := nums[l] + nums[r]
				if cur < tarij {
					r++
				} else if cur > tarij {
					l--
				} else {
					has := genHash(nums[i], nums[l], nums[j], nums[r])
					if !retMap[has] {
						retMap[has] = true
						ret = append(ret, []int{nums[i], nums[l], nums[j], nums[r]})
					}
					l--
					r++
				}
				if l <= i || r > len(nums)-1 {
					break
				}
			}
		}
	}
	return ret
}

func genHash(nums ...int) string {
	var s string
	for _, n := range nums {
		s = s + strconv.Itoa(n) + "_"
	}
	return s
}
