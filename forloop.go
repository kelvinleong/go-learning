package main

import "fmt"

func main() {
	i := 1
	for i <= 3 {
		fmt.Println(i);
		i += 1
	}

	// while (true)
	for {
		fmt.Println("loop")
		break
	}

	// traditional for loop syntax
	for n := 0; n <= 5; n++ {
		if n % 2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
