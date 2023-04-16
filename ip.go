package main

import (
	"fmt"
	"os"
)

var ch (chan any)

func main() {
	go op()
	<-ch

	os.Exit(0)

}
func op() {
	for i := 0; i < 100; i++ {
		fmt.Println("6", i)
		ch <- 0
	}
	close(ch)

}
