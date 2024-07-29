package main

import (
	"errors"
	"fmt"
)

func f(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't work with 42")
	}

	return arg + 3, nil
}

var ErrOutOfTea = fmt.Errorf("no more tea available")
var ErrPower = fmt.Errorf("can't boil water")

func makeTea(arg int) error {
	if arg == 2 {
		return ErrOutOfTea
	} else if arg == 4 {

		return fmt.Errorf("making tea: %w", ErrPower)
	}
	return nil
}


//custom error

type customError struct{
	arg int
	str string
}

func (e *customError) Error() string{
	return fmt.Sprintf("%d --> %s", e.arg, e.str)
}

func fun(arg int) (int, error) {
	if arg == 42{
		return -1, &customError{arg,"Custom: can't work with 42"}
	}

	return arg + 3, nil
}

func main() {
	for _, i := range []int{7, 42} {

		if r, e := f(i); e != nil {
			fmt.Println("f failed:", e)
		} else {
			fmt.Println("f worked:", r)
		}
	}

	for i:=1;i<=5;i++{
		if e:=makeTea(i); e!=nil{
			if errors.Is(e,ErrOutOfTea ){
				fmt.Println("New tea Required:")
			} else if errors.Is(e, ErrPower){
				fmt.Println("No Power")
			} else{
				fmt.Println("Unknown Error: %s ",e)
			}
		}
	}

	//custom error

	_, err := fun(43)
	var ae *customError
	if errors.As(err, &ae) {
	// `errors.As` is a more advanced version of `errors.Is`.
	// It checks that a given error (or any error in its chain)
	// matches a specific error type and converts to a value
	// of that type, returning `true`. If there's no match, it
	// returns `false`.
		fmt.Println(ae.arg)
		fmt.Println(ae.str)
	} else {
		fmt.Println("err doesn't match argError")
	}
}