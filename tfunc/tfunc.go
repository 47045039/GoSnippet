// 函数相关课程，自定义类、全局和局部变量、匿名和延时函数、实现接口、panic和recovery等
package tfunc

import (
	"fmt"
	"reflect"
	"strconv"
)

var VAR int = 10

type stack struct {
	data  [10]int
	index int
}

func (s *stack) Push(v int) {
	if s.index <= 9 {
		s.data[s.index] = v
		s.index++
	} else {
		fmt.Println("stack is full: ", s.index)
	}
}

func (s *stack) Pop() (v int) {
	if s.index > 0 {
		s.index--
		v = s.data[s.index]
		s.data[s.index] = 0

		return
	} else {
		fmt.Println("stack is empty: ", s.index)
		return
	}
}

func (s *stack) Size() (v int) {
	return s.index
}

// 实现Stringer接口
func (s stack) String() string {
	var str string
	for i := 0; i < s.index; i++ {
		str += "[" + strconv.Itoa(i) + " : " + strconv.Itoa(s.data[i]) + "]"
	}
	return str
}

// 递归函数调用
func TestFunc1(i int) {
	fmt.Println("@@@@@@@@@@@@ TestFunc 1")
	if i == 2 {
		return
	}

	TestFunc1(i + 1)
	fmt.Printf(" %d ", i) // 2 1 0
	fmt.Println("")
}

func TestFunc2() {
	fmt.Println("@@@@@@@@@@@@ TestFunc 2")
	fmt.Println("global var = ", VAR)

	VAR := 5 // 局部变量
	fmt.Println("local var = ", VAR)
}

func TestFunc3() {
	fmt.Println("@@@@@@@@@@@@ TestFunc 3")

	// defer将多个动作加入一个LIFO（后进先出）的栈内，函数返回时，再执行多个动作。
	// 相当于java的try {} finally {}
	defer fmt.Println("delay do something 1111")
	defer fmt.Println("delay do something 2222")
	defer fmt.Println("delay do something 3333")

	fmt.Println("sync do something")

	// 函数执行结果如下：
	//	sync do something
	//	delay do something 3333
	//	delay do something 2222
	//	delay do something 1111
}

func TestFunc4() (result int) { // 返回参数result，初始化为0
	fmt.Println("@@@@@@@@@@@@ TestFunc 4")

	defer func() {
		result++
		fmt.Println("匿名函数调用 111: ", result) // result = 2
	}()
	defer func() {
		result++
		fmt.Println("匿名函数调用 222: ", result) // result = 1
	}()

	return result // result = 2
}

func TestFunc5() (result int) { // 返回参数result，初始化为0
	fmt.Println("@@@@@@@@@@@@ TestFunc 5")

	var arg int = 5
	defer func() {
		arg += 2
		fmt.Println("匿名函数调用 111: ", arg, &arg, result, &result) // arg = 8
	}()
	defer func() {
		arg += 3
		fmt.Println("匿名函数调用 222: ", arg, &arg, result, &result) // arg = 7
	}()

	arg++
	fmt.Println("PRINT: ", arg, &arg, result, &result) // arg=6,result=0

	// result = arg	// XXX：return是隐晦的执行了该行代码，然后执行了延时2个匿名函数，所以最终返回结果=6
	return arg // arg = 6	???? why
}

func TestFunc6(index int) (result int) {
	fmt.Println("@@@@@@@@@@@@ TestFunc 6")

	// map的key为int型，value为func型，而func有一个int型参数，返回一个int型结果
	// map ([int]  (func(int) int))
	mmm := map[int](func(int) int){
		1: func(arg int) int {
			return arg * 10
		},
		2: func(arg int) int {
			return arg * 100
		},
		3: func(arg int) int {
			return arg * 1000
		},
	}

	return mmm[index](index) // index = 2, 返回结果 = 2*100
}

func TestPanicAndRecover(test func()) (b bool) {
	//recover是一个内建的函数，可以让进入令人恐慌的流程中的goroutine恢复过来。Recover
	//仅在延迟函数中有效。在正常的执行过程中，调用recover会返回nil并且没有其他任何效果。
	//如果当前的goroutine陷入恐慌，调用recover可以捕获到panic的输入值，并且恢复正常的执行。
	defer func() {
		if recover() != nil { // recover() != nil，说明已经发生了panic，
			b = true
		}
	}()

	test()
	return // 返回值为b，无需指定，== return b
}

func TestFunc7() {
	fmt.Println("@@@@@@@@@@@@ TestFunc 7")

	//函数F调用panic，函数F的执行被中断，并且F中的延迟函数会正常执行，然后F返回到调
	//用它的地方。在调用的地方，F的行为就像调用了panic。这一过程继续向上，直到程
	//序崩溃时的所有goroutine返回。

	f := func() {
		fmt.Println("##### before panic #####")
		panic(1) // 程序崩溃
		fmt.Println("##### after panic #####")
	}

	fmt.Println("is panic = ", TestPanicAndRecover(f))

	fmt.Println("panic recovery")
}

func TestFunc8() {
	fmt.Println("@@@@@@@@@@@@ TestFunc 8")

	var ptr *stack = new(stack)
	ptr.Push(10)
	ptr.Push(20)
	ptr.Push(30)
	ptr.Push(40)

	fmt.Printf("after push, stack = %v \n", ptr)
	fmt.Printf("after push, stack size = %d \n", ptr.Size())

	fmt.Printf("pop stack = %d \n", ptr.Pop())
	fmt.Printf("pop stack = %d \n", ptr.Pop())

	fmt.Printf("after pop, stack = %v \n", ptr)
	fmt.Printf("after pop, stack size = %d \n", ptr.Size())
}

func sort(arr []int) {
	fmt.Printf("arr type: %s \n", reflect.TypeOf(arr))

	var ll int = len(arr)
	for i := 0; i < ll-1; i++ {
		for j := i + 1; j < ll; j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}

func TestFunc9() {
	fmt.Println("@@@@@@@@@@@@ TestFunc 9")
	var arr []int = []int{3, 5, 4, 7, 1, 2}
	fmt.Printf("arr type: %s \n", reflect.TypeOf(arr))

	fmt.Printf("before sort: %v \n", arr)
	sort(arr)
	fmt.Printf("after sort: %v \n", arr)
}

func TestFunc() {
	TestFunc1(0)
	TestFunc2()
	TestFunc3()

	fmt.Println("result = ", TestFunc4())
	fmt.Println("result = ", TestFunc5())
	fmt.Println("result = ", TestFunc6(2))

	TestFunc7()
	TestFunc8()
	TestFunc9()
}
