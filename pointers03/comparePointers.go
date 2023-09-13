package main

import (
	"fmt"
)

func main() {

	/*
		In Go, Two pointers can be compared with each other.
		The two pointers will be equal if they point to the same memory
		location and hence have the same value. The == (equal to)
		and != (not equal to) operators are used for comparison.
	*/

	var kilos int = 34 // declare & init variable kilos

	// Declare pointer and initialize to the address of the variable kilos
	var pointerK1 *int = &kilos

	// Declare pointer and point to pointer pointerK1
	var pointerK2 **int = &pointerK1

	fmt.Println(&kilos)     // 0xc000018188
	fmt.Println(pointerK1)  // 0xc000018188
	fmt.Println(pointerK2)  // 0xc000012028
	fmt.Println(*pointerK2) // 0xc000018188
	// Time for comparison
	fmt.Println(pointerK1 == &kilos)     // true
	fmt.Println(*pointerK1 == kilos)     // true
	fmt.Println(pointerK2 == &pointerK1) // true
	fmt.Println(*pointerK2 == pointerK1) // true

}
