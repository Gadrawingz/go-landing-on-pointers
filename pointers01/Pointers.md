# POINTERS
https://www.developer.com/languages/pointers-in-golang/

## Meaning

In Go lang, pointers are special variables that are used to store the memory address of other variables and point to their memory address. Pointers also provide ways to access the value of the variable stored in that address.

**Pointer** is a variable that stores the address of another variable.

Unlike other variables that held values of a certain type, pointer holds the address of a variables.

A simple variable typically holds values such as intVar :=10 or strVar := “Hello”. But there are also variables that refer to functions, channels, methods, maps, slices, and so forth. The latter use a type of variable called references.

A pointer is also a variable but it holds the memory address of another variable. It is created to point to a particular type of variable.

This means that the address of an integer variable type can only be stored in a pointer variable declared of the same type.

This ensures that the Go compiler knows the size of bytes of the pointer to the value it occupies. As the address of the variable is stored in a pointer, we can modify the content of a variable through it.

## Memory Address

When we create a variable, a memory address is allocated for the variable to store the value.

In Go, we can access the memory address using the & operator. For example,


>>var num int = 5// prints the value stored in variable
>>fmt.Println("Variable Value:", num)
>>// prints the address of the variable
>>fmt.Println("Memory Address:", &num)

## What is the point of a pointer in Go?

One particular need for pointer variables is that they are cheap. This means they occupy a fixed size regardless of the size of the value they point to. The size of a pointer variable is 8-bytes for 64-bit machines and 4-bytes for 32-bit machines. Imagine a string variable which can be of any size but the pointer that points to it is of fixed size. 

This is important because we typically pass data to functions as arguments. If the function needs a local copy of the data, a normal variable would do. But if we want to alter the value of the original variable then instead of passing a variable by its value we may pass by reference through a pointer type variable. 

This ensures that even if we are in different function scope we are actually pointing to the original data.

Another aspect is that a variable passed by value actually creates a copy of itself. This occupies memory doubly. With pointers it actually references the original variable and does not create a copy and hence has less memory footprint.

Interestingly, in Go we can increase the lifetime of a variable independent of its scope using pointers. How? The internal memory management of Go will not claim the memory as long as there is at least one pointer pointing to that variable. (C/C++ do not have an internal memory manager, hence it is the programmer’s responsibility to do all the clean-ups.)

## Stackoverflow Overview
https://stackoverflow.com/questions/1863460/whats-the-point-of-having-pointers-in-go

Rather than answering it in the context of “Go”, I would answer this question in the context of any language (e.g. C, C++, Go) which implements the concept of "pointers"; and the same reasoning can be applied to “Go” as well.

There are typically two memory sections where the memory allocation takes place: the Heap Memory and the Stack Memory (let’s not include “global section/memory” as it would go out of context).

Heap Memory: this is what most of the languages make use of: be it Java, C#, Python… But it comes with a penalty called the “Garbage Collection” which is a direct performance hit.

Stack Memory: variables can be allocated in the stack memory in languages like C, C++, Go, Java. Stack memory doesn’t require garbage collection; hence it is a performant alternative to the heap memory.

But there is a problem: when we allocate an object in the heap memory, we get back a “Reference” which can be passed to “multiple methods/functions” and it is through the reference, “multiple methods/functions” can read/update the same object(allocated in the heap memory) directly. Sadly, the same is not true for the stack memory; as we know whenever a stack variable is passed to a method/function, it is “passed by value”(e.g. Java) provided you have the “concept of pointers”(as in the case of C, C++, Go).

Here is where pointers come into picture. Pointes let “multiple methods/functions” read/update the data which is placed in the stack memory.

>In a nutshell, “pointers” allow the use of “stack memory” instead of the heap memory in order to process variables/structures/objects by “multiple methods/functions”; hence, avoiding performance hit caused by the garbage collection mechanism.

Another reason for introducing pointers in Go could be: Go is ought to be an "Efficient System Programming Language" just like C, C++, Rust etc. and work smoothly with the system calls provided by the underlying Operating System as many of the system call APIs have pointers in their prototype.

One may argue that it can done by introducing a pointer-free layer on top of the system call interface. Yes, it can be done but having pointers would be like acting very close to the system call layer which is trait of a good System Programming Language.


## Advantages of Pointer
-Pointer provide direct access to the memory.
-Pointer provide a way to returns more than one value to the functions
-Pointers provides an alternate way to access array elements.
-Pointers reduces the storage space and complexity of the program.
-Pointer reduces the execution of the program.

## Disadvantages of Pointer
-Pointer are slower than normal variables.
-Uninitialized pointer might cause segmentation fault.
-Dynamically allocated block needs to be freed explicitly. Otherwise, it would lead to memory task.
-If pointer bugs are updated with incorrect values, it might lead to memory corruption.
-Basically, pointer bugs are difficult to handle. So, use pointer effectively and correctly.
-By using * operator we can access the value of a variable through a pointer.

> Ref: https://www.ques10.com/p/68382/what-is-pointer-explain-its-advantages-and-disad-1/

## To summarize the pointer operations:
-Pointer variables have to be declared b4 they can be used.
-It's then assigned the address of another variable by using the & operator.
-The value of the data in the address is accessed using the * operator.
