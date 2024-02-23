package main

import (
	"errors"
	"fmt"
)

type argError struct {
	name string
}

func (ae argError) Error() string {
	return fmt.Sprintf("%v, You cannot divide by 0 ", ae.name)
}

func (ae argError) newError() string {
	err := errors.New("cannot divide by zero")
	if err != nil {
		return err.Error()
	}
	return " "
}

func main() {
	ae := argError{}
	ae.name = "Aniket"
	s := error(ae)
	fmt.Println(s)
	s2 := ae.newError()
	fmt.Println(s2)

}
