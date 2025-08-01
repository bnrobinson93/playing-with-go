package main

import (
	"fmt"
)

func zeroval(ival int) { // this will have it's own copy of the value
	ival = 0
}

func zeroptr(iptr *int) { // this will reference the real value
	*iptr = 0 // dereference to update the value
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i) // this won't do anything to "i" since it's passed as a copy
	fmt.Println("zeroval:", i)

	zeroptr(&i) // this will update "i" since it's passed by reference
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)
}
