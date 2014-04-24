// GoTcp project main.go
package main

func main() {
	hostAndPort := "localhost:54321"
	listener := initServer(hostAndPort)
	for {
		conn, err := listener.Accept()
		checkError(err, "Accept : ")
		go connectionHandler(conn)
	}
}
