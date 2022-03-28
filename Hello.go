package main

import (
	"fmt"
)

func HelloWorld() string {

	messageHelloWorld := "Hello World from Go !!"

	return messageHelloWorld
}

func main() {
	message := HelloWorld()
	fmt.Println(message)

}
