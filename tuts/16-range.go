package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 3, 4} // create an initialize a slice of ints; arrays work the same
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"} // same for a map
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println("Key:", k)
	}

	for i, c := range "go" { // this will iterate over the byte array (runes)
		fmt.Println(i, c) // i is the starting point, c is the rune
		fmt.Printf("%d %#U\n", i, c)
	}
}
