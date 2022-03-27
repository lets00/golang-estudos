// Programs start running in package main
package main

// You can also import multiple statement like:
// import "fmt"
// import "math/rand"
// good style below
import (
	"fmt"
	"math/rand"
)

// Intn is in files that begin with statement package rand
func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
}
