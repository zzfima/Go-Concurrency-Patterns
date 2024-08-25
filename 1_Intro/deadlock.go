package main

import "fmt"

func deadlock() {
	ch := make(chan int)

	//deadlock situation - because no goroutine ready to receive from it
	ch <- 5

	fmt.Println(<-ch)
}
