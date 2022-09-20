package server

import (
	"fmt"
	"net"
	"time"
)

var connectList = make([]net.Conn, 0)

func pushToConnLink(conn net.Conn) {
	connectList = append(connectList, conn)
}

func pingPone() {
	for {
		i := 0
		activeConn := 0
		deadConn := 0
		for _, conn := range connectList {
			i++
			_, err := conn.Write([]byte("aaa"))
			if err != nil {
				conn.Close()
				deadConn++
			} else {
				activeConn++
			}
		}

		fmt.Printf("当前连接数 %d 活连 %d 死连 %d \r\n", i, activeConn, deadConn)
		time.Sleep(time.Second * 5)
	}

}
