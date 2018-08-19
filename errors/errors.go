package main

import (
	"fmt"
	"strconv"
)

func main() {
	i, err := strconv.Atoi("42")
	if err != nil {
		fmt.Printf("could not convert %v to int type", err)
		return
	}
	fmt.Printf("Converted to integer: %v \n", i)
}
