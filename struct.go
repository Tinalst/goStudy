package main

import "fmt"

// 声明一个struct类型
type person struct {
	name string
	age int
}



func main()  {
	var tom person
	// 初始化 方式1
	tom.name, tom.age = "TOM", 18

	// 初始化 方式2
	bob := person{"Bob", 19}

	// 初始化 方式3
	paul := person{age: 20, name: "paul"}

	// 初始化 方式4
	alice := new(person)
	alice.age = 21
	alice.name = "Alice"

	fmt.Println(tom, bob, paul, alice)
}