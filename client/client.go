package client

import "net"

func NewClient() {
	conn, err := net.Dial("tcp", ":9000")
	if err != nil {
		panic(err)
	}
	conn.Write()
}
