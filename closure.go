package main

import "fmt"

func nextSeq() func() int{
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	next := nextSeq()

	fmt.Println(next())
	fmt.Println(next())
}
