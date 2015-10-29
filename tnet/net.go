// 演示网络相关
package tnet

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
)

func testNet(protecol string, addr string) {
	c, e := net.Dial(protecol, addr)
	if e != nil {
		fmt.Fprintf(os.Stderr, "connect to %s:%s err: %v\n", protecol, addr, e)
		return
	}

	var r = bufio.NewReader(c)
	var w = bufio.NewWriter(os.Stdout)
	defer func() {
		c.Close()
		fmt.Fprintf(os.Stdout, "close connection of %s:%s\n", protecol, addr)
	}()

	defer func() {
		w.Flush()
		fmt.Fprintf(os.Stdout, "flush system output\n")
	}()

	var buff = make([]byte, 1024)
	for {
		n, _ := r.Read(buff)
		if n > 0 {
			w.Write(buff[0:n])
		} else {
			fmt.Fprintf(os.Stdout, "read for %s:%s finished\n", protecol, addr)
			return
		}
	}
}

func TestNet() {
	testNet("tcp", "127.0.0.1:80")
}

func testHttp(url string) {
	r, e := http.Get(url)
	if e != nil {
		fmt.Fprintf(os.Stderr, "err: %s, %v\n", url, e)
		return
	}

	defer r.Body.Close()

	b, e2 := ioutil.ReadAll(r.Body)
	if e2 != nil {
		fmt.Fprintf(os.Stderr, "err 2: %s, %v\n", url, e2)
		return
	}

	fmt.Fprintf(os.Stdout, "read: %s\n%s\n", url, string(b))
}

func TestHttp() {
	testHttp("http://www.baidu.com")
}
