// 测试chan用法，并发，sleep等
package tconc

import (
	"fmt"
	"runtime"
	"strconv"
	"time"
)

func sleep(key string, sec int) {
	fmt.Printf("before sleep %d seconds: %s.\n", sec, key)
	time.Sleep(time.Duration(sec) * time.Second) // sleep i 秒，time.Sleep()
	fmt.Printf("after sleep %d seconds: %s.\n", sec, key)
}

func TestChan1() {
	fmt.Println("@@@@@@@@@@@@ TestChan 1")
	go sleep("goroutine 1", 4) // 在goroutine中运行，sleep 4
	go sleep("goroutine 2", 2) // 在goroutine中运行，sleep 2

	// 同步运行，sleep 5，必须比2个goroutine中的sleep时间更长。
	// 否则，主goroutine结束，两个副的goroutine将立即停止运行
	sleep("main goroutine", 5)

	// 打印结果如下。可以看出main goroutine最后退出
	// before sleep 5 seconds: main goroutine.
	// before sleep 4 seconds: goroutine 1.
	// before sleep 2 seconds: goroutine 2.
	// after sleep 2 seconds: goroutine 2.
	// after sleep 4 seconds: goroutine 1.
	// after sleep 5 seconds: main goroutine.
}

func read(ch chan int) {
	fmt.Printf("before read from channel: %v\n", ch)
	var val int = <-ch
	fmt.Printf("after read from channel: %d %v\n", val, ch)
}

func write(ch chan int, val int) {
	fmt.Printf("before write to channel: %d %v\n", val, ch)
	ch <- val
	fmt.Printf("after write to channel: %d %v\n", val, ch)
}

func TestChan2() {
	fmt.Println("@@@@@@@@@@@@ TestChan 2")
	var ch1 chan int = make(chan int, 0) // 定义一个chan，这个chan可以读写int
	//	var ch2 chan<- int // 定义一个chan，这个chan只可以读int
	//	var ch3 <-chan int // 定义一个chan，这个chan只可以写int

	go write(ch1, 2)
	go write(ch1, 3)
	go read(ch1)
	go read(ch1)

	sleep("TestChan2", 2)

	// 顺序：goroutine 1 write -> goroutine 1 wait -> goroutine 2 read -> goroutine 1 continue
	// 打印结果如下。
	//	before sleep 2 seconds: TestChan2.
	//	before write to channel: 2 0xc0820101e0
	//	before write to channel: 3 0xc0820101e0

	//	before read from channel: 0xc0820101e0
	//	after read from channel: 2 0xc0820101e0
	//	after write to channel: 2 0xc0820101e0

	//	before read from channel: 0xc0820101e0
	//	after read from channel: 3 0xc0820101e0
	//	after write to channel: 3 0xc0820101e0

	//	after sleep 2 seconds: TestChan2.
}

func TestChan3() {
	fmt.Println("@@@@@@@@@@@@ TestChan 3")

	fmt.Printf("cpu num: %d\n", runtime.NumCPU()) // 8核cpu

	// 虽然goroutine是并发执行的，但是它们并不是并行运行的。如果不告诉Go额外的东西，同
	// 一时刻只会有一个goroutine执行。利用runtime.GOMAXPROCS(n)可以设置goroutine
	// 并行执行的数量。GOMAXPROCS 设置了同时运行的CPU 的最大数量，并返回之前的设置。
	val := runtime.GOMAXPROCS(runtime.NumCPU() * 4)
	fmt.Printf("last goroutine num: %d\n", val) // 8个

	fmt.Printf("goroutine num: %d\n", runtime.NumGoroutine()) // 4个goroutine同时运行

	var ch1 chan int = make(chan int, 0)
	var ch2 chan int = make(chan int, 0)
	var ch3 chan int = make(chan int, 0)

	go write(ch1, 22)
	go write(ch2, 33)
	go write(ch3, 44)
	go read(ch1)
	go read(ch2)
	go read(ch3)

	fmt.Printf("goroutine num: %d\n", runtime.NumGoroutine()) // 10个goroutine同时运行
	sleep("TestChan3", 3)
}

func shower(work chan int, exit chan bool) {
	defer func() {
		fmt.Println("exit shower 1.") // return之前执行
	}()

	for { // 死循环
		select {
		case i := <-work:
			fmt.Printf("worker received: %d\n", i)
		case b := <-exit:
			fmt.Printf("exiter received: %s\n", strconv.FormatBool(b))
			return // 退出函数，goroutine结束运行
		}
	}

	fmt.Println("exit shower 2.") // 不会执行
}

func writer(work chan int, exit chan bool) {
	for i := 0; i < 3; i++ {
		fmt.Printf("worker write: %d\n", i)
		work <- i
	}
	exit <- false

	fmt.Println("exit writer.")
}

func TestChan4() {
	fmt.Println("@@@@@@@@@@@@ TestChan 4")
	var work chan int = make(chan int, 5)   // 缓冲区大小为5，最多可以连续写5个数据
	var exit chan bool = make(chan bool, 1) // 缓冲区大小为1

	go writer(work, exit)
	go shower(work, exit)

	sleep("TestChan4", 4)
}
