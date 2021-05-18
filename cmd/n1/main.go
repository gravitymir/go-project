package main

import "fmt"

func main() {

	fmt.Printf("%e", run())
}

func run() error {
	fmt.Println("func \"run\" is work!!!")
	return nil
}
