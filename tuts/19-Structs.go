package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p // NOTE: this will only be cleaned up when there are no references to it. It gets garbage collected
}

func main() {
	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name: "Alice", age: 30})
	fmt.Println(person{name: "Fred"})
	fmt.Println(&person{name: "Ann", age: 40})
	fmt.Println(newPerson("Jon"))

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age) // this is automatically dereferenced

	sp.age = 51
	fmt.Println(sp.age)
	fmt.Println(s.age) // Since it was reference, this is updated too

	dog := struct { // anonymous struct; often used for table-driven tests
		name   string
		isGood bool
	}{
		"Dr. Franklin",
		true,
	}
	fmt.Println(dog)
}
