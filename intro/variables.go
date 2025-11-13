package main

import (
	"fmt"
)

func main() {
	/// neewd the type when defining
	var blankvar int
	fmt.Println(blankvar)

	var a = "initial"
	fmt.Println(a)

	var astring string = "this is a string"
	fmt.Println(astring)

	var b, c int = 1, 2
	fmt.Println(b, c)

	// can only be used in functions
	// can only be used when initializing
	myvar := 100
	fmt.Println(myvar)

	// constants
	const constvar string = "golang"
	const otherconst = "golangagain"
	const intconst = 12.34e2
	const constmath = intconst / 90
	const lastcmath = intconst / constmath
}
