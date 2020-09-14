package main

import (
	"errors"
	"fmt"
)

// go communicate errors with an explicit separate return value
// not like JAVA exceptions
// program does not stop when meet error
// type error is a built-in interface

func f1(arg int) (int, error) {
	// return two values, one int type and one error type
	if arg == 42 {
		return -1, errors.New("Can't work with 42")
	}
	// a nil indicated there is no error
	return arg, nil
}

func main() {
	fmt.Println("Hello Errors in Go")

	// range return two values, index and element
	for ii, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println(ii, "f1 failed:", e)
		} else {
			fmt.Println(ii, "f2 worked:", r)
		}
	}
}
