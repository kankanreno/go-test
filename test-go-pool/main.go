package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func worker(id int, input <-chan int, output chan<- int) {
	// Decreasing internal counter for wait-group as soon as goroutine finishes
	defer wg.Done()

	// Consumer: Process items from the input channel and send output to output channel
	for j := range input {
		fmt.Println("workder", id, "start job", j)
		time.Sleep(time.Second)
		fmt.Println("workder", id, "finished job", j)
		output <- j * 2
	}
}

func main() {
	input := make(chan int, 100)
	output := make(chan int, 100)

	// Increment waitgroup counter and create go routines
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, input, output)
	}

	// Producer: load up input channel with input
	for i := 1; i <= 5; i++ {
		input <- i
	}

	// Close input channel since no more input are being sent to input channel
	//close(input)

	// Wait for all goroutines to finish processing
	wg.Wait()

	// Close output channel since all workers have finished processing
	close(output)

	// Read from output channel
	for result := range output {
		fmt.Println(result)
	}
}