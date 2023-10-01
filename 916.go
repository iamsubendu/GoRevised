package main

import (
	"fmt"
)

func main() {
	num:= 1

	// their is no need of break in Go
	switch num{
	case 1:
		fmt.Print(1)
	case 2:
		fmt.Print(2)
	default:
		fmt.Print("none")
	}
}
