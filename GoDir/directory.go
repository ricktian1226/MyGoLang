// directory
package main

import (
	"fmt"
	"os"
)

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
	//1.mkdir
	Mk_dir("name")
}
