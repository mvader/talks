package main

import "fmt"

func main() {
	a := 5
	switch a {
	case 5:
		fmt.Println("a is 5")
	case 6, 7, 8:
		fmt.Println("a is 6, 7 or 8")
	default:
		fmt.Println("i don't know what is a")
	}
}
