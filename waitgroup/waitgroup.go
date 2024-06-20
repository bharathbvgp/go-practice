package main

import (
	"fmt"
	"sync"
	"time"
)

// workers

func worker(id int) {
	fmt.Println("started working ... " , id);
	time.Sleep(time.Second);
	fmt.Println("completed .... " , id);
}

func main() {
	var wg sync.WaitGroup;
	for i := 0 ; i < 5 ; i++ {
		wg.Add(1);
		go func() {
			defer wg.Done();
			worker(i);
		} ()
	}
	wg.Wait();
}