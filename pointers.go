package main

import "fmt"

func changePointer(ptr *int) {
	*ptr = 99;
}

func main()  {
	i := 0;
	changePointer(&i)
	fmt.Println(i)
}