package main

import (
	"log"
	"time"
)

type task struct {
	result int
	worker int
}

func main() {
	number := 30
	numberOfWorker := 5
	jobs := make(chan int, number)
	results := make(chan task, number)
	for i := 1; i <= numberOfWorker; i++ {
		go worker(i, jobs, results)
	}

	for i := 1; i <= number; i++ {
		jobs <- i
	}
	close(jobs)
	for i := 0; i < number; i++ {
		result := <-results
		log.Printf("worker %v - result ::: %v", result.worker, result.result)
	}
}
func fib(n int, id int) int {
	log.Printf("Starting fib ::: %v", id)
	time.Sleep(time.Second)
	//if n <= 1 {
	//	return n
	//}
	//result := fib(n-1, 0) + fib(n-2, 0)
	log.Printf("Finished fib ::: %v", id)
	return n * n
}

/*
*
<-chan int read-only
chan <- int write-only
*/
func worker(id int, jobs <-chan int, results chan<- task) {
	for n := range jobs {
		results <- task{
			result: fib(n, id),
			worker: id,
		}
	}
}
