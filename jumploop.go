package main

import "fmt"

func main() {
	JumpLoop:
		for j := 0; j < 5; j++ {
			for i :=0; i < 5; i++ {
				if i > 2 {
					break JumpLoop
				}
				fmt.Println(i)
			}
		}

	fmt.Println("here")

	// δΌεηζ¬
}