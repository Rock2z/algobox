package _059SpiralMatrixII

func generateMatrix(n int) [][]int {
	ret := make([][]int, n)
	for i := range ret {
		ret[i] = make([]int, n)
	}
	c := 1
	x, y := 0, 0
	flag := false
	for {
		ret[x][y] = c
		c++
		if c == n*n+1 {
			return ret
		}
		if y+1 < n && ret[x][y+1] == 0 && !flag {
			y++
			continue
		}
		if x+1 < n && ret[x+1][y] == 0 && !flag {
			x++
			continue
		}
		if y-1 >= 0 && ret[x][y-1] == 0 && !flag {
			y--
			continue
		}
		if x-1 >= 0 && ret[x-1][y] == 0 {
			x--
			flag = true
			continue
		}
		c--
		flag = !flag
	}
}
