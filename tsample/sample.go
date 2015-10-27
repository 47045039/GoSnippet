package tsample

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
)

var needLineNumber = flag.Bool("n", false, "print number of lines")

func cat(r *bufio.Reader, path string) {
	lineNumber := 1

	for { // 死循环
		line, err := r.ReadBytes('\n')
		if err == io.EOF {
			break
		}

		if *needLineNumber {
			fmt.Fprintf(os.Stdout, "%5d - %s", lineNumber, line)
			lineNumber++
		} else {
			fmt.Fprintf(os.Stdout, "%s", line)
		}
	}
}

func TestCat() {
	flag.Parse() // 解析命令行传递的参数列表

	for idx, arg := range os.Args {
		fmt.Fprintf(os.Stdout, "@@@ read arg: %d, %s \n", idx, arg)
	}
	fmt.Fprintf(os.Stdout, "@@@ needLineNumber: %s \n", strconv.FormatBool(*needLineNumber))

	var NARGS int = flag.NArg()
	if NARGS < 1 {
		fmt.Fprintf(os.Stderr, "@@@ Invalid arguments count: %d\n", NARGS)
		return
	}

	for i := 0; i < NARGS; i++ {
		var p string = flag.Arg(i)
		fmt.Fprintf(os.Stdout, "### cat file : %s\n", p)

		f, e := os.OpenFile(p, os.O_RDONLY, 0)
		if e != nil {
			fmt.Fprintf(os.Stderr, "### err: file=%s, err=%s\n", p, e)
			continue
		}

		cat(bufio.NewReader(f), p)
		defer func() { // 延时关闭文件
			f.Close() // 延时函数中的参数f和path都是在for循环中定义的
			fmt.Fprintf(os.Stdout, "### close file : %s\n", p)
		}()

		defer func(file *os.File, path string) { // 延时关闭文件，测试延时函数的参数
			file.Close()
			fmt.Fprintf(os.Stdout, "### close file 2 : %s\n", path)
		}(f, p)
	}
}
