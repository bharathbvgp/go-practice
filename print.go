package main

import "fmt"

func main() {
	const name = "bharath";
	const age = 23;
	// fmt.Printf("My name is : %s , age is  : %d" , name , age);
	fmt.Printf("My name is : %s , age is  : %d and type is : %T and %T" , name , age , name , age);
}