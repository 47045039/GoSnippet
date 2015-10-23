// CmdProject project main.go
package main

import "fmt"

func TestBase() {
	test1()

	str1, str2 := test2("E:\\android-5.1.1\\Makefile1")
	fmt.Printf("test2 result: %s, %s\n", str1, str2)

	test3()

	test4()

	println("'0xc' to byte: ", test5('c'))

	test6()

	test7()

	test8()
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

func main() {
	fmt.Println("Hello World!")

	//	TestBase()

	TestFunc()

}
