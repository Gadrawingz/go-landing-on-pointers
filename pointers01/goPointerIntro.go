package main

import "fmt"

func main() {
	fmt.Println("Getting on with Pointers!")

	// WHAT IS POINTER?
	// It is a variable that stores the memory address of another variable.
	var XVariable int = 394
	var XPointer *int = &XVariable
	// Here, the variable 'XVariable' has value 394 and is stored at memory
	// address '0xc0000be000'. The var XPointer holds the address of XVariable.
	// Now XPointer is said to point to XVariable.
	fmt.Println("XVariable's value:", XVariable)    // 394
	fmt.Println("XVariable's address:", &XVariable) // 0xc0000be000
	fmt.Println("XPointer.'s value:", XPointer)     // 0xc0000be000

	// 1. Declaring Pointers
	// *T is the type of the pointer variable w/c points to a value of type T.
	// Variable declaration
	var variable10 = "Wayne"
	var variable20 = 58
	var variable30 = 7.80
	var pointer01 *string  // pointer variable of string type
	var pointer02 *int     // pointer variable of complex type
	var pointer03 *float64 // pointer variable of double type
	pointer01 = &variable10
	pointer02 = &variable20
	pointer03 = &variable30
	fmt.Println(pointer01)
	fmt.Println(pointer02)
	fmt.Println(pointer03)

	// 2. Memory address
	// When we create a variable, a memory address is allocated for the
	// variable to store the value.
	// => The & operator, used 2get memory address of a variable.

	// 3. How 2access(get) value pointed by pointers in Go?
	// => The * operator, used 2access the value present in the memory address
	// pointed by the pointer.

	var age int = 24
	// Prints the value stored in variable
	fmt.Println("Variable Value:", age) // Variable Value: 24
	// Prints the address of the variable
	fmt.Println("Memory Address:", &age) // Memory Address: 0xc000128048

	var firstName = "Iradufasha"
	fNamePointer := &firstName
	fmt.Println(fNamePointer) // 0xc000014270
	// 2get the value pointed by fNamePointer we can use *
	fmt.Println(*fNamePointer) // Iradufasha

	// Note1: Above, fNamePointer is a pointer, not *fNamePointer.
	// You cannot and should not do something like *fNamePointer = &firstName

	myVariable := 300 // Variable
	// Both ways is for assigning the variable addr 2the pointer
	var myPointer1 *int = &myVariable
	myPointer2 := &myVariable

	fmt.Printf("Type of myVariable is %T \n", myPointer1)
	fmt.Println("The address of myVariable is", myPointer1) // 0xc0000be000
	fmt.Println("The address of myVariable is", myPointer2) // 0xc0000be000

	// NB: This hexadecimal memory address of the variable 'myVariable'
	// can be different in different machine/OS. So,
	// All &myVariable, myPointer1 and myPointer2 display the same address.

	// 4. Dereferencing a pointer:
	// Dereferencing a pointer means accessing the value of the variable to which the pointer points. *pointer1 is the syntax to deference pointer1.

	// The * is called 'dereference operator' when working with pointers. It operates on a pointer and gives the value stored in that pointer.

	variable1 := 430
	pointer1 := &variable1
	fmt.Println("--------------------------")
	// We deference pointer1 and print its value. As expected it prints the value of variable1.
	fmt.Println("Address of variable1 is", pointer1) // 0xc000018188
	fmt.Println("value of variable1 is", *pointer1)  // 0xc000018188

	// 5: Nil Pointers (Zero value of a pointer)
	// The zero value of a pointer is nil.
	// If a pointer variable is not assigned to an address of a variable
	// then it is called nil pointer.
	// -Any uninitialized pointer will have <nil> value.

	var intPointer *int
	var newPiPointer *float64
	fmt.Println(intPointer)   //output: <nil>
	fmt.Println(newPiPointer) //output: <nil>

	// 6. Deep Working of Pointers in Golang
	var academicYear int32 // Variable
	var yearPointer *int32 // Pointer

	academicYear = 2021
	fmt.Println("Address of academicYear is:", &academicYear) //0xc0000181a0
	fmt.Println("The Value of academicYear :", academicYear)  // 2021

	yearPointer = &academicYear
	fmt.Println("Address of yearPointer is: ", yearPointer)  //0xc0000181a0
	fmt.Println("The Value of academicYear :", *yearPointer) // 2021

	academicYear = 2022
	// 0xc0000181a0
	fmt.Println("Address of pointer yearPointer is:", yearPointer)
	fmt.Println("Content of pointer yearPointer is:", *yearPointer) // 2022
	fmt.Println("***************")

	academicYear = 2024
	// 0xc0000181a0
	fmt.Println("Address of pointer yearPointer is:", yearPointer)
	fmt.Println("Content of pointer yearPointer is:", *yearPointer) // 2024
	fmt.Println("***************")

	*yearPointer = 2070
	fmt.Println("Address of pointer academicYear is:", &academicYear)
	fmt.Println("Content of pointer academicYear is:", academicYear) // 2022
	fmt.Println("***************")

	// 7. Creating pointers using the new function.
	// The new function takes a type as an argument and returns a pointer to
	//a newly allocated zero value of the type passed as argument.
	theSize := new(int)
	fmt.Printf("Value: %d, Type: %T, Address: %v\n", *theSize, theSize, theSize) // Value: 0, Type: *int, Address: 0xc0000181c8
	*theSize = 132
	fmt.Printf("New Value: %d, Type: %T, Address: %v\n", *theSize, theSize, theSize) // New Value: 132, Type: *int, Address: 0xc0000181c8

	var nPointer = new(float64)
	*nPointer = 40.75

	fmt.Println(nPointer)  // 0xc0000181d0
	fmt.Println(*nPointer) // 40.75

	// 8 Again, How to create pointer without * ?
	university := "Harvard"
	// To directly assign the &university(address of university)
	// to the noStarPointer variable.
	var noStarPointer = &university
	fmt.Println("________________________")
	fmt.Println("Value :", noStarPointer)  // Value : 0xc0000a6250
	fmt.Println("Value :", *noStarPointer) // Value : Harvard
	fmt.Println("Address of university :", &noStarPointer)
	// Address of university : 0xc000012030

}
