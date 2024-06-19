package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0 ; i < 3 ; i++ {
		fmt.Println(from , ":" , i);
	}
}

func say (s string) {
	for i := 0 ; i < 5 ; i ++ {
		fmt.Println(s);
		time.Sleep(100 * time.Millisecond);
	}
}


func main() {
	// synchronous call
	f("bharath")

	go f("goroutine")
	go func (msg string) {
		fmt.Println(msg);
	}("going")
	time.Sleep(time.Second);
	fmt.Println("done");

	go say("hello from goroutine");
	say("hello synchronous");

}