package main

import "fmt"

func Add(a, b int) int {
	return a + b
}

func main() {
	if sum := Add(1, 2); sum != 3 {
		panic(fmt.Sprintf("sum expected to be 3; got %d", sum))
	}

	fmt.Println("Well done!")
}
