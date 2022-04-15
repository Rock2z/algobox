package _001twoSum

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

//https://leetcode.com/problems/two-sum/
func Test_twoSum(t *testing.T) {
	t.Run("example1", func(t *testing.T) {
		res := twoSum([]int{2, 7, 11, 15}, 9)
		assert.EqualValues(t, []int{0, 1}, res)
	})
	t.Run("example2", func(t *testing.T) {
		res := twoSum([]int{3, 2, 4}, 6)
		assert.EqualValues(t, []int{1, 2}, res)
	})
	t.Run("example3", func(t *testing.T) {
		res := twoSum([]int{3, 3}, 6)
		assert.EqualValues(t, []int{0, 1}, res)
	})
}

//Benchmark_solve-16    	      27	  47900711 ns/op
func Benchmark_solve(b *testing.B) {
	a := make([]int, 10002)
	for i := 0; i < 10000; i++ {
		a[i] = 7 + rand.Intn(1<<30)
	}
	a = append(a, []int{2, 4}...)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = solve(a, 6)
	}
	b.StopTimer()
}

//Benchmark_solve2-16    	    2949	    380416 ns/op
func Benchmark_solve2(b *testing.B) {
	a := make([]int, 10002)
	for i := 0; i < 10000; i++ {
		a[i] = 7 + rand.Intn(1<<30)
	}
	a = append(a, []int{2, 4}...)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = solve2(a, 6)
	}
	b.StopTimer()
}
