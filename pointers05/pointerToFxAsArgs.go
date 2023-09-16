package main

import "fmt"

/*
Passing Pointers to Functions as Arguments
------------------------------------------
A pointer can be passed as an argument to a function just like passing
any other variable. While creating the function, the argument has to be
declared as of type pointer.

Below fx is created to accept a pointer as its argument.
*/

func acceptPointerAsArg(pointerVar *int) {
	*pointerVar += 10 // OR *pointerVar = *pointerVar + 10
}

func main() {

	//  1. Some testing
	var n1 int = 25
	fmt.Println("Before function call, n1: ", n1) // 25

	// Assign the memory address of n1 to pointer n1Pointer and
	var n1Pointer *int = &n1

	// Call the fx and pass pointer n1Pointer as its argument.
	acceptPointerAsArg(n1Pointer)
	// 2. Test again!
	fmt.Println("After. function call, n1: ", n1) // 35

	// Pass address of variable as pointer parameter
	acceptPointerAsArg(&n1)
	// 3. Test again!
	fmt.Println("After. function call, n1: ", n1) // 45

	// Pass address of variable as pointer parameter
	acceptPointerAsArg(&n1)
	// 4. Test again!
	fmt.Println("After. function call, n1: ", n1) // 55

}
