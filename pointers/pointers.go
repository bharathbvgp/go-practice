package main

import "fmt"

func zeroval(val int) {
	val = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1;
	fmt.Println("initial value : " , i);
	zeroval(i);
	fmt.Println("after calling zeroval");
	fmt.Println(i)
	// pass by reference 
	zeroptr(&i);
	fmt.Println("AFTER CALLING ZEROPTR");
	fmt.Println(i);
}