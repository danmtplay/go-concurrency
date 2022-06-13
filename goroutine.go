// goroutine.go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	start := time.Now()

	decoupled()

	elapsed := time.Since(start)

	fmt.Println(elapsed)
}

func decoupled() {
	c := fanIn(foo("Alice"), foo("Bob"))
	for i := 0; i < 20; i++ {
		fmt.Println(<-c)
	}
}

func foo(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1337)) * time.Millisecond)
		}
	}()
	return c
}

func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case s := <-input1:
				c <- s
			case s := <-input2:
				c <- s
			}
		}
	}()
	return c
}
