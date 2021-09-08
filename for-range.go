package main

import "fmt"

func eachArray() {
	for k, v := range []int{0, 1,2,3} {
		fmt.Printf("k: %d, v: %d", k, v)
	}
}

func eachString() {
	var str string = "abcdefe"
	for k, v := range str {
		fmt.Printf("k: %d, v: 0x%x\n", k, v)
		fmt.Println(k,v)
	}
}

func eachMap() {
	m := map[string]int {
		"go": 100,
		"web": 100,
	}

	for k, v := range m {
		fmt.Println(k, v)
	}
}

func eachChannel() {
	// 创建通道
	c := make(chan int)
	// 开启通道，赋值
	//顺序从上到下
	go func(){
		c <- 2
		c <- 1
		// 关闭通道
		close(c)
	}()

	// 通道只能获得值
	for v := range c {
		println(v)
	}
}

func main() {
	//eachArray()
	//eachString()
	//eachMap()
	eachChannel()
}