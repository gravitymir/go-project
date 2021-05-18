package main

import (
	"errors"
	"fmt"
)

func main() {

	fmt.Printf("%v", run())
}

func run() error {
	return errors.New("func \"run\" is work!!!\n")
}
