package main

import "fmt"




func main() {
	var a [5]int;
	// initially all zeros [0,0,0,0,0]
	fmt.Println(a);
	a[4] = 100;
	fmt.Println("set a[4] - ");
	fmt.Println(a);
	fmt.Println("get a[4] - " , a[4]);
	// printing len of a 

	fmt.Println("length of a is  : " , len(a));

	// initializing and decalring both at the same time
	b := [5]int {1,2,3,4,5};
	fmt.Println(b);

	// compiler will count the elements 
	b = [...]int {100,2,3,400,6};
	fmt.Println(b);

	// if we specify index elements will be zero in between 
	b  = [...]int{100, 3: 400, 500};
	fmt.Println(b);

	// two dimensional arrays

	var twoD [3][3] int;
	cnt := 1;
	for i := 0 ; i < 3 ; i++ {
		for j := 0 ; j < 3 ; j++ {
			twoD[i][j] = cnt;
			cnt += 1;
		}
	}

	fmt.Println("2D array" , twoD);


	// declaration and initialization at the same time  

	second := [2][3] int {
		{1,2,3},
		{4,5,6},
	}
	fmt.Println(second);

	third := [2][2] int {
		{100,200},
		{300,400},
	}
	fmt.Println(third);




}