package main

import (
	"fmt"
)

type User struct {
	Name string
	ID   int
}

func main() {
	us := []*User{
		&User{Name: "Amelia", ID: 1},
		&User{Name: "Allison", ID: 2},
	}

	for _, u := range us {
		fmt.Println(*u)
	}

	fmt.Println()

	for _, u := range us {
		fmt.Println(u)
	}
}
