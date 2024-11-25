package main

import (
	"fmt"
)

func main() {
	array := [4]int8{1, 3, 3, 7}

	slice := array[:]

	slice[0] = 9

	fmt.Print(array) // [9,3,3,7]

}
