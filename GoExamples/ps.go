/*
编写一个程序，列出所有正在运行的进程，并打印每个进程执行的子进程个数。
输出应当类似：
Pid 0 has 2 children: [1 2]
Pid 490 has 2 children: [1199 26524]
Pid 1824 has 1 child: [7293]
*/

package main

import (
	"fmt"
	"os/exec"
	"sort"
	"strconv"
	"strings"
)

func main() {
	ps := exec.Command("ps", "-e", "-opid,ppid,comm")
	output, _ := ps.Output()
	child := make(map[int][]int) //child是一个以key是int，value是元素为int的切片
	for i, s := range strings.Split(string(output), "\n") {
		if i == 0 { //跳过第一行
			continue
		}

		if len(s) == 0 { //跳过最后一行
			continue
		}

		f := strings.Fields(s)
		fpp, _ := strconv.Atoi(f[1])        //父进程
		fp, _ := strconv.Atoi(f[0])         //子进程
		child[fpp] = append(child[fpp], fp) //将子进程放在父进程的列表中
	}

	schild := make([]int, len(child))
	i := 0
	for k, _ := range child {
		schild[i] = k
		i++
	}
	sort.Ints(schild)

	for _, ppid := range schild {
		fmt.Printf("Pid %d has %d child", ppid, len(child[ppid]))
		if len(child[ppid]) == 1 {
			fmt.Printf(": %v\n", child[ppid])
			continue
		}
		fmt.Printf("ren : %v\n", child[ppid])
	}
}
