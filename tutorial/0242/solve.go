package _242

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m := make(map[int32]int, len(s))
	for _, c := range s {
		m[c] = m[c] + 1
	}
	for _, c := range t {
		if m[c] <= 0 {
			return false
		}
		m[c] -= 1
	}
	return false
}
