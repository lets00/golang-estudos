package main

import "fmt"

// Variable name comes before the type
func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
