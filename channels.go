package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, j)
		time.Sleep(time.Second) // Simulate some work
		fmt.Printf("Worker %d finished job %d\n", id, j)
		results <- j * 2
	}
}

func main() {
	numJobs := 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// Start three worker goroutines
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// Send jobs to the worker goroutines
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// Collect the results
	for a := 1; a <= numJobs; a++ {
		result := <-results
		fmt.Println("Result:", result)
	}
}