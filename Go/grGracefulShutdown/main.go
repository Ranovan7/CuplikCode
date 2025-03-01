package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		for {
			select {
			case <-ctx.Done():
				log.Println("Stopping Program")
				return
			default:
				log.Println("Right now is :", time.Now())
				time.Sleep(2 * time.Second)
			}
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT)

	sig := <-sigChan
	log.Println("\nReceived signal :", sig)
	log.Println("Initiating Graceful Shutdown")

	cancel()
	time.Sleep(3 * time.Second)

	log.Println("Shutting Down Program...")
}

