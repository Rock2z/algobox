package _541

func reverseStr(s string, k int) string {
	b := []byte(s)
	for i := range b {
		if (1+i)%(2*k) == 0 {
			reverse(b[i+1-k-k : i+1-k])
		}
	}
	leftLen := len(b) - (2*k)*(len(b)/(2*k))
	if leftLen > 0 {
		if leftLen < k {
			reverse(b[len(b)-leftLen : len(b)])
		} else {
			reverse(b[len(b)-leftLen : len(b)-leftLen+k])
		}
	}
	return string(b)
}

func reverseStr2(s string, k int) string {
	b := []byte(s)
	for i := 0; i < len(b); i += 2 * k {
		if i+k < len(b) {
			reverse(b[i : i+k])
		} else {
			//fix: if left less than k elements, need to reverse all of them
			reverse(b[i:])
		}
	}
	return string(b)
}

func reverse(b []byte) {
	for i := 0; i < len(b)/2; i++ {
		b[i], b[len(b)-1-i] = b[len(b)-1-i], b[i]
	}
}
