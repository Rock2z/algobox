package _383

func canConstruct(ransomNote string, magazine string) bool {
	if len(magazine) < len(ransomNote) {
		return false
	}
	m := make(map[int32]int)
	uniqueChar := 0
	for _, c := range ransomNote {
		if _, exist := m[c]; !exist {
			uniqueChar++
		}
		m[c] += 1
	}
	for _, c := range magazine {
		m[c] -= 1
		if m[c] == 0 {
			uniqueChar--
		}
	}
	return uniqueChar == 0
}
