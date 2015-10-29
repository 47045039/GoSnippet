// 冒泡排序
package tsample

import (
	"fmt"
	"reflect"
)

type Comparer interface {
	Len() int
	Less(int, int) bool
	Swap(int, int)
}

type SI []int

type SS []string

func (si SI) Len() int {
	return len(si)
}

func (si SI) Less(i int, j int) bool {
	return si[j] < si[i]
}

func (si SI) Swap(i int, j int) {
	si[i], si[j] = si[j], si[i]
}

func (ss SS) Len() int {
	return len(ss)
}

func (ss SS) Less(i int, j int) bool {
	return ss[j] < ss[i]
}

func (ss SS) Swap(i int, j int) {
	ss[i], ss[j] = ss[j], ss[i]
}

func Sort(s Comparer) {
	for i := 0; i < s.Len()-1; i++ {
		for j := i + 1; j < s.Len(); j++ {
			if s.Less(i, j) {
				s.Swap(i, j)
			}
		}
	}
}

func TestSort() {
	si := SI{3, 9, 7, 1, 8, 2, 6, 4, 5}
	fmt.Printf("SI type = %s\n", reflect.TypeOf(si))

	fmt.Printf("ss addr 1: %d\n", si)
	fmt.Printf("ss addr 2: %d\n", &si)
	fmt.Printf("ss addr 3: %d\n", si[0])
	fmt.Printf("ss addr 3: %d\n", &(si[0]))

	fmt.Printf("before sort: %v\n", si)
	Sort(si)
	fmt.Printf("after sort: %v\n", si)

	ss := SS{"bcd", "abc", "123", "321", "aaa"}
	fmt.Printf("SS type = %s\n", reflect.TypeOf(ss))
	fmt.Printf("before sort: %v\n", ss)
	Sort(ss)
	fmt.Printf("after sort: %v\n", ss)
}
