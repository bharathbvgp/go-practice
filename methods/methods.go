package main

import "fmt"

// define a struct of rect with two fields

type rect struct {
	width , height int
}

// create two associated methods for calculating area and perimeter 

func area(r *rect) int {
	return r.height * r.width;
}

func perimeter(r *rect) int {
	return 2*r.height + 2*r.width;
}

func main() {
	// create an instance of rect struct
	r := rect {
		width : 10,
		height : 10,
	}
	fmt.Println("area is : ",area(&r));
	fmt.Println("perimeter is : " , perimeter(&r));
}