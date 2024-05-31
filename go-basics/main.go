package main

import "fmt"

func main() {
	fmt.Println("Hello World!")
	sl := []int{1, 2, 3}
	if len(sl) > 2 {
		fmt.Println("unreachable code")
	}
}
