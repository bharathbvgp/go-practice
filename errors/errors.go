package main

import (
	"errors"
	"fmt"
)

func f(arg int) (int , error){
	if arg == 42 {
		return -1 , errors.New("cant work with 42");
	} 
	return arg + 3 , nil;
}

var ErrorOutOftea = fmt.Errorf("no more tea available")
var ErrorPower = fmt.Errorf("can't boil water")

func makeTea(arg int) error {
	if arg == 2 {
		return ErrorOutOftea;
	} else if arg == 4 {
		return fmt.Errorf("making tea : %w" , ErrorPower);
	}
	return nil
}


// custom errors

type argError struct {
	arg int
	message string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s" , e.arg , e.message);
}

func k(arg int) (int , error) {
	if arg == 42 {
		return -1 , &argError{arg , "Cant work with 42"};
	}
	return arg + 3 , nil
}


func main() {

	_ , err := k(42);
	var ae *argError 
	if errors.As(err , &ae) {
		fmt.Println(ae.arg);
		fmt.Println(ae.message);
	} else {
		fmt.Println("Argument did not match argError");
	}

	for _,i := range []int {7 , 42} {
		if r , e := f(i) ; e != nil {
			fmt.Println("f failed" , e);
		} else {
			fmt.Println("f worked: " , r);
		}
	}

	for i := range 5 {
		if err := makeTea(i); err != nil {
			if errors.Is(err , ErrorOutOftea) {
				fmt.Println("We should buy new Tea!");	
			} else if errors.Is(err , ErrorPower) {
				fmt.Println("Now it is dark");
			} else {
				fmt.Printf("Unknow error: %s\n" , err);
			}
			continue
		} 
		fmt.Println("Tea is ready!!!!");
	}



}