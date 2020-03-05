package main

import "fmt"

type argError struct {
	arg int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprint("%d - %s", e.arg, e.prob)
}

func main()  {
	error := argError{arg: 404, prob: "NOT FOUND"}
	print(&error)
}
