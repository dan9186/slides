package main

import (
	"fmt"
	"os"
)

func main() {
	defer func() {
		fmt.Println("Executed the defer") // The anonymous func will not be executed
	}()
	fmt.Println("Defer added to the stack")

	os.Exit(1)
}
