package main

import "fmt"


func main() {
	nums := [5]int {1,2,3,4,5};
	sum := 0
	for _ , num := range nums {
		sum += num;
	}
	fmt.Println("the sum of array is  : " , sum);

	family := [4]string {"anil" , "bharath" , "gurumurthy" , "pallavi"};
	for _ , members := range family {
		fmt.Println(members);
	}

	// finding index

	idx := -1
	for i , num := range nums {
		if num == 3 {
			idx = i;
		}
	}
	fmt.Println(idx);


	// iterating over map using range

	m := map[string]int {"anil" : 4 , "pallavi" : 6 , "raghu" : 5};
	for k,v := range m {
		fmt.Println("key and value is : " , k , v);
	}

	for k := range m {
		fmt.Println(k);
	}

	for i , c := range "bharath" {
		fmt.Println(i , string(c));
	}


}


