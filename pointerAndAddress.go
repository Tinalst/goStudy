package main

import "fmt"

// a int - 值传递
func add(a int) int {
	a += 1
	return a
}

// a *int - 引用传递 - 拷贝的是一个指针
// *int - 指针类型
// 适合用在传递大的结构体 - 比较轻量
func add2(a *int) int {
	// *a - 直接操作地址上的值
	*a += 1
	return *a
}

func main() {
	x := 3

	x1 := add(x)

	fmt.Println(x1) // 4
	fmt.Println(x)  // 3

	// &x - 将x的引用地址传入
	x2 := add2(&x)

	fmt.Println(x2) // 4
	fmt.Println(x)  // 4
}