package main

import "fmt"

type  Direction int
const (
	North Direction = iota
	North2
	North3
	North4
)

func main() {
	fmt.Printf("North: %d, North2: %d,  North3: %d, North4: %d", North, North2, North3, North4)
}