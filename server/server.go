package server

import (
	"fmt"
	"net"
)

func handleConnect(conn net.Conn) {
	defer conn.Close()
	//加入到链表
	pushToConnLink(conn)
	//消息buffer
	buffer := make([]byte, 1)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Printf("read from connect failed, err: %v\n", err)
			break
		}
		fmt.Printf("receive from client, data: %v\n", string(buffer[:n]))
	}
}

func NewServer() {
	listenr, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}
	fmt.Print("=============== listen 9000  \r\n")
	go pingPone()
	for {
		conn, err := listenr.Accept()
		if err != nil {
			panic(err)
		}
		go handleConnect(conn)
	}
}
