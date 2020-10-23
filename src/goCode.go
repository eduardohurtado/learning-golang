package main

import "fmt"

var i float32 = 18

var (
	actorName    string = ""
	companion    string = ""
	doctorNumber int    = 22
	season       int    = 3
)

const (
	a3 = iota
	a4 = iota
	a5 = iota
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

	// Any var initialized has 0 or false as value
	var poop int
	fmt.Println("Default value:", poop)

	// Numeric types

	// Unsigned integer
	var noSign uint = 12
	fmt.Println("Unsigned:", noSign)

	// Logical operators
	fmt.Println("Logical operators:")
	a := 10 // 1010
	b := 3  // 0011

	fmt.Println(a & b)  // 0010
	fmt.Println(a | b)  // 1011
	fmt.Println(a ^ b)  // 1001
	fmt.Println(a &^ b) // 0100

	// Bit shift ONLY INTEGER
	fmt.Println("Bit shift")
	a = 8               // 2^3
	fmt.Println(a << 3) // 2^3 * 2^3 = 2^6
	fmt.Println(a >> 3) // 2^3 / 2^3 = 2^0

	// Type string
	s := "This is a string"
	fmt.Println(s)
	fmt.Printf("%v, %T\n", string(s[2]), s[2])

	// Constants, you cant set your constants at run time ex: const a int = func(1.57)
	fmt.Println("My constant:")
	const myConst int = 98
	fmt.Printf("%v, %T\n", myConst, myConst)
}
