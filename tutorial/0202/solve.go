package _202

func isHappy(n int) bool {
	m := make(map[int]struct{}, 1000)
	for i := 0; i < 1000; i++ {
		n = getSum(n)
		if n == 1 {
			return true
		}
		if _, exist := m[n]; exist {
			return n == 1
		}
		m[n] = struct{}{}
	}
	return false
}

func getSum(n int) int {
	ret := 0
	for {
		d := n % 10
		ret += d * d
		n /= 10
		if n == 0 {
			break
		}
	}
	return ret
}
