package _209

//1. brute force
//	- 用i遍历数组作为subslice的起点
//	- 用j遍历 num[i:] 作为subslice的终点
//	- 对于每个subslice求出sum 然后更新答案 -> improve 先用一个数组算好sums[i] = nums[0] + ... + nums[i] -> 计算subslice nums[i:j] 的时候 就可以直接用sum[j]-sum[i]来算
//	- space complexity: O(1) 或者 O(n)
//	- time complexity: O(n^3) 或者 O(n^2)
//2. 二分+记忆化
//	- 先用sums 数组存好 sums[i]=sums[i−1]+nums[i]
//	- 遍历nums, 对于nums[i], 问题就退化成了，在一个递增数组sums[i+1:] 中，找到第一个 sums[x] >= target-sums[i]，可以用二分查找
//	- space complexity: O(n)
//	- time complexity: O(nlogn)
//3. 双指针（滑动窗口法）
//	- 维护两个指针i和j
//	- 一直往后移动j，直到sum >= target
//	- 如果发现当前已经有 sum >= target 了，那说明至少已经找到了一个解，需要尝试找到更优解，所以把i往后移，直到sum<target为止
//	- 继续往后移动j，然后重复上面两步，这样就遍历了所有的可能性

func minSubArrayLen(target int, nums []int) int {
	ret := len(nums) + 1
	sum := 0
	l := 0
	for r := 1; r <= len(nums); r++ {
		sum += nums[r-1]
		for sum >= target {
			if r-l < ret {
				ret = r - l
			}
			sum -= nums[l]
			l++
		}
	}

	if ret == len(nums)+1 {
		return 0
	}
	return ret
}

//func minSubArrayLen2(target int, nums []int) int {
//	sums := make([]int, len(nums)+1)
//	sums[0] = 0
//	for i := 1; i < len(sums); i++ {
//		sums[i] = sums[i-1] + nums[i-1]
//	}
//
//	ret := 1<<31 -1
//	for i := range nums {
//		for j := i+1; j < len(nums)+1; j++ {
//			if sums[j] - sums[i] >= target {
//				if j-i < ret {
//					ret = j-i
//					break
//				}
//			}
//		}
//	}
//	if ret == 1<<31-1 {
//		return 0
//	}
//	return ret
//}
