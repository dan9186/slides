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

	for idx := range us {
		coolUsers = append(coolUsers, &us[idx])
	}

	for _, u := range coolUsers {
		fmt.Println(u)
	}
}
