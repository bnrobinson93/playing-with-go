package main

import "fmt"

func vals() (int, int) { // multiple return values
	return 3, 7
}

func main() {
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals() // to ignore one, use the blank identifier
	fmt.Println(c)
}
