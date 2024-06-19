package main

import "fmt"

type base struct {
	num int
}

type container struct {
	// EMBEDDING A STRUCT INSIDE OF STRUCT 
	base 
	str string
}

func (b base) describe() string {
	return fmt.Sprintf("base with num = %v" , b.num);
}

func main() {

	co := container {
		base : base {
			num : 10,
		},
		str : "bharath",
	}

	fmt.Println(co.num);
	fmt.Println(co.base.num); // accessing via full path
	fmt.Println(co.describe());

	type describer interface {
		describe() string
	}

	var d describer = co;
	fmt.Println(d.describe());

}