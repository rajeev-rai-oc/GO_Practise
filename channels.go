package main

import "fmt"

func main(){

	channel := make(chan string)

	go func(){
		channel<-"Hello World"
	}()

	fmt.Println(<-channel)

}