package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

func main() {
	inpt := [][2]int{
		{234, 3214},
		{421514, 321415},
		{3214, 4356},
		{12, 23},
		{21512125,3212151},
	}	

	var wg sync.WaitGroup
	results := make(chan int, len(inpt))

	for _, ip := range inpt {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(time.Duration(rand.Intn(3) + 1) * time.Second)
			
			log.Printf("-- calculate %v", ip)
			results <- ip[0] + ip[1]
		}()
	}

	wg.Wait()
	close(results)

	log.Println("Done Calculating :")
	for res := range results {
		log.Println("--", res)
	}
}

