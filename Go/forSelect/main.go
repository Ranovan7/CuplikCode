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

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT)

	go func(){
		for {
			select {
			case <- ctx.Done():
				log.Println("Cancel received, Shutting Down Program...")
				return
			default:
				for {
					log.Println("Running Program...")
					time.Sleep(1 * time.Second)
				}
			}
		}
	}()

	<-sigChan
	cancel()
	time.Sleep(3 * time.Second)
}


