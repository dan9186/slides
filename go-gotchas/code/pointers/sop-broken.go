package main

import (
	"fmt"
)

type User struct {
	Name string
	ID   int
}

func main() {
	us := []User{
		User{Name: "Amelia", ID: 1},
		User{Name: "Allison", ID: 2},
	}

	coolUsers := []*User{}

	for _, u := range us {
		coolUsers = append(coolUsers, &u)
	}

	for _, u := range coolUsers {
		fmt.Println(u)
	}
}
