// client
package main

import (
	"fmt"
	"net"
	"time"
)

const RECV_BUF_LEN = 1024

func connect(hostAndPort string) {
	conn, err := net.Dial("tcp", hostAndPort)
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	buf := make([]byte, RECV_BUF_LEN+1)

	for i := 0; i < 5; i++ {
		//准备要发送的字符串
		msg := fmt.Sprintf("Hello ricktian, %03d", i)
		_, err := conn.Write([]byte(msg))
		if err != nil {
			println("Write Buffer Error : ", err.Error())
			break
		}
		fmt.Println(msg)
	}

	//从服务器端接受字符串
	n, err := conn.Read(buf)
	if err != nil {
		println("Read Buffer Error : ", err.Error())
	}

	println(string(buf[0:n]))

	//等一秒钟
	time.Sleep(time.Second)
}
