package main

import (
	"fmt"
	"math/rand"
	"errors"
)

func main() {
	s, err := myFunc()
	if err != nil {
		panic(err)
	}
	fmt.Println(s)
}

func myFunc() (string, error) {
	if rand.Float64() > 0.5 {
		return "Success", nil
	}
	return "Fail", errors.New("super serious error")
}