package main

import "fmt"

func main(){
	var num int
	fmt.Println(num)
	var num2 uint = 3
	fmt.Print(num2)
}
// types
// bool
// string
// int int8 int16 int32 int64
// uint uint8 uint16 uint32 uint64

// byte-> alias for unit8

// rune-> alias for int32

// float32 float64

//by default int value is 0

// bool: Represents boolean values, which can 
// be either true or false.

// string: Represents a sequence of characters, 
// such as text. Strings are enclosed in double
// quotes, like "Hello, World!".

// int, int8, int16, int32, int64: Represent signed 
// integer types with different sizes. The number 
// indicates the number of bits used to represent 
// the integer (e.g., int32 uses 32 bits). int is 
// the platform-dependent integer type.

// uint, uint8, uint16, uint32, uint64: Represent 
// unsigned integer types with different sizes. They 
// are used to store positive whole numbers and have 
// the same bit size distinctions as their signed counterparts.

// byte: An alias for uint8, used to represent small 
// positive integers, often used for storing ASCII characters.

// rune: An alias for int32, used to represent Unicode 
// code points, which are used for characters in text.

// float32, float64: Represent floating-point numbers 
// with single and double precision, respectively. They 
// are used for real numbers and can store decimal values.