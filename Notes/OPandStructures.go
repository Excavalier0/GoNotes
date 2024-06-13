package main
import "fmt"

func main (){
   fmt.Println("Hello World")

   //Control Structures:----
   //If statements:---
   if err := SomeFunction(); err == nil {
    // do something
} else {
    return err
}
//As go has return continue break and goto we can write:--

if err := SomeFunction(); err != nil {
    return err
}
// do something
func myfunc() {
    i := 0
Here:
    fmt.Println(i)
    i++
    goto Here
}
// goto statements can be used to move the program
//control to the label in this case "Here" when goto executes
// it moves the control to here 
 sum := 0
for i := 0; i < 10; i++ {
    sum = sum + i
}
/*

    for init; condition; post { } - a loop using the syntax borrowed from C;
    for condition { } - a while loop, and;
    for { } - an endless loop.
*/

//RANGE:---
list := []string{"a", "b", "c", "d", "e", "f"}
for k, v := range list {
    // do something with k and v
}
//range returns the value of index and the value stored at the index (Yup it returns 2 values)
//The return value depends on where you use range as:
for pos, char := range "Gő!" {
    fmt.Printf("character '%c' starts at byte position %d\n", char, pos)
}
//It prints:
/*
character 'G' starts at byte position 0
character 'ő' starts at byte position 1
character '!' starts at byte position 3 
(Note that ő took 2 bytes, so ‘!’ starts at byte 3.)
*/


//ARRAYS:----
var arr [10]int
arr[0] = 42
arr[1] = 13

a := int[4]{1,22,3,4}
ab := int[2][2]{ {1,2} , {1,2} } // multidimensional array

//arrays need to have a defined length and a defined type


//SLICES:-----
// Similar to arrays but can grow in size
//Points to underlying array , Its points to the array


slice:= a[n:m] //New slice 
//variable a , starts at index n , ends before m and length is n-m
a := [...]int{1, 2, 3, 4, 5} 
s1 := a[2:4] 
s2 := a[1:5] 
s3 := a[:]   
s4 := a[:4]  
s5 := s2[:] 
s6 := a[2:4:5]
//append:
s0 := []int{0, 0}
s1 := append(s0, 2) 1
s2 := append(s1, 3, 5, 7) 
s3 := append(s2, s0...) // ... required to specify that adding a slice 
//Append returns a slice 

//Copy 
var a = [...]int{0, 1, 2, 3, 4, 5, 6, 7}
var s = make([]int, 6)
n1 := copy(s, a[0:]) 
n2 := copy(s, s[2:])
//It returns no. of elements copied (n1 = 6 , n2 = 4)

//MAPS:--
monthdays := map{string}int(
     "Jan": 31, "Feb": 28, "Mar": 31,
    "Apr": 30, "May": 31, "Jun": 30,
    "Jul": 31, "Aug": 31, "Sep": 30,
    "Oct": 31, "Nov": 30, "Dec": 31, 
) 
// Map that converts from a string to int
fmt.Println(monthdays["Dec"]) // to print days in Dec

year := 0
for _, days := range monthdays {
    year += days
}
fmt.Printf("Numbers of days in a year: %d\n", year)
// range returns the days (index is discarded with _) and loops over
// year is the sum of no of days in an entire year

monthdays["undecim"]= 30 // new element
monthdays["Feb"] = 29 // new element to existing key

present := monthdays["Jan"] // returns true if key is present

map:delete(monthdays,"Mar") // delete the "Mar" key

 }

