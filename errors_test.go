package gotalk_test

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func ExampleErrors_wrapping() {
	i := errors.New("this is an error")
	j := fmt.Errorf("this error wraps error %w", i)

	fmt.Println(errors.Is(j, i))

	// Output:
	//
	// true
}

func ExampleErrors_returnErrors() {
	i, err := strconv.Atoi("42")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	fmt.Println(i)

	// Output:
	//
	// 42
}

func ExampleErrors_deferPanicRecover() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("calm")
		}
	}()
	panic("panic")

	fmt.Println("I don't print")

	// Output:
	//
	// calm
}
