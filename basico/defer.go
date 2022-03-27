package main

import "fmt"

// defer runs in last-in-first-out order
func main() {
	defer fmt.Println("World")
	defer fmt.Println("!")
	defer fmt.Println("@")
	defer fmt.Println("#")
	fmt.Println("Hello")
}
