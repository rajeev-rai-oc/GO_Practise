// A _goroutine_ is a lightweight thread of execution.
package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {


	// To invoke this function in a goroutine, use
	// `go f(s)`. This new goroutine will execute
	// concurrently with the calling one.
	go f("goroutine")

	// You can also start a goroutine for an anonymous
	// function call.
	go func(msg string) {
		fmt.Println(msg)
	}("going 1:")

	// func(){
	// 	fmt.Println("Running:")
	// }()

	go func(msg string) {
		fmt.Println(msg)
	}("going 2:")


	f("direct")

	time.Sleep(time.Second)
	fmt.Println("done")
}
