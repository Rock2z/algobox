package _004findMedianSortedArrays

import "sort"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	return solve2(nums1, nums2)
}

// merge then find
func solve(nums1 []int, nums2 []int) float64 {
	c := append(nums1, nums2...)
	sort.Ints(c)
	return (float64(c[(len(c)-1)/2]) + float64(c[(len(c))/2])) / 2
}

// two pointers
func solve2(nums1 []int, nums2 []int) float64 {
	i, j := 0, 0
	tLen := len(nums1) + len(nums2)
	ret := float64(0)
	pre := 0

	for {
		curIdx := i + j - 1
		if curIdx >= (tLen-1)/2 {
			ret += float64(pre)
			if curIdx == tLen/2 {
				if tLen%2 == 0 {
					return ret / 2
				} else {
					return ret
				}
			}
		}
		if i == len(nums1) {
			pre = nums2[j]
			j++
			continue
		}
		if j == len(nums2) {
			pre = nums1[i]
			i++
			continue
		}
		if nums1[i] <= nums2[j] {
			pre = nums1[i]
			i++
		} else {
			pre = nums2[j]
			j++
		}
	}
}

// binary search
//func solve3(nums1 []int, nums2 []int) float64 {
//
//}
