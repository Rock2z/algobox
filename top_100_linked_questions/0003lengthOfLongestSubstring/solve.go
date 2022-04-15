package _003lengthOfLongestSubstring

func lengthOfLongestSubstring(s string) int {
	return solve2(s)
}

func solve(s string) int {
	m := make(map[int32]int)
	maxLen := 0
	for i, e := range s {
		if idx, exist := m[e]; exist {
			for k, v := range m {
				if v <= idx {
					delete(m, k)
				}
			}
		}
		m[e] = i
		if len(m) > maxLen {
			maxLen = len(m)
		}
	}
	return maxLen
}

func solve2(s string) int {
	m := make(map[int32]int)
	maxLen := 0
	from, to := 0, 0
	for i, e := range s {
		if idx, exist := m[e]; exist {
			if idx+1 > from {
				from = idx + 1
			}
		}
		m[e] = i
		to += 1
		if to-from > maxLen {
			maxLen = to - from
		}
	}
	return maxLen
}
