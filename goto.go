package main

import "fmt"

func main()  {
	for x := 0; x < 20; x++ {
		for y := 0; y < 20; y++ {
			if y == 2 {
				goto breakTag
			}
		}
	}
	breakTag:
		fmt.Println("here")
}