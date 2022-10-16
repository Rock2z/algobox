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

func generateMatrix02(n int) [][]int {
	ret := make([][]int, n)
	for i := range ret {
		ret[i] = make([]int, n)
	}

	i, j, cur := 0, 0, 1
	for {
		ret[i][j] = cur
		cur++
		if j < n-1 && ret[i][j+1] == 0 {
			j++
			continue
		}
		if i < n-1 && ret[i+1][j] == 0 {
			i++
			continue
		}
		if j > 0 && ret[i][j-1] == 0 {
			j--
			continue
		}
		for i > 0 && ret[i-1][j] == 0 {
			i--
			ret[i][j] = cur
			cur++
		}
		cur--
		if cur >= n*n {
			break
		}
	}
	return ret
}
