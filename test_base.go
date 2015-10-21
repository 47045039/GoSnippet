// CmdProject project main.go
package main

import (
	"fmt"
	"os"
	"unicode/utf8"
)

func test1() {
	println("@@@@@@@@@@@@@@ test 1")
	println("@@@@@@@@@@@@@@ test 1")
	println("@@@@@@@@@@@@@@ test 1")
	var a int32 = 1
	a++
	b := a + 2

	fmt.Printf("a=%d, b=%d\n", a, b)

	const (
		x = iota
		y = iota
	)

	fmt.Printf("x=%d, y=%d\n", x, y)

	const (
		c string = "abc" + "edf"
		//d byte   = c[0]
	)

	var d byte = c[0]
	var e []byte = []byte(c)
	fmt.Printf("c=%s, c[0]=%d, d=%d, e=%s \n", c, c[0], d, e)
}

func test2(path string) (string, string) {
	println("@@@@@@@@@@@@@@ test 2")
	println("@@@@@@@@@@@@@@ test 2")
	println("@@@@@@@@@@@@@@ test 2")

	file, err := os.OpenFile(path, os.O_RDONLY, 0)
	if err != nil {
		fmt.Println(err)
		return "error", err.Error()
	} else {
		fmt.Println(file)
		return "result", "result2"
	}
}

func test3() {
	println("@@@@@@@@@@@@@@ test 3")
	println("@@@@@@@@@@@@@@ test 3")
	println("@@@@@@@@@@@@@@ test 3")
	var sum int
	for i := 0; i < 10; i++ {
		sum += i
	}

	fmt.Printf("sum = %d\n", sum)

LOOP1:
	for i := 1; ; /*i < 3*/ i++ {
		for j := 5; j > 0; j-- {
			if i > 1 {
				break LOOP1 // 跳出LOOP1循环
			}

			println("loop 1: ", i, j)
		}
	}

	i := 0
	for i < 2 { // 相当于for(;;;) {} 和 while () {}
		i++
		println("loop 2: ", i)
	}

	i = 0
	for { // 相当于死循环
		i++
		if i > 2 {
			break
		}

		println("loop 3: ", i)
	}
}

func test4() {
	println("@@@@@@@@@@@@@@ test 4")
	println("@@@@@@@@@@@@@@ test 4")
	println("@@@@@@@@@@@@@@ test 4")

	var list []string = make([]string, 5, 10)
	list = []string{"abc", "def", "hij"}
	for k, v := range list {
		println("list [", k, "] = "+v)
	}

	for pos, char := range "abcd" {
		println("abcd[ ", pos, "]  ==  ", char)
	}
}

func test5(c byte) byte {
	println("@@@@@@@@@@@@@@ test 5")
	println("@@@@@@@@@@@@@@ test 5")
	println("@@@@@@@@@@@@@@ test 5")
	i := 2
	switch i {
	case 1:
		println(" i  =  1 ") // 省略了break，i=1时调用
	case 2:
		println(" i  =  2 ")
	}

	switch i {
	case 1:
		fallthrough // i=1时自动向下匹配
	case 2:
		println(" i  = ", i) // XXXXX: 打印 i = 2 ？？？？
	}

	switch {
	case '0' <= c && c <= '9':
		return c - '0'
	case 'a' <= c && c <= 'f':
		return c - 'a' + 10
	case 'A' <= c && c <= 'F':
		return c - 'A' + 10
	default:
		return 0
	}
}

