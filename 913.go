package main

import "fmt"

var a=10 //package level

func demo(){
	var a=9 //functional level
	fmt.Println("demo",a)
}

func main(){
	demo()
	fmt.Println("main",a)
}