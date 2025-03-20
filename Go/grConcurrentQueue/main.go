package main

import (
	"fmt"
	"log"
	"math/rand/v2"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	batchSize   = 200
	concurrency = 20 // Number of concurrent workers
)

func main() {
	// Channel to send batches of data to workers
	dataChan := make(chan []int)

	// Semaphore to limit concurrency (acts like RabbitMQ QoS)
	semaphore := make(chan struct{}, concurrency)

	// Start a goroutine to fetch data in batches and send to workers
	go func() {
		for i := range 100 {
			log.Println("Loop Number : ", i)

			batch := []int{}
			for _ = range batchSize {
				batch = append(batch, rand.IntN(5))
			}

			dataChan <- batch // Send batch to workers
		}
		close(dataChan) // Close the channel when done fetching data
	}()

	// Process data using workers with concurrency control
	for batch := range dataChan {
		// Start a worker goroutine
		for _, row := range batch {
			// Acquire a semaphore slot (blocks if concurrency limit is reached)
			semaphore <- struct{}{}

			go func(row int) {
				defer func() {
					<-semaphore // Release the semaphore slot when done
				}()

				processRow(row)
			}(row)
		}
	}

	// Wait for all workers to finish
	for len(semaphore) > 0 {
		time.Sleep(100 * time.Millisecond) // Poll to check if all workers are done
	}

	fmt.Println("All data processed!")

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT)

	sig := <-sigChan
	log.Println("Signal Received : ", sig)
	log.Println("Closing Worker")
}

func processRow(row int) {
	// Simulate processing (e.g., some execution based on the data)
	// fmt.Printf("Processing row: %d\n", row)

	time.Sleep(time.Duration(row) * 10 * time.Millisecond) // Simulate work
}
