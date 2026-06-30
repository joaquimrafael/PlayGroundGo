package main

import (
	"fmt"
	"sync"
)

func ProcessData() {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := range 10 {
			ch <- i
		}
	}()
	go func() {
		defer wg.Done()
		for i := range 10 {
			ch <- i*100 + 1
		}
	}()

	go func() {
		wg.Wait()
		close(ch)
	}()

	var wg2 sync.WaitGroup
	wg2.Go(func() {
		for v := range ch {
			fmt.Println(v)
		}
	})
	wg2.Wait()
}

func main() {
	ProcessData()
}
