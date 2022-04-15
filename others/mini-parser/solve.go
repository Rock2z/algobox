package mini_parser

import (
	"regexp"
	"strconv"
	"strings"
)

// This is the interface that allows for creating nested lists.
// You should not implement it, or speculate about its implementation
type NestedInteger struct {
	Val  *int
	Nest []*NestedInteger
}

// Return true if this NestedInteger holds a single integer, rather than a nested list.
func (n NestedInteger) IsInteger() bool { return n.Val != nil }

// Return the single integer that this NestedInteger holds, if it holds a single integer
// The result is undefined if this NestedInteger holds a nested list
// So before calling this method, you should have a check
func (n NestedInteger) GetInteger() int { return *n.Val }

// Set this NestedInteger to hold a single integer.
func (n *NestedInteger) SetInteger(value int) { n.Val = &value }

// Set this NestedInteger to hold a nested list and adds a nested integer to it.
func (n *NestedInteger) Add(elem NestedInteger) { n.Nest = append(n.Nest, &elem) }

// Return the nested list that this NestedInteger holds, if it holds a nested list
// The list length is zero if this NestedInteger holds a single integer
// You can access NestedInteger's List element directly if you want to modify it
func (n NestedInteger) GetList() []*NestedInteger { return n.Nest }

func deserialize(s string) *NestedInteger {
	return solve(s)
}

func solve(s string) *NestedInteger {
	if len(s) == 0 {
		return nil
	}
	return parseNest(s)
}

func parseNest(s string) *NestedInteger {
	ret := &NestedInteger{}
	if isNum(s) {
		ret.SetInteger(parseNum(s))
		return ret
	}
	if strings.HasPrefix(s, "[") {
		s = strings.TrimPrefix(s, "[")
		s = strings.TrimSuffix(s, "]")
		for _, integer := range parseNestContent(s) {
			ret.Add(*integer)
		}
		return ret
	}
	return nil
}

func parseNestContent(s string) []*NestedInteger {
	slice := split(s)
	ret := make([]*NestedInteger, 0, len(s))
	for _, e := range slice {
		ret = append(ret, parseNest(e))
	}
	return ret
}

func split(s string) []string {
	idx := 0
	ret := make([]string, 0, len(s))
	for idx <= len(s)-1 {
		if ok := isNum(getTok(s, idx)); ok {
			tmp := getTok(s, idx)
			for {
				idx++
				if idx > len(s)-1 {
					ret = append(ret, tmp)
					break
				}
				if ok = isNum(tmp + getTok(s, idx)); !ok {
					ret = append(ret, tmp)
					break
				}
				tmp += getTok(s, idx)
			}
		}
		if isLS(getTok(s, idx)) {
			tmp := getTok(s, idx)
			leftLS := 1
			for leftLS > 0 {
				idx++
				if idx > len(s)-1 {
					idx--
					break
				}
				if isLS(getTok(s, idx)) {
					leftLS++
				}
				if isRS(getTok(s, idx)) {
					leftLS--
				}
				tmp += getTok(s, idx)
			}
			ret = append(ret, tmp)
			idx++
		}
		if isCOMMA(getTok(s, idx)) {
			idx++
		}
	}
	return ret
}

func getTok(s string, idx int) string {
	if idx < 0 {
		idx = 0
	}
	if idx > len(s)-1 {
		idx = len(s) - 1
	}
	return s[idx : idx+1]
}

//NEST
//LS NUM COMMA NEST RS
//LS NEST COMMA NUM RS

func isLS(s string) bool {
	return s == "["
}

func isRS(s string) bool {
	return s == "]"
}

func isCOMMA(s string) bool {
	return s == ","
}

func isNum(s string) bool {
	res, err := regexp.Match("^-*[0-9]*$", []byte(s))
	if err != nil {
		return false
	}
	return res
}

func parseNum(s string) int {
	ret, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return ret
}
