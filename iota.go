package main

import (
	"fmt"
)

// this is a const block to demo "iota"
// also these constants are available for the whole package
// TODO: explore iota of other data types
const (
	first  = iota
	second //iota auto increases this value within this constant block
)

const (
	//iota gets reset and starts again from 0 since this is a new constant block
	third  = iota
	fourth = iota
)

func main() {
	fmt.Println(first, second, third, fourth) //prints 0, 1, 0, 1
}
