package _002

func commonChars(words []string) []string {
	if len(words) == 0 {
		return []string{}
	}
	m := make(map[int32]int, 26)
	for _, c := range words[0] {
		m[c] = m[c] + 1
	}
	for _, word := range words {
		tm := make(map[int32]int, 26)
		for _, c := range word {
			tm[c] = tm[c] + 1
		}
		m = filterMap(m, tm)
	}
	ret := make([]string, 0, len(m))
	for k, v := range m {
		for i := 0; i < v; i++ {
			ret = append(ret, string(k))
		}
	}
	return ret
}

func filterMap(a, b map[int32]int) map[int32]int {
	if len(b) < len(a) {
		a, b = b, a
	}
	for k, v := range a {
		if _, exist := b[k]; !exist {
			delete(a, k)
			continue
		}
		if b[k] < v {
			a[k] = b[k]
		}
	}
	return a
}
