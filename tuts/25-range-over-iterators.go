package main

import (
	"fmt"
	"iter"
	"slices"
)

type element[T any] struct {
	next *element[T]
	val  T
}

type List[T any] struct {
	head, tail *element[T]
}

func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

// technically, this is just a generator, returns a function
func (lst *List[T]) All() iter.Seq[T] {
	return func(yield func(T) Bool) {
		for e := lst.head; e != nil; e = e.next {
			if !yield(e.val) {
				break
			}
		}
	}
}

// This will just keep running until yeild failes
func genFib() iter.Seq[int] {
	return func(yield func(int) bool) {
		a, b := 1, 1

		for {
			if !yield(a) {
				return
			}
			a, b = b, a+b
		}
	}
}

func main() {
	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)

	for e := range lst.All() { // this will return an iterator, which can be "ranged"
		fmt.Println(e)
	}

	all := slices.Collect(lst.All()) // stores an iterator as a slice
	fmt.Println("all:", all)

	for n := range genFib() {
		if n >= 10 {
			break // or return
		}
		fmt.Println()
	}
}
