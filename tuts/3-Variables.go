package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
	a := "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	d := true
	fmt.Println(d)

	var e int // NOTE: zero value: 0, false, ""
	fmt.Println(e)

	// Initialize at same time as assignment
	// NOTE: Only available within a function
	f := true
	fmt.Println(f)

	// const
	fmt.Println(s)

	const n = 500000000
	const value = 3e20 / n
	fmt.Println(value)
	fmt.Println(int64(value))
	fmt.Println(math.Sin(n))
}
