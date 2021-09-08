package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a string = "hahahaa"

	fmt.Println(strconv.FormatBool(len(a) > 6))

	switch a {
	case "love", "ha" :
		fmt.Println("love")
	case "ps":
		fmt.Println("ps")
	// 这样写不行，因为case的类型必须是switch 里的类型下面这个虽然会返回true 或者false的string类型，但是a是等于hahaha的所以这里
	// 永远不会被触发
	case strconv.FormatBool(len(a) > 6):
		fmt.Println("6666")
	default:
		fmt.Println("default")
	}
}