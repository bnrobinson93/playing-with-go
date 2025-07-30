package main

import (
	"errors"
	"fmt"
)

type argError struct {
	arg     int
	message string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.message)
}

func f(arg int) (int, error) {
	if arg == 42 {
		// return -1, errors.New("Can't work with 42") // or a custom error:
		return -1, &argError{arg, "can't work with it"}
	}

	return arg + 3, nil
}

var (
	ErrOutOfTea = fmt.Errorf("no more tea available")
	ErrPower    = fmt.Errorf("can't boil water")
)

func makeTea(arg int) error {
	if arg == 2 {
		return ErrOutOfTea
	} else if arg == 4 {
		return fmt.Errorf("making tea: %w", ErrPower)
	}
	return nil
}

func main() {
	for _, i := range []int{7, 42} {
		if r, e := f(i); e != nil { // this is a commont pattern: run and check if error
			fmt.Println("f failed:", e)
		} else {
			fmt.Println("f worked:", r)
		}
	}

	_, err := f(42)
	var ae *argError
	if errors.As(err, &ae) { // maps an error to custom type. Works like errors.Is but for custom types, since it may not have the same sig
		fmt.Println(ae.arg)
		fmt.Println(ae.message)
	} else {
		fmt.Println("err doesn't match argError")
	}

	for i := range 5 {
		if err := makeTea(i); err != nil {
			if errors.Is(err, ErrOutOfTea) {
				fmt.Println("We should buy new tea!")
			} else if errors.Is(err, ErrPower) {
				fmt.Println("Now it is dark")
			} else {
				fmt.Printf("unknown error: %s\n", err)
			}
		}
	}
}
