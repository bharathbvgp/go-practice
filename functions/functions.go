package main

import "fmt"

// variadic functions 

func sum (nums ...int) int {
	total := 0;
	fmt.Println(nums);
	for _ , v := range nums {
		total += v;
	}
	return total;
}

func plus(a int , b int) int  {
	return a + b;
}

func plusplus(a int , b int , c int) int {
	return a + b + c;
}

// Multiple return values 

func vals() (int , int) {
	return 3,7;
}


// closures


func intSeq () func() int {
	i := 0;
	return func() int {
		i++;
		return i
	}
}



func main() {

	res := plus(2,3);
	fmt.Println("The result of 2 and 3 is : " , res);
	fmt.Println(plusplus(20,30,40));
	a,b := vals();
	fmt.Println(a);
	fmt.Println(b);

	// LEAVE blank if we dont use it 

	_,c := vals();
	fmt.Println(c);

	fmt.Println(sum(1,2,3,4));
	fmt.Println(sum(100,200,200));

	// next int seq

	nextSeq := intSeq();
	fmt.Println("next number : " , nextSeq());
	fmt.Println("next number : " , nextSeq());
	fmt.Println("next number : " , nextSeq());
	fmt.Println("next number : " , nextSeq());

	nextSeq = intSeq();
	fmt.Println("initialized to zero");
	fmt.Println("Next number is  : " ,  nextSeq());



}