func test6() {
	println("@@@@@@@@@@@@@@ test 6")
	println("@@@@@@@@@@@@@@ test 6")
	println("@@@@@@@@@@@@@@ test 6")

	var arr [6]int // 定义数组
	arr[0] = 111
	arr[1] = 222
	arr[2] = 333

	arr2 := []int{111, 222} // 定义数组，不指定长度，指定元素
	arr3 := [2]int{111}     // 定义数组，指定长度，指定部分元素

	for i := 0; i < cap(arr); /*len(arr)*/ i++ {
		//		fmt.Printf("array element %d value is %d \n", i, arr[i])
		println("array element", i, "value is", arr[i])
	}

	// 数组的len()和cap()是一样的
	println("array len:", len(arr), "   array cap:", cap(arr))
	println("array2 len:", len(arr2), "   array2 cap:", cap(arr2))
	println("array3 len:", len(arr3), "   array3 cap:", cap(arr3))

	slic := make([]int, 10, 20) // 定义一个len=10，cap=20的slice
	slic2 := arr[2:5]           // 由array生成一个slice，元素为arr[2]--arr[4]，len=3，cap=4
	println("slice len:", len(slic), "   slice cap:", cap(slic))
	println("slice2 len:", len(slic2), "   slice2 cap:", cap(slic2))
	for i := 0; i < len(slic2); i++ {
		println("slice2 element", i, "value is", slic2[i]) // 333,0,0
	}

	slic3 := append(slic2, 2, 3)
	for i := 0; i < len(slic3); i++ {
		println("slice3 element", i, "value is", slic3[i]) // 333,0,0,2,3
	}

	slic4 := append(slic3, []int{11, 22} /*slic*/ ...)
	for i := 0; i < len(slic4); i++ {
		println("slice4 element", i, "value is", slic4[i]) // 333,0,0,2,3,11,22
	}

	var a []int = []int{0, 1, 2, 3, 4, 5, 6}
	var s = make([]int, 5, 10)
	n1 := copy(s, a[0:]) // n1=6, s= []int {0,1,2,3,4}
	println("n1 = ", n1)
	for i := 0; i < len(s); i++ {
		println("s element", i, "value is", s[i]) // 0,1,2,3,4
	}

	n2 := copy(s, s[2:]) // n2=3, s= []int {2,3,4,3,4}
	println("n2 = ", n2)
	for i := 0; i < len(s); i++ {
		println("s element", i, "value is", s[i]) // 2,3,4,3,4
	}
}

func test7() {
	println("@@@@@@@@@@@@@@ test 7")
	println("@@@@@@@@@@@@@@ test 7")
	println("@@@@@@@@@@@@@@ test 7")

	monthdays := make(map[string]int, 5)      // 定义一个len=5的map
	println("monthdays len:", len(monthdays)) // 注意，map没有cap
	monthdays = map[string]int{
		"Jan": 31, "Feb": 28, "Mar": 31, "Apr": 30, // 注意，最后的，是必须的
	}

	monthdays["May"] = 31 // 添加元素
	monthdays["Jun"] = 30 // 添加元素
	monthdays["Jul"] = 31 // 添加元素
	monthdays["Aug"] = 31 // 添加元素
	monthdays["Sep"] = 30 // 添加元素
	monthdays["Oct"] = 31 // 添加元素
	monthdays["Nov"] = 30 // 添加元素
	monthdays["Dec"] = 31 // 添加元素

	var year int = 0
	for _, v := range monthdays { // 变量_声明但未使用，不会报错，换个名字将报错
		year += v
	}
	println("Numbers of days in a year:", year)

	//	_ := monthdays	// 报错：no new variables on left side of :=

	value, present := monthdays["May"] // 判断monthdays中是否存在以“May”为key的元素
	println("monthdays[\"May\"] =", value, present)

	delete(monthdays, "May")          // 从monthdays中移除以“May”为key的元素
	value, present = monthdays["May"] // 判断monthdays中是否存在以“May”为key的元素
	println("monthdays[\"May\"] =", value, present)
}

func test8() {
	var str string = "abcdefghijk"
	var result map[int]int = make(map[int]int, 10)

	for c := range str {
		_, present := result[c]
		if present {
			result[c] = result[c] + 1
		} else {
			result[c] = 1
		}
	}

	for k, v := range result {
		println("result[", k, "] =", v)
	}

	var arr []byte = []byte(str)
	println("result:", utf8.RuneCount(arr))

	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
//		println("swap", i, j)
		arr[i], arr[j] = arr[j], arr[i]
	}
	println("result:", string(arr))
}

