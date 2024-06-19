package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	// looping through jobs channel to handle all jobs pushed into jobs channel
	for job := range jobs {
		fmt.Println("WORKER" , id , "started job" , job)
		time.Sleep(time.Second * 1)
		fmt.Println("worker" , id , "finished job" , job)
		results <- job * 2
	}
}

func main() {
	const numJobs = 5
	jobs := make(chan int , numJobs);
	results := make(chan int , numJobs);
	// spawning three workers to finish five jobs 
	for w := 1 ; w <= 3 ; w++ {
		go worker(w , jobs , results);
	}
	// send the jobs through the channel
	for i := 1 ; i <= 5 ; i++ {
		jobs <- i;
	}
	// eventually jobs will be picked by the workers (3 workers are working to finish jobs)

	close(jobs);
	// printing the results to the stdout

	for i := 1 ; i <= 5 ; i++ {
		// fmt.Println(<-results);
		<-results
	}

}