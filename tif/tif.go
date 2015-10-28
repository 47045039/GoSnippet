// 接口相关课程：接口定义，接口实现，接口做函数参数，接口判断和接口转换等
package tif

import "fmt"

type I interface { // 定义一个interface
	Get() int
	Set(int)
}

type S struct { // 定义一个类S
	i int
}

type S2 struct { // 定义一个类S2
	i2 int
}

func (s *S) Get() int { // *S实现I接口的Get() int方法，所以在调用testI(I)时参数是一个*S
	return s.i
}

func (s *S) Set(i int) { // *S实现I接口的Set(int)方法，所以在调用testI(I)时参数是一个*S
	s.i = i
}

func (s2 *S2) Get() int { // *S2实现I接口的Get() int方法
	return s2.i2
}

func (s2 *S2) Set(i int) { // *S2实现I接口的Set(int)方法
	s2.i2 = i
}

func testIf(i I) {
	fmt.Printf("@@@ I.Get() = %d\n", i.Get())

	i.Set(2)
	fmt.Printf("@@@ I.Get() = %d\n", i.Get())
}

func TestIf1() {
	fmt.Println("@@@@@@@@@@@@ TestIf 1")
	s := new(S) // new创建一个S类型的指针变量s
	testIf(s)   // *S实现了I接口，所以参数必须是*S

	var s2 S    // 创建一个S类型的变量s2
	testIf(&s2) // *S实现了I接口，所以参数必须是*S
}

func testIf2(i I) {
	switch i.(type) {
	//	case S:
	//		fmt.Printf("I is a S instance \n")
	//		//		fallthrough		// 自动向下匹配
	//	case S2:
	//		fmt.Printf("I is a S2 instance \n")
	case *S:
		fmt.Printf("I is a *S instance \n")
	case *S2:
		fmt.Printf("I is a *S2 instance \n")
	default:
		fmt.Printf("I is a unknown type \n")
	}
}

func TestIf2() {
	fmt.Println("@@@@@@@@@@@@ TestIf 2")

	s := new(S)
	testIf2(s)

	var s2 S2
	testIf2(&s2)
}

func testIf3(any interface{}) int {
	return any.(I).Get() // any.(I) 表示将any强制转换为I接口类型，如果转换失败将报错
}

func testIf3_2(any interface{}, v int) {
	any.(I).Set(v) // any.(I) 表示将any强制转换为I接口类型，如果转换失败将报错
}

func TestIf3() {
	fmt.Println("@@@@@@@@@@@@ TestIf 3")

	s := new(S)
	fmt.Println("s.Get() 1 =", testIf3(s))
	//	fmt.Println("s.Get() 2 =", s.(I).Get()) // 报错，x.(T)语法只能使用在x是interface时

	var s2 S2
	testIf3_2(&s2, 5) // 设置s2
	fmt.Println("s2.Get() 1 =", testIf3(&s2))
	//	fmt.Println("s2.Get() 2 =", (&s2).(I).Get()) // 报错，x.(T)语法只能使用在x是interface

}

type Emitter interface {
	Emit()
}

type Foo int

// 扩展自定义类型成功
// 参数是*Foo类型。在调用Emit()方法时，传递指针，没有发生拷贝。改变*foo，将改变原值
func (foo *Foo) Emit() {
	fmt.Printf("foo = %v, addr = %d \n", *foo, foo)

	*foo = 2
}

// 参数是Foo类型。在调用Emit2()方法时，实际上做了参数的拷贝。改变foo2，未改变原值
func (foo2 Foo) Emit2() {
	fmt.Printf("foo2 = %v, addr = %d \n", foo2, &foo2)

	foo2 = 2
}

//// 扩展内建类型错误
//func (foo float32) Emit() {
//	fmt.Printf("foo = %f\n", foo)
//}

//// 扩展内建类型错误
//func (foo os.File) Emit() {
//	fmt.Printf("foo = %v\n", foo)
//}

func TestIf4() {
	fmt.Println("@@@@@@@@@@@@ TestIf 4")

	foo := new(Foo)
	*foo = 5
	fmt.Printf("foo = %d, addr = %d \n", *foo, foo) // foo = 5, addr = 826814792992
	foo.Emit()                                      // foo = 5, addr = 826814792992
	fmt.Printf("foo = %d, addr = %d \n", *foo, foo) // foo = 2, addr = 826814792992

	var foo2 Foo
	foo2 = 5
	fmt.Printf("foo2 = %d, addr = %d \n", foo2, &foo2) // foo2 = 5, addr = 826814793056
	foo2.Emit2()                                       // foo2 = 5, addr = 826814793088
	fmt.Printf("foo2 = %d, addr = %d \n", foo2, &foo2) // foo2 = 5, addr = 826814793056

	// 从上边的打印信息可以看出：使用指针做函数的参数，函数中操作形参，将真正改变实参
	// 而使用普通类型做函数的参数，函数中操作形参，实际操作的是实参的拷贝，并不会影响到实参
}
