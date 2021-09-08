package main

import "fmt"

// Shaper 接口
type Shaper interface {
	Area() float32
}

// Square 结构体
type Square struct {
	side float32
}

func(sq Square) Area() float32 {
	return sq.side * sq.side
}

func main() {
	// 创建Square实例
	sq1 := new(Square)
	sq1.side = 5

	// var areaIntf Shaper
	// 因为结构体Square实现了Shaper接口，所以可以将这种结构体定义的变量赋值给接口类型变量
	// 那么这个接口类型变量就能使用结构体实现了的接口方法了
	//areaIntf := Shaper(sq1)
	fmt.Println("sq1 area:", sq1.Area())
}