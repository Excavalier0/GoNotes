// Main Struct
package main

import "fmt"

func main() {
	fmt.Println("Hello World")

	//All go files start with package name (In this case for a standalone main)

	// First func to be executed is main.main() which is similar to C
	// non ASCII characters are to be included in fmt.Println
	var a int
	a = 2
	y := true

	var (
		x int
		s string
	)
	//Declaring Variables in golang
  var b int32 //32 bit machines and int64 for 64 bit machines

  /*If you want to be explicit about the length, you can have that too
  , with int32, or uint32. The full list for (signed and unsigned) integers
  is int8, int16, int32, int64 and byte, uint8, uint16, uint32, uint64, with byte 
  being an alias for uint8. For floating point values there is float32 and float64 (there is no float type)
  . A 64 bit integer or floating point value is always 64 bit, also on 32 bit architeturestures.*/

	//int32 and int are 2 distinct types(Goes for all others..)

  //CONSTANTS:--
  const marine int = 4
  // const value cant be mutated and has to be assigned a value at Declaring
  str := "Hello World"
  //strings are immutable in go must be assigned at Declaring

	var dumb rune 
	//rune is an alias for int32
var(
 r2 complex64 // 32 bit real and imaginary parts
 r1 complex128 // 64 bit real and imaginary parts
 )
 // Errors:--
 error.var e error
 //Creates a var e of type error with a value of nil

}
