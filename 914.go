package main

import (
	"fmt"
	"math"
)

func main(){
	var num float64=18
	fmt.Println(num)
	res:= math.Sqrt(num)
	//for using Sqrt we can use int type
	fmt.Println("Sqrt is :", res)
	fmt.Printf("Sqrt is :", res)
	fmt.Println()
	// we can use Printf, but have to use placeholder
	fmt.Printf("Sqrt is %f", res)
	fmt.Println()
	fmt.Printf("Sqrt is %g", res)
	fmt.Println()
	fmt.Printf("Sqrt is %.2g", res)
	fmt.Println()
	fmt.Printf("Sqrt is %.3g", res)
	fmt.Println()
	fmt.Printf("Sqrt is %.2f", res)
	fmt.Println()
	fmt.Printf("Sqrt is %.3f", res)
	fmt.Println("==============")
	var intRes = math.Round(res)
	fmt.Println(intRes)
	fmt.Printf("Sqrt is %.3f", intRes)
	fmt.Println("==============")
	var intRes2 = math.Ceil(res)
	//Ceil will go up
	fmt.Println(intRes2)
	fmt.Printf("Sqrt is %.3f", intRes2)
}