package main

import (
	"fmt"
)

type User struct {
	Name string
	ID   int
}

func main() {
	u := &User{"Jimbo", 1}

	if u, err := nilOutUser(); err != nil {
		fmt.Printf("Couldn't get user: %v\n", err.Error())
		fmt.Printf("User should be nil here: %v\n", u)
		return
	}

	fmt.Printf("User: %v", u)
}

func nilOutUser() (*User, error) {
	return nil, nil
}
