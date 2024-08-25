package main

import (
	"fmt"
	"time"
)

func buffChannel() {
	chBuf := make(chan int, 2)

	go func() {
		chBuf <- 1
		chBuf <- 2
		chBuf <- 3
		time.Sleep(time.Second * 3)
		chBuf <- 4
	}()
	fmt.Println(<-chBuf)
	fmt.Println(<-chBuf)
	fmt.Println(<-chBuf)
	fmt.Println(<-chBuf)
}
