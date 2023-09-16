package main

import "fmt"

func changeName(name *string) {
	*name = "Donnekt"
}

func main() {

	// Testing (1)
	gadName := "Gad Durant"
	fmt.Println("Value before Fx call: ", gadName)

	// Testing (2)
	newName := &gadName
	changeName(newName)
	fmt.Println("Value after. Fx call: ", gadName)

	// LOGIC: Here, we are passing the pointer variable 'newName' w/c
	// holds the address of 'gadName' to the fx 'changeName'.
	// Inside 'changeName' fx, the value of gadName is changed using dereference.
}
