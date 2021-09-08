package main

import "fmt"

// 声明一个函数类型
type testInt func(int) bool

// 把函数类型当作另外一个函数的参数类型
// 回调的意思
func filter(slice []int, f testInt) []int {
	var result []int
	for _, v := range slice {
		// 当 f(v)返回true，追加到result切片中
		res := f(v)
		if res {
			result = append(result, v)
		}
	}

	return result
}


// 定义两个回调函数

// 奇数过滤器
func isOdd(num int) bool{
	if num%2 == 0 {
		return false
	}
	return true
}

// 偶数过滤器
func isEven(num int) bool  {
	if num%2 == 0 {
		return true
	}
	return false
}


func main()  {
	slice := []int{1,2,3,4,5}
	var res []int

	res = filter(slice, isOdd)
	fmt.Println(res)

	res = filter(slice, isEven)
	fmt.Println(res)
}