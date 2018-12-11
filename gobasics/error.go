package main

import (
	"errors"
)

func MyDiv(a, b int) (result int, err error) {
	if b == 0 {
		err = errors.New("除数为0")
	} else {
		result = a / b
	}
	return
}

func test() {
	defer func() {
		if err := recover(); err != nil {
			return err

		}
	}()
	panic("panicccc")
}

func main() {
	test()

}
