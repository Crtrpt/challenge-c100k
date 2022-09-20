package main

import (
	"fmt"

	"challenge-c100k.com/container"
)

func main() {

	link := container.NewLinkList()
	link.PushHead(20)
	link.PushHead(30)
	link.PushTail(40)
	link.PushTail(50)

	link.Filter(func(v any) bool {
		return v.(int) > 30
	})

	link.Each(func(a any) {
		fmt.Printf("数据 %v \n", a)
	})
	fmt.Printf("长度 %d", link.Len())
	// sigs := make(chan os.Signal, 1)
	// done := make(chan bool, 1)
	// signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	// go func() {
	// 	<-sigs
	// 	done <- true
	// }()
	// go server.NewServer()
	// <-done
	// fmt.Println("应用程序已经退出")
}
