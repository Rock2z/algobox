package _015

import (
	"sort"
	"strconv"
)

func threeSum(nums []int) [][]int {
	type ans struct {
		a, b, c int
	}
	m := make(map[int]int, len(nums))
	ret := make([][]int, 0)
	dup := make(map[ans]bool)
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = i
	}

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			t := nums[i] + nums[j]
			if index, exist := m[0-t]; exist {
				if index == i || index == j {
					continue
				}
				a, b, c := i, j, index
				if nums[c] > nums[b] {
					b, c = c, b
				}
				if nums[b] > nums[a] {
					a, b = b, a
				}
				if nums[c] > nums[b] {
					b, c = c, b
				}
				tmp := ans{
					a: nums[a],
					b: nums[b],
					c: nums[c],
				}
				if !dup[tmp] {
					dup[tmp] = true
					ret = append(ret, []int{nums[a], nums[b], nums[c]})
				}
			}
		}
	}
	return ret
}

func threeSum2(nums []int) [][]int {
	sort.Ints(nums)
	tar := 0
	ret := make([][]int, 0)

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			cur := nums[l] + nums[i] + nums[r]
			nl := nums[l]
			nr := nums[r]
			if cur > tar {
				r--
			}
			if cur < tar {
				l++
			}
			if cur == tar {
				ret = append(ret, []int{nums[l], nums[i], nums[r]})
				for l < r && nums[l] == nl {
					l++
				}
				for l < r && nums[r] == nr {
					r--
				}
			}
		}
	}
	return ret
}

// not the best solution, but this is done by myself TAT
func threeSum3(nums []int) [][]int {
	sort.Ints(nums)
	ret := make([][]int, 0)
	retMap := make(map[string]bool)
	tar := 0
	for i := range nums {
		if i == 0 || i == len(nums)-1 {
			continue
		}
		l, r := i-1, i+1

		for {
			cur := nums[l] + nums[i] + nums[r]
			if cur < tar {
				r++
			} else if cur > tar {
				l--
			} else {
				ha := genHash(nums[l], nums[i], nums[r])
				if retMap[ha] == false {
					retMap[ha] = true
					ret = append(ret, []int{nums[l], nums[i], nums[r]})
				}
				l--
				r++
			}
			if r > len(nums)-1 || l < 0 {
				break
			}
		}
	}
	return ret
}

func genHash(a, b, c int) string {
	return strconv.Itoa(a) + strconv.Itoa(b) + strconv.Itoa(c)
}
