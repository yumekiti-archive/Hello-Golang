package main

import (
	"fmt"
)

func main() {

	var a int = 123

	var b = 456

	c := 789

	var (
		d int = 123
		e int = 456
	)

	var f int
	f = 123

	var g int
	var h string
	g, h = 456, "hoge"

	fmt.Printf("test %d %d %d %d %d %d %d %s", a, b, c, d, e, f, g, h)
	fmt.Printf("\nend\n")

}
