package main

import (
	"fmt"
	ratelimiter "github.com/dallaer/ratelimpkg"
	"time"
)

func say(str string) {
	fmt.Println(str)
	time.Sleep(2 * time.Second)
}
func square(x, y int) int {
	return x * y
}
func main() {
	ratelimiter.Initialization(3, 24)
	ch := make(chan func(), 3)
	ch <- func() { say("smth...") }
	ch <- func() { example1() } // compares  content of two URLby byte slice, and signals when they start to differ.
	// (one of the urls does not exist yet, but should appear soon, the second is a redirect link)
	ch <- func() { square(5, 4) }
	ratelimiter.Ratelimiter(ch)

}
