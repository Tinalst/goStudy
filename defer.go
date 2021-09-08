package main

import "fmt"

func main() {
	// defer执行顺序  - 先进后出
	// 函数内又return，对于外部调用该函数，先return了，再执行defer
	for i := 0; i < 3; i++ {
		defer fmt.Println(i) // 输出顺序应该是 - 2 1 0
	}
}