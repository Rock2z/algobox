package _707DesignLinkedList

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"algobox/top_100_linked_questions/common"
)

func TestConstructor(t *testing.T) {
	linkedList := Constructor()
	linkedList.AddAtHead(1)
	linkedList.AddAtTail(3)
	linkedList.AddAtIndex(1, 2)
	fmt.Println(linkedList.Get(1))
	linkedList.DeleteAtIndex(0)
	fmt.Println(linkedList.Get(0))
	fmt.Println(common.MakeSliceFromNode(linkedList.head))
}

func TestConstructor2(t *testing.T) {
	fns := []string{
		"addAtHead", "addAtTail", "addAtTail",
		"get", "get", "addAtTail", "addAtIndex", "addAtHead", "addAtHead",
		"addAtTail", "addAtTail", "addAtTail", "addAtTail", "get", "addAtHead",
		"addAtHead", "addAtIndex", "addAtIndex", "addAtHead", "addAtTail",
		"deleteAtIndex", "addAtHead", "addAtHead", "addAtIndex", "addAtTail",
		"get", "addAtIndex", "addAtTail", "addAtHead", "addAtHead", "addAtIndex",
		"addAtTail", "addAtHead", "addAtHead", "get", "deleteAtIndex", "addAtTail",
		"addAtTail", "addAtHead", "addAtTail", "get", "deleteAtIndex", "addAtTail",
		"addAtHead", "addAtTail", "deleteAtIndex", "addAtTail", "deleteAtIndex",
		"addAtIndex", "deleteAtIndex", "addAtTail", "addAtHead", "addAtIndex",
		"addAtHead", "addAtHead", "get", "addAtHead", "get", "addAtHead", "deleteAtIndex",
		"get", "addAtHead", "addAtTail", "get", "addAtHead", "get", "addAtTail", "get",
		"addAtTail", "addAtHead", "addAtIndex", "addAtIndex", "addAtHead", "addAtHead",
		"deleteAtIndex", "get", "addAtHead", "addAtIndex", "addAtTail", "get", "addAtIndex",
		"get", "addAtIndex", "get", "addAtIndex", "addAtIndex", "addAtHead", "addAtHead",
		"addAtTail", "addAtIndex", "get", "addAtHead", "addAtTail", "addAtTail", "addAtHead",
		"get", "addAtTail", "addAtHead", "addAtTail", "get", "addAtIndex",
	}
	args := [][]int{
		[]int{84},
		[]int{2},
		[]int{39},
		[]int{3},
		[]int{1},
		[]int{42},
		[]int{1, 80},
		[]int{14},
		[]int{1},
		[]int{53},
		[]int{98},
		[]int{19},
		[]int{12},
		[]int{2},
		[]int{16},
		[]int{33},
		[]int{4, 17},
		[]int{6, 8},
		[]int{37},
		[]int{43},
		[]int{11},
		[]int{80},
		[]int{31},
		[]int{13, 23},
		[]int{17},
		[]int{4},
		[]int{10, 0},
		[]int{21},
		[]int{73},
		[]int{22},
		[]int{24, 37},
		[]int{14},
		[]int{97},
		[]int{8},
		[]int{6},
		[]int{17},
		[]int{50},
		[]int{28},
		[]int{76},
		[]int{79},
		[]int{18},
		[]int{30},
		[]int{5},
		[]int{9},
		[]int{83},
		[]int{3},
		[]int{40},
		[]int{26},
		[]int{20, 90},
		[]int{30},
		[]int{40},
		[]int{56},
		[]int{15, 23},
		[]int{51},
		[]int{21},
		[]int{26},
		[]int{83},
		[]int{30},
		[]int{12},
		[]int{8},
		[]int{4},
		[]int{20},
		[]int{45},
		[]int{10},
		[]int{56},
		[]int{18},
		[]int{33},
		[]int{2},
		[]int{70},
		[]int{57},
		[]int{31, 24},
		[]int{16, 92},
		[]int{40},
		[]int{23},
		[]int{26},
		[]int{1},
		[]int{92},
		[]int{3, 78},
		[]int{42},
		[]int{18},
		[]int{39, 9},
		[]int{13},
		[]int{33, 17},
		[]int{51},
		[]int{18, 95},
		[]int{18, 33},
		[]int{80},
		[]int{21},
		[]int{7},
		[]int{17, 46},
		[]int{33},
		[]int{60},
		[]int{26},
		[]int{4},
		[]int{9},
		[]int{45},
		[]int{38},
		[]int{95},
		[]int{78},
		[]int{54},
		[]int{42, 86},
	}

	l := Constructor()
	lv := reflect.ValueOf(&l)
	for i := range fns {
		paramList := make([]reflect.Value, 0, len(args[i]))
		for _, param := range args[i] {
			paramList = append(paramList, reflect.ValueOf(param))
		}
		if i == 83 {
			fmt.Println('w')
		}
		fnName := strings.ToUpper(fns[i][:1]) + fns[i][1:]
		vals := lv.MethodByName(fnName).Call(paramList)
		for _, v := range vals {
			println(v.Interface().(int))
		}
		fmt.Printf("%s, %v, result=%v\nnow=%v\n\n", fnName, args[i], vals, common.MakeSliceFromNode(l.head))
	}
}
