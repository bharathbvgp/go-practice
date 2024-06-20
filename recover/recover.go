package main

import "fmt"

func mayPanic() {
	panic("Custom panic!!!!")
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered from the error" , r);
		}
	}()
	mayPanic()
	fmt.Println("After mayPanic()")
}