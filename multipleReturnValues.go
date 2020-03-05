package main

import "fmt"

func vals() (int, int, int) {
	return 2, 4, 6
}

func main() {
	a, b, c := vals()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	_, e, _ := vals()
	fmt.Println(e)
}
