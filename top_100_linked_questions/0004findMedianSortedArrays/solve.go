package _004findMedianSortedArrays

import "sort"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	return solve(nums1, nums2)
}

func solve(nums1 []int, nums2 []int) float64 {
	c := append(nums1, nums2...)
	sort.Ints(c)
	if len(c)%2 == 0 {
		return (float64(c[len(c)/2]) + float64(c[(len(c)-1)/2])) / 2
	} else {
		return float64(c[len(c)/2])
	}
}
