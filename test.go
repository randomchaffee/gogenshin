// you can edit this code!
// click here and start typing.
package main

import "fmt"

// return an int of num raised to pow
func exp(num int, pow int) int {
	result := num

	for i := pow; i > 0; i -= 1 {
		result = result * num
	}

	return result
}

func test() {
	x := 0
	fmt.Printf("%v\n%T\n", x, x)

	fmt.Printf("%v raised to %v is %v\n", 2, 5, exp(2, 5))
}
