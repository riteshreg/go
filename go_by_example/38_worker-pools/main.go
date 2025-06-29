package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for ch := range jobs {
		fmt.Println("worker", id, "started  job", ch)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", ch)

		results <- ch * 2
	}
}

func main() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}

	close(jobs)

	for a := 1; a <= numJobs; a++ {
		<-results
	}

}
