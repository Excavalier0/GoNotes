package main

import "fmt"

func main() {
   func (p mytype) funcname(q int) (r,s int) { return 0,0 }
   // p is used to bind type to reciever
   // q represents func type
   // r,s represents return values(if not specified returns final parameter value)
   //nesting is not allowed in funcs
   //Can be declared in any order as go reads the whole file bfr compile
   func rec(i int) {
    if i == 10 { 1
        return
    }
    rec(i+1) 2
    fmt.Printf("%d ", i)
}
 //SCOPE :== 
/*
Variables declared outside any functions are global in Go, those defined in 
functions are local to those functions. If names overlap - a local variable is declared 
with the same name as a global one - the local
variable hides the global one when the current function is executed.
*/

//FUNCS AS VALUES:--=
import "fmt"

func main() {
	a := func() { 
		fmt.Println("Hello")
	} 
	a() 
}
//Here the var a is given the value of an anonymous func 
//As Maps:--
:

var xs = map[int]func() int{
    1: func() int { return 10 },
    2: func() int { return 20 },
    3: func() int { return 30 },
}
//Callbacks:--
//Define a function that does something:
func printit(x int) {
    fmt.Printf("%v\n", x)
}
/*

This function does not return a value and just prints its argument.
The signature of this function is: func printit(int), or without the function
name: func(int). To create a new function that uses this one as a callback we need 
to use this signature: */
func callback(y int, f func(int)) {
    f(y)
}
//Here the func callback takes 2 parameters y int and the func(int) which returns nothing
//"f" is the variable holding that func and can be used as shown(taking parameter y)

//DEFERRED CODE:--

//Defer when executed moves the func to a stack which follows LIFO(last in first out)
//Defer when called moves line to the end :-
defer fmt.Println("World")
fmt.Println("Hello")
//The output will still be Hello /n World
//In a loop 
for i:= 0 ; i< 5 ; i++ {
  defer fmt.Println(i)
}
//The output will be 43210 instead of 01234 because of defer LIFO

//Variadic parameters(Infinite parameters)
func myfunc(arg ...int) {}
//Here myfunc takes an Infinite amount of int parameters
//Those Parameters are stored as slices
//arg ... packs the parameters into arg
for _, n := range arg {
    fmt.Printf("And the number is: %d\n", n)
}
//Print the parameters
xint := []int{1,2,3,4}
myfunc(xint...) //unpacks the slice and passes to myfunc

}
//PANIC AND RECOVERY:----
//IT stops the normal flow of the program 
// any DEFERRED funcs are executed normally and rest is stopped


func Panic(f func()) (b bool) { 
    defer func() { 
        if x := recover(); x != nil {
            b = true
        }
    }()
    f() 
    return 
}
//Recover returns nil on normal execution
//And captures the panic 
