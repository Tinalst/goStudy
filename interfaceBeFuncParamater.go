package main

import (
	"fmt"
	"strconv"
)

type Human struct {
	name string
	age int
	phone string
}

func (h Human) String() string  {
	return "❰"+h.name+" - "+strconv.Itoa(h.age)+" years -  ✆ " +h.phone+"❱"
}


func main() {
	Bob := Human{
		name: "bob",
		age: 20,
		phone: "11122-2222",
	}

	fmt.Println(Bob)
}