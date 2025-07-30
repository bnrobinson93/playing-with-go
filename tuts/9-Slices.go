package main

import (
	"fmt"
	"slices"
)

func main() {
	var s []string // Just like an array but with no size specified. A slice has a length (len) and a capacity (cap)
	fmt.Println("uninit:", s, s == nil, len(s) == 0)

	s = make([]string, 3) // Creates a non-zero length slice with three items, each with a zero value
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("append:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	l := s[2:5] // slices only values at index 2-5 into a new slice
	fmt.Println("sl1:", l)

	l = s[:5] // up-to-and-including index 5
	fmt.Println("sl2:", l)

	l = s[2:] // index 2 to the end
	fmt.Println("sl3:", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dc1:", t)

	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	}

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
