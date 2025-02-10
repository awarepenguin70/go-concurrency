package main

import (
	"fmt"
	"time"
)

func someFunc(num string) {
	fmt.Println(num)
}

func main() {
	//channels
	myChannel := make(chan string)
	anotherChannel := make(chan string)

	go func() { myChannel <- "data" }()
	go func() { anotherChannel <- "cow" }()

	//select
	select {
	case msgFromMyChannel := <-myChannel:
		fmt.Println(msgFromMyChannel)
	case msgFromAnotherChannel := <-anotherChannel:
		fmt.Println(msgFromAnotherChannel)

	}

	//asynchronous
	go someFunc("1")
	go someFunc("2")
	go someFunc("3")

	time.Sleep(2 * time.Second)

	fmt.Println("hi")
}
