package _150

import (
	"strconv"
)

func evalRPN(tokens []string) int {
	st := make([]int, 0, len(tokens))
	for _, s := range tokens {
		switch {
		case s == "+":
			a, b := st[len(st)-1], st[len(st)-2]
			st = st[:len(st)-2]
			st = append(st, a+b)
		case s == "-":
			a, b := st[len(st)-1], st[len(st)-2]
			st = st[:len(st)-2]
			st = append(st, a-b)
		case s == "*":
			a, b := st[len(st)-1], st[len(st)-2]
			st = st[:len(st)-2]
			st = append(st, a*b)
		case s == "/":
			a, b := st[len(st)-1], st[len(st)-2]
			st = st[:len(st)-2]
			st = append(st, a/b)
		default:
			num, _ := strconv.Atoi(s)
			st = append(st, num)
		}
	}
	return st[0]
}
