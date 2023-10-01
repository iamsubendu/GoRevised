package main

import "fmt"

func calc(x, y int) (res1,res2 int) {
	res1= x+y
	res2= x-y
	return
}

func main(){
	num1:= 2
	num2:=3
	result1,result2:= calc(num1,num2)
	fmt.Print(result1,",",result2)
}