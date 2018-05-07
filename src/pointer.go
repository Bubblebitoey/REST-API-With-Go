package main

import "fmt"

func main() {
	
	x := 15

	a := &x         // point to x 
	fmt.Println(a)
	fmt.Println(*a)
	
}