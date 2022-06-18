// exported name are all capitalized
// ex: fmt.Println math.Pi
package main

import (
	"fmt"
	"math"
)

func main ()  {
	// wrong: math.pi
	fmt.Println(math.Pi)
}