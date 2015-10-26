// 进阶课程，指针、内存分配、自定义类、类型转换
package tadv

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
	"sync"
)

func TestAdv1() {
	fmt.Println("@@@@@@@@@@@@ TestAdv 1")
	i := 5
	p := &i

	fmt.Printf("i=%d, *p=%d \n", i, *p) // i=5, *p=5

	i++
	fmt.Printf("i=%d, *p=%d \n", i, *p) // i=6, *p=6

	*p++                                // *p++ 等价于 (*p)++，等价于 i++
	fmt.Printf("i=%d, *p=%d \n", i, *p) // i=7, *p=7
}

type SyncedBuffer struct {
	lock sync.Mutex
	buff bytes.Buffer
}

func TestAdv2() {
	fmt.Println("@@@@@@@@@@@@ TestAdv 2")

	// new(T): 返回一个零值填充的T类型的内存空间的地址，即返回一个T类型的指针。该指针指向T类型的零值。
	// make(T, args): 只能创建slice、map和channel，并且返回一个有初始值的T类型。

	buffPointer := new(SyncedBuffer) // type *SyncedBuffer
	var buffer SyncedBuffer          // type SyncedBuffer

	fmt.Printf("type 1 = %s \n", reflect.TypeOf(buffPointer))      // type 1 = *tadv.SyncedBuffer
	fmt.Printf("type 2 = %s \n", reflect.TypeOf(buffer))           // type 2 = tadv.SyncedBuffer
	fmt.Printf("type 3 = %s \n", reflect.TypeOf(buffPointer.lock)) // type 3 = sync.Mutex
	fmt.Printf("type 4 = %s \n", reflect.TypeOf(buffer.lock))      // type 4 = sync.Mutex

	var p *[]int = new([]int)                       // *p = nil
	*p = make([]int, 100)                           // *p = array int 100
	fmt.Printf("type 5 = %s \n", reflect.TypeOf(p)) // *[]int

	var v []int = make([]int, 100)                    // v = array int 100
	v2 := make([]int, 100)                            // v2 = array int 100
	fmt.Printf("type v = %s \n", reflect.TypeOf(v))   // []int
	fmt.Printf("type v2 = %s \n", reflect.TypeOf(v2)) // []int
}

type NameAndAge struct {
	name string
	age  int
}

func Init(in *NameAndAge, name string, age int) {
	in.Init2(name, age)
}

func (in *NameAndAge) Init2(name string, age int) {
	in.name = name
	in.age = age
}

func (in NameAndAge) String() string {
	return "[name=" + in.name + " age=" + strconv.Itoa(in.age) + "]"
}

func TestAdv3() {
	fmt.Println("@@@@@@@@@@@@ TestAdv 3")

	nameAge := new(NameAndAge)

	fmt.Printf("before init: %s \n", nameAge.String())
	fmt.Printf("before init: %v \n", nameAge)
	Init(nameAge, "test_name", 10)
	fmt.Printf("after init: %s \n", nameAge.String())
	fmt.Printf("after init: %v \n", nameAge)

	nameAge = new(NameAndAge)
	fmt.Printf("before init 2: %s \n", nameAge.String())
	fmt.Printf("before init 2: %v \n", nameAge)
	nameAge.Init2("test_name", 10)
	fmt.Printf("after init 2: %s \n", nameAge.String())
	fmt.Printf("after init 2: %v \n", nameAge)
}

type NewMutex sync.Mutex
type NewMutex2 struct {
	sync.Mutex
	bytes.Buffer
}

func TestAdv4() {
	fmt.Println("@@@@@@@@@@@@ TestAdv 4")

	//	newMutex := new(NewMutex)
	// NewMutux等同于Mutex，但是它没有任何Mutex的方法。即它的方法是空的。所以，下边的代码无法运行

	//	newMutex.lock()
	//	newMutex.unlock()

	newMutex2 := new(NewMutex2)
	// NewMutex2已经从Mutex和Buffer继承了方法集合，包含了Lock/Unlock/Cap等方法，被绑定到其匿名字段Mutex和Buffer。

	newMutex2.Lock()
	newMutex2.Unlock()

	err := newMutex2.WriteByte(byte(1))
	if err == nil {
		fmt.Printf("NewMutex2 cap: %d \n", newMutex2.Cap())

		b, e := newMutex2.ReadByte()
		if e == nil {
			fmt.Printf("NewMutex2 read 1 byte: %d \n", b)
		} else {
			fmt.Printf("NewMutex2 read 1 byte failed: %s \n", e)
		}
	} else {
		fmt.Printf("NewMutex2 write 1 byte failed: %s \n", err)
	}
}

func TestAdv5() {
	fmt.Println("@@@@@@@@@@@@ TestAdv 5")

	str := "hello golang"

	bytesli := []byte(str)
	fmt.Printf("byte slice: %s \n", bytesli)

	//	intsli := []int(str)
	//	fmt.Printf("int slice: %s \n", intsli)

	str2 := string(bytesli)
	fmt.Printf("string 2: %s \n", str2)

	//	str3 := string(intsli)
	//	fmt.Printf("mystr3: %s \n", str3)
}

type itface interface{}

func mult(v itface) itface {
	switch v.(type) {
	case byte:
		return v.(byte) * v.(byte)
	case int:
		return v.(int) * v.(int)
	case int8:
		return v.(int8) * v.(int8)
	case int32:
		return v.(int32) * v.(int32)
	case int64:
		return v.(int64) * v.(int64)
	case string:
		return v.(string) + v.(string)
	default:
		return "unsupport type: " + reflect.TypeOf(v).String()
	}
}

func testMap(vs []itface, f func(itface) itface) []itface {
	mm := make([]itface, cap(vs))
	for idx, v := range vs {
		mm[idx] = f(v)
	}
	return mm
}

func TestAdv6() {
	fmt.Println("@@@@@@@@@@@@ TestAdv 6")
	fmt.Println("result: ", testMap([]itface{2, 3, 4, 5}, mult))
	fmt.Println("result 2: ", testMap([]itface{"string1", "string2", "string3"}, mult))
}
