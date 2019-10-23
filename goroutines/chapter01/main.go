package main

import (
	"fmt"
	"time"
)

func main() {

	go sayHello()

	time.Sleep(2 * time.Second)
}

func sayHello(){
	fmt.Printf("Hello!!")
}
