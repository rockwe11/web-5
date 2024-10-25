package main

import "fmt"

// реализовать calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int
func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
	output := make(chan int)
	go func() {
		defer close(output)
		select {
		case a := <-firstChan:
			output <- a * a
		case b := <-secondChan:
			output <- b * 3
		case <-stopChan:
			return
		}

	}()
	return output
}

func main() {
	ch1, ch2 := make(chan int), make(chan int)
	stop := make(chan struct{})

	r := calculator(ch1, ch2, stop)
	ch1 <- 4
	// ch2 <- 3
	close(stop)
	fmt.Println(<-r)
}
