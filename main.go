// CmdProject project main.go
package main

import (
	"fmt"
)


func main() {
	fmt.Println("Hello World!")

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
