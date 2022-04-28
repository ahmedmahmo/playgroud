package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		time.Sleep(time.Second)
		fmt.Println(from, ":", i)
	}
}

func main() {

	f("Wait for me")

	go f("you can go")

	go func(msg string) {
		fmt.Println(msg)
	}("I will go with you")

	time.Sleep(10 * time.Second)
	fmt.Println("done")
}
