package main

import (
	"fmt"
)

/**
goroutine
 */
func main() {
	c := make(chan bool)
	go func() {
		fmt.Println("go go go!")
		c <- true
	}()
	<- c

}
