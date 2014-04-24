// reader
package main

import (
	"bufio"
	"fmt"
	"os"
)

//没有缓冲的全读写文件
func io_without_buf() {
	buf := make([]byte, 1024)
	f, _ := os.Open("/etc/profile")
	defer f.Close()
	for {
		n, _ := f.Read(buf)
		if n == 0 {
			break
		} else {
			//buf[n-1] = 0  //这样写是没用的，和c的字符串区分开
			os.Stdout.Write(buf[0:n])
		}
	}
}

//带缓冲的全读写文件
func io_with_buf() {
	buf := make([]byte, 1024)
	f, _ := os.Open("/etc/profile")
	defer f.Close()
	r := bufio.NewReader(f)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	for {
		n, _ := r.Read(buf)
		if n == 0 {
			break
		} else {
			w.Write(buf[0:n])
		}
	}
}

//按行读取文件 ReadString/ReadLine

//判断目录是否存在，不存在则创建之
func Mk_dir(name string) {
	if _, err := os.Stat(name); err != nil {
		fmt.Printf("dir %s does't exist, mk it.\n", name)
		os.Mkdir(name, 0666)
	} else {
		fmt.Printf("dir %s exists.\n", name)
	}
}

func main() {
	//1.无缓存读写
	//io_without_buf()

	//2.带缓存读写
	//io_with_buf()

	//3.目录是否存在，不存在则创建之
	//Mk_dir("name")
}
