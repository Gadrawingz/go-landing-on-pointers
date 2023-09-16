package main

import "fmt"

// Returning pointer from a function
// =================================
// It is perfectly legal for a fx to return a pointer of a local variable.
// The Go compiler is intelligent enough and it will allocate this variable
// on the heap.

func bringLocalVarAddress() *int {
	number := 10
	return &number
}

func main() {

	// Testing (1)
	num := bringLocalVarAddress()
	fmt.Println("Value of num: ", num) //  0xc000018188

	/*
		In line no. 9 of the program above, we return the address of the local variable i from the function hello. The behavior of this code is undefined in programming languages such as C and C++ as the variable i goes out of scope once the function hello returns. But in the case of Go, the compiler does an escape analysis and allocates i on the heap as the address escapes the local scope. Hence this program will work and it will print,
	*/

}
