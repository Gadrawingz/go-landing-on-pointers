package main

import "fmt"

func main() {

	/*
		Pointer to a pointer:
		It is declared with two asterisks prefixed to the type e.g.
		var ptrB **int. In the above figure, pointer ptrA contains the address
		of the variable X and pointer ptrB holds the address of the pointer ptrA.
		You can get the value of X using *ptrA. To get the the value of X from
		the double pointer ptrB, use **ptrB.
	*/

	var myVariable int = 240

	// Declare pointer & initialize to the address of the variable
	var pointerOne *int = &myVariable

	// Declare pointer and point to pointer ptrA
	var pointerTwo **int = &pointerOne

	fmt.Println(pointerOne)
	fmt.Println(pointerTwo)

	fmt.Printf("Value of variable myVariable: %d\n", myVariable)    // 240
	fmt.Printf("Address of variable myVariable: %d\n", &myVariable) // 824633819528
	fmt.Printf("Value of pointer pointerOne: %d\n", pointerOne)     // 824633819528
	fmt.Printf("Address of pointer pointerOne: %d\n", &pointerOne)  // 824633794600
	fmt.Printf("Value of pointer pointerTwo: %d\n", pointerTwo)     // 824633794600

	// Dereferencing of pointers. Getting the value of X
	fmt.Printf("Value of myVariable from pointerOne : %d\n", *pointerOne)  // 240
	fmt.Printf("Value of myVariable from pointerTwo : %d\n", **pointerTwo) // 240

}
