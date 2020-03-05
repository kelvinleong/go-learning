package main

import (
	"fmt"
	"math"
)

type person struct {
	name	string
	age 	int
	gender  string
}

// method definition
func (p *person) describe() {
	fmt.Printf("name is %v and age is %v \n", p.name, p.age)
}

func (p *person) setAge(age int) {
	p.age = age
}

func (p *person) setName(name string) {
	p.name = name
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}

func testPerson() {
	pp := &person{name: "Bob", age: 42, gender: "Male"}
	pp.describe()

	pp.setAge(25)
	pp.setName("Alice")
	pp.describe()
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

	var ap *int
	a := 1
	ap = &a

	fmt.Println(*ap)

	testPerson()
}
