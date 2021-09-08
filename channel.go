package main

import "fmt"

func sum(a []int, c chan int) {
	total := 0
	for _, v := range a {
		total += v
	}
	c <- total
}

func main()  {
	a := []int{1, 1, 1, 2, 2, 2}

	c := make(chan int)
	go sum(a[:len(a) / 2 ], c)
	go sum(a[len(a) / 2:], c)

	x, y := <-c, <-c

	fmt.Println(x, y, x+y)
}