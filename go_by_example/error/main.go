package main

import (
	"errors"
	"fmt"
)

func f(arg int) (int, error) {
	if arg == 10 {
		return -1, errors.New("10 will not work")
	}
	return arg + 1, nil
}

func main() {

	if val, err := f(10); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("hi", val)
	}

	for _, item := range []int{5, 7, 9, 10, 12} {
		if val, err := f(item); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("hi from loop", val)
		}
	}

}
