package main

import (
	"fmt"
	"time"
)

func worker(ch chan int) {
    time.Sleep(2 * time.Second)
    ch <- 42
}



func main() {

	// i need to create two channels

	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)

	go func() {
		time.Sleep(1 * time.Second);
		ch1 <- "First";
	}()

	go func() {
		time.Sleep(2 * time.Second);
		ch2 <- "second";
	}()

	timeout := time.After(3 * time.Second);

	for i := 0 ; i < 2 ; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("msg1" , msg1);
		case msg2 := <-ch2:
			fmt.Println("msg2" , msg2);

		case <- timeout:
			fmt.Println("Timeout");
			return 
		
		}

	}


	ch5 := make(chan string);
	go func () {
		time.Sleep(2 * time.Second);
		ch5 <- "43"; // blocking operation

		// here the point is that this line of code will never execute because the above line is
		// a blocking operation once it completes then this line of code will be executed
		// the send and receive operation on a channel can be blocking if there is no sender for th 
		// e receving operation and there is no receiver for the sending operation
		fmt.Println("this is running after blocking call ....");
	} ()

	fmt.Println("waiting for message .... ");
	// fmt.Println(<-ch5);  // BLOCKING Operation
	fmt.Println("WAIT COMPLETED");


	// Non blocing channels can be acheived using select

	ch6 := make(chan string , 1);

	select {
	case ch6 <- "bharath":
		time.Sleep(2 * time.Second)
		fmt.Println("CH6 received my data");
	default:
		fmt.Println("timeout");
	}
	fmt.Println("After select")


	ch := make(chan int)
    go worker(ch)
    fmt.Println("Waiting for value...")
    val := <-ch
    fmt.Println("Received:", val)
	fmt.Println("END")

}