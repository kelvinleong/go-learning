package main

import (
	"fmt"
	"time"
)

func f (from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, "::", i)
	}
}

func main() {
	f("function")

	go f("routine")

	go func(msg string) {
		fmt.Println(msg)
	} ("routine function")

	time.Sleep(time.Second);
	fmt.Println("Done")
}