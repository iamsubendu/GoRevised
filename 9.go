package main

import "fmt"

func add(x int, y int) int {
	//add(parameter) return valueType
	res:= x+y
	return res
}

func main(){
	num1:= 2
	num2:=3
	result:= add(num1,num2)
	fmt.Print(result)
}