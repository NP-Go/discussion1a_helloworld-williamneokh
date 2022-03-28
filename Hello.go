package main

import (
	"fmt"
)

func HelloWorld() string {
	//Insert your code here
	messageHelloWorld := "Hello World from Go !!"
	//Do not remove this line here
	return messageHelloWorld
}

func main() {
	message := HelloWorld()
	fmt.Println(message)

}
