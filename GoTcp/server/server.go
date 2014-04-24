// server
package main

import (
	"fmt"
	"net"
)

var (
	maxread  = 1100
	msgStop  = []byte("cmdStop")
	msgStart = []byte("cmdStart")
)

func initServer(hostAndPort string) *net.TCPListener {
	serverAddr, err := net.ResolveTCPAddr("tcp", hostAndPort)
	checkError(err, "Resolving address : port failed :'"+hostAndPort+"'")
	listener, err := net.ListenTCP("tcp", serverAddr)
	checkError(err, "ListenTCP : ")
	fmt.Println("Listen to : ", listener.Addr().String())
	return listener
}

func connectionHandler(conn net.Conn) {
	connFrom := conn.RemoteAddr().String()
	fmt.Println("Connection from : ", connFrom)
	talkToClients(conn)
	for {
		var ibuf []byte = make([]byte, maxread+1)
		length, err := conn.Read(ibuf[0:maxread])
		ibuf[maxread] = 0
		switch err {
		case nil:
			handleMsg(length, err, ibuf)
		default:
			goto DISCONNECT
		}
	}
DISCONNECT:
	err := conn.Close()
	fmt.Println("close connection : ", connFrom)
	checkError(err, "Close : ")
}

func handleMsg(length int, err error, msg []byte) {
	if length > 0 {
		for i := 0; ; i++ {
			if msg[i] == 0 {
				break
			}
		}
		fmt.Printf("Receive data : [%v], ", string(msg[0:length]))
		fmt.Println(" length: ", length)
	}
}

func talkToClients(to net.Conn) {
	wrote, err := to.Write(msgStart)
	checkError(err, "Write : wrote "+string(wrote)+" bytes.")
}

func checkError(error error, info string) {
	if error != nil {
		panic("Error : " + info + error.Error())
	}
}
