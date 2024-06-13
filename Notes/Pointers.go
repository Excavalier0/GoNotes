//Pointers:---
p := &i
//Here p is the assigned the address of i
//New and make func

/*
The built-in function new is essentially the same as its namesakes in other languages: new(T) allocates zeroed storage for a 
new item of type T and returns its address, a value of type *T. Or in other words, it returns a pointer to a newly 
allocated zero value of type T. This is important to remember.

The documentation for bytes.Buffer states that “the zero value for Buffer is an empty buffer ready to use.”.
Similarly, sync.Mutex does not have an explicit constructor or Init method. Instead, the zero value for a sync
.Mutex is defined to be an unlocked mutex.

Allocation with make

The built-in function make(T, args) serves a purpose different from new(T). It creates slices, maps, and channels only, 
aid it returns an initialized (not zero!) value of type T, and not a pointer: *T. The reason for the distinction is that these three 
types are, under the covers, references to data structures that must be initialized before use. A slice, for example, is a three-item descriptor
containing a pointer to the data (inside an array), the length, and the capacity; until those items are initialized, the slice is nil. For slices, maps, and channels, make
initializes the internal data structure and prepares the value for use.

For instance, make([]int, 10, 100) allocates an array of 100 ints and then creates a slice structure with length 10 and a capacity of 100 pointing at the first
10 elements of the array. In contrast, new([]int) returns a pointer to a newly allocated, zeroed slice structure, 
that is, a pointer to a nil slice value. These examples illustrate the difference between new and make.*/ 

var p *[]int = new([]int)       
var v  []int = make([]int, 100) 

var p *[]int = new([]int)       
*p = make([]int, 100, 100)

v := make([]int, 100)


//    new(T) returns *T pointing to a zeroed T
//    make(T) returns an initialized T

//STRUCTS:____
package main

import "fmt"

type NameAge struct {
	name string // Both non exported fields.
	age  int
}

func main() {
	a := new(NameAge)
	a.name = "Pete"
	a.age = 42
	fmt.Printf("%v\n", a)
}

struct {
    T1        // Field name is T1.
    *T2       // Field name is T2.
    P.T3      // Field name is T3.
    x, y int  // Field names are x and y.
}

//Methods

