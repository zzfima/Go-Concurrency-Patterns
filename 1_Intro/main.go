package main

import (
	"fmt"
	"time"
)

func main() {
	var ch chan int = make(chan int)
	go listener(ch)
	ch <- 5
	time.Sleep(time.Second * 1)
}

func listener(ch chan int) {
	fmt.Println(<-ch)
}
