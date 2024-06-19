package main

import (
	"fmt"
	"slices"
)

func main() {
	var s1 []string;
	fmt.Println(s1 , len(s1) , cap(s1));

	var s = make([]string , 3);

	s[0] = "Gu";
	s[1] = "ru";
	s[2] = "murthy";
	fmt.Println("Set : " , s);
	fmt.Println("get : " , s[2]);
	fmt.Println(len(s));
	// fmt.Println(len(s));
	var s2 = append(s , "B" , " " , "Bharath");
	fmt.Println(s2);
	fmt.Println(s);
	l := s2[2:5];
	fmt.Println(l);

	// declare and initialize at same time
	s4 := [] string { "anil" , "pallavi" , "gurumurthy" };
	s5 := [] string { "anil" , "pallavi" , "gurumurthy" };
	if slices.Equal(s4,s5) {
		fmt.Println("Slices are equal");
	}
	fmt.Println(s4);

	twoD := make([][] int , 3);
	fmt.Println(twoD);
	for i := 0 ; i < 3 ; i++ {
		innerLen := i + 1;
		twoD[i] = make([]int , innerLen);
		for j := 0 ; j < innerLen ; j++ {
			twoD[i][j] = i + j;
		}
	}
	fmt.Println(twoD);
	

}