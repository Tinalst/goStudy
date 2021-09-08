package main

import "fmt"

func main() {
	product := 1
	for i := 1; i < 5; i++ {
		product *= i
	}

	fmt.Printf("product: %d", product)
}