package main

import (
	"fmt"
	"slices"
)

// using maps package

func main() {
	// map[key]value type
	m := make(map[string]int);
	m["k1"] = 7;
	m["k2"] = 13;
	fmt.Println("map : " , m);
	v1 := m["k1"];
	fmt.Println("val1 is : " , v1);
	delete(m , "k2");
	fmt.Println("deleted k2");
	fmt.Println(m);
	// clearing m 
	clear(m);
	fmt.Println(m);
	_ , prs := m["k2"];
	fmt.Println(prs);

	// declare and initialize both at same time

	count := map[string]int {"sss" : 3 , "aas" : 2};
	fmt.Println(count);

	


}