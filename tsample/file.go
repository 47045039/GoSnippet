// 读取文件，创建文件夹，命令行参数解析等
package tsample

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
)

func readFile(path string) {
	var w *bufio.Writer = bufio.NewWriter(os.Stdout)
	defer w.Flush()

	fmt.Fprintf(w, "read file start: %s\n", path)

	file, err := os.OpenFile(path, os.O_RDONLY, 0666) // 打开文件

	if err != nil {
		fmt.Printf("open file err: %v\n", err)
		return
	}

	defer file.Close() // 延时关闭文件

	var buff []byte = make([]byte, 1024)
	var r *bufio.Reader = bufio.NewReader(file)

	for {
		n, _ := r.Read(buff)
		if n <= 0 {
			fmt.Fprintf(w, "read file end: %s %d\n", path, n)
			return
		}

		w.Write(buff[0:n])
	}
}

func TestReadFile() {
	flag.Parse()

	NARGS := flag.NArg()
	for i := 0; i < NARGS; i++ {
		readFile(flag.Arg(i))
	}
}

func mkDir(dir string) {
	f, e := os.Stat(dir)
	if e != nil {
		fmt.Printf("path stat err: %s %v\n", dir, e)
		os.Mkdir(dir, os.FileMode(0755))
	} else if !f.IsDir() {
		fmt.Printf("path is not a dir: %s\n", dir)
		os.Remove(dir)
		os.Mkdir(dir, os.FileMode(0755))
	} else {
		fmt.Printf("path is existed: %s\n", dir)
	}
}

// ./CmdProject.exe -port 512 -n docc
var FLAG_INT = flag.Int("port", 80, "test int flag")
var FLAG_BOOL = flag.Bool("p", false, "test bool flag")

func TestMakeDir() {
	// os.Args是所有的输入参数。例如：./CmdProject.exe -port 512 -p docc就包含了5个参数
	for idx, arg := range os.Args {
		fmt.Fprintf(os.Stdout, "@@@ read arg: %d, %s \n", idx, arg)
	}

	// 该函数在./CmdProject.exe -h或者./CmdProject.exe --help时调用，输出命令行帮助信息
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [OPTIONS] [name ...]\n", os.Args[0])
		flag.PrintDefaults()
	}

	// 解析命令行参数
	flag.Parse()

	fmt.Fprintf(os.Stdout, "test int flag: %d\n", *FLAG_INT)                          // 512
	fmt.Fprintf(os.Stdout, "test boolean flag: %s\n", strconv.FormatBool(*FLAG_BOOL)) // true

	// 除开了cmd和flag之后的部分，./CmdProject.exe -port 512 -p docc中的docc
	NARGS := flag.NArg()
	for i := 0; i < NARGS; i++ {
		mkDir(flag.Arg(i))
	}
}
