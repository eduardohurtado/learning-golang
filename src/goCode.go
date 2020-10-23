package main

import "fmt"

var i float32 = 18

var (
	actorName    string = ""
	companion    string = ""
	doctorNumber int    = 22
	season       int    = 3
)

func main() {

	// var u int = 44	// Never used error

	myvar := 12.0

	fmt.Println("Value types:")
	fmt.Printf("%v, %T\n", myvar, myvar)

	var foo int
	foo = int(myvar)
	fmt.Printf("%v, %T\n", foo, foo)

	// Boolean data
	var n bool = false
	l := 1 == 1
	m := 1 == 2

	fmt.Println("Boolean data:")
	fmt.Println(n)
	fmt.Println(l)
	fmt.Println(m)

	// Numeric types
}
