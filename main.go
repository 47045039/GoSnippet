// CmdProject project main.go
package main

import (
	"fmt"

	"CmdProject/tadv"
	"CmdProject/tbase"
	"CmdProject/tfunc"
	"CmdProject/tsample"
)

func TestBase() {
	tbase.TestBase1()

	str1, str2 := tbase.TestBase2("E:\\android-5.1.1\\Makefile1")
	fmt.Printf("test2 result: %s, %s\n", str1, str2)

	tbase.TestBase3()

	tbase.TestBase4()

	println("'0xc' to byte: ", tbase.TestBase5('c'))

	tbase.TestBase6()

	tbase.TestBase7()

	tbase.TestBase8()
}

func TestFunc() {
	tfunc.TestFunc1(0)
	tfunc.TestFunc2()
	tfunc.TestFunc3()

	fmt.Println("result = ", tfunc.TestFunc4())
	fmt.Println("result = ", tfunc.TestFunc5())
	fmt.Println("result = ", tfunc.TestFunc6(2))

	tfunc.TestFunc7()
	tfunc.TestFunc8()
	tfunc.TestFunc9()
}

func TestAdv() {
	tadv.TestAdv1()
	tadv.TestAdv2()
	tadv.TestAdv3()
	tadv.TestAdv4()
	tadv.TestAdv5()
	tadv.TestAdv6()
}

func TestSample() {
	tsample.TestCat() // 命令行下执行：./CmdProject.exe -n doc.go main.go
}

func main() {
	fmt.Println("Hello World!")

	//	TestBase()

	//	TestFunc()

	//	TestAdv()

	TestSample()
}
