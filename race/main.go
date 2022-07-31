package main

import (
	"fmt"
)

func main() {
	a := 0
	// a starting with zero we use add and subs functions to manipulate a variable
	// Since we are adding and substracting 200 everytime at the end we should have zero
	// But since we are using goroutines and accessing the same variable sometimes we end up with different values
	for i := 0; i <= 500; i++ {
		go add(&a, 200)
		go subs(&a, 200)
	}
	fmt.Println(a)
}

func add(a *int, amount int) {
	*a += amount
}

func subs(a *int, amount int) {
	*a -= amount
}
