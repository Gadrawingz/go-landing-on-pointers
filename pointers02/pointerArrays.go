package main

import "fmt"

func main() {

	theArray := []int{100, 200, 300, 400}

	var k int // for increment

	// We can create array which stores pointers in Go.
	// pointerArr is an array of int type pointers that can store 3 pointers
	var pointerArr [4]*int

	for k = 0; k < 4; k++ {
		// Assign the address of each element of theArray to pointerArr.
		pointerArr[k] = &theArray[k]
	}

	for k = 0; k < 4; k++ {
		fmt.Printf("*pointerArr[%d] = %d\n", k, *pointerArr[k])
	}
}
