package main

import (
	"fmt"
)

func ProcessData() {
	ch := make(chan int)
	ch2 := make(chan int)

	go func() {
		for i := range 10 {
			ch <- i
		}
		close(ch)
	}()
	go func() {
		for i := range 10 {
			ch2 <- i*100 + 1
		}
		close(ch2)
	}()

	count := 2
	for count != 0 {
		select {
		case data, ok := <-ch:
			if !ok {
				ch = nil
				count--
				break
			}
			fmt.Printf("Number %d from chanel 1\n", data)
		case data, ok := <-ch2:
			if !ok {
				ch2 = nil
				count--
				break
			}
			fmt.Printf("Number %d from chanel 2\n", data)
		}
	}
}

func main() {
	ProcessData()
}
