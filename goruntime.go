package main

import (
	"fmt"
	"runtime"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		for i := 0; i < 5; i++ {
			// 让cpu把事件片让给别人，下一次恢复执行
			runtime.Gosched()
			fmt.Println(s)
		}
	}
}

func main()  {
	go say("world") // 开一个新的 Goroutines 执行
	say("hello") // 当前 Goroutines 执行
}