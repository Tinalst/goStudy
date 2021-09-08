package main

import "fmt"

// 与js的...rest一样
func foo(arg...int) {
	fmt.Println(arg[0], arg[1])
	fmt.Println(arg[0], arg[1], arg[2])
}

func main() {
	foo(1,2,3)
}