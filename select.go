package main

import "fmt"

func main()  {
	c := make(chan int)
	quit := make(chan int)

	go func(){
		for i := 0; i < 10; i++ {
			// 读取通道数据
			fmt.Println(<-c)
		}

		// 通道写入数据
		quit <- 0
	}()


}