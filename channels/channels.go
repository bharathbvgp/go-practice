package main

import "fmt"


func sum (arr []int , c chan int) {
	total := 0
	for i := range len(arr) {
		total += arr[i];
	}
	c <- total
}

func ping(pings chan<- string , msg string) {
	pings <- msg;
}

func pong(pongs chan<- string , pings <- chan string) {
	msg := <- pings
	pongs <- msg
}


func main() {
	c := make(chan int)
	arr := []int {1,2,3,4,5,6}
	n := len(arr)
	go sum(arr[:n/2] , c);
	go sum(arr[n/2:] , c);
	x , y := <-c , <-c;
	fmt.Println(x,y,x+y);


	ch1 := make(chan string)
	go func () {
		ch1 <- "hello"
		ch1 <- "names"
	}();
	// msg := <- ch1;
	// msg2 := <- ch1;
	fmt.Println(<- ch1 , <- ch1);

	// channel buffering 

	// buffer space upto 2
	ch3 := make(chan string , 2);
	ch3 <- "hello"
	ch3 <- "goodies"
	fmt.Println(<-ch3, <-ch3);

	pings := make(chan string , 1);
	pongs := make(chan string , 1);
	ping(pings , "hello frm main");
	pong(pings , pongs);
	fmt.Println(<- pongs);
	


}