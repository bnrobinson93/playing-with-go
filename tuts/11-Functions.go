package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

// If multiple args are the same time, just specify on the last one
func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {
	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	res = plusPlus(1, 2, res)
	fmt.Println("1+2+3 =", res)
}
