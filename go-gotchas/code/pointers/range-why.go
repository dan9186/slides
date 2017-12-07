package main

import (
	"fmt"
)

type User struct {
	Name string
	ID   int
}

func main() {
	u1 := User{Name: "Amelia", ID: 1}
	u2 := User{Name: "Allison", ID: 2}

	fmt.Printf("User 1 Addr: %p\n", &u1)
	fmt.Printf("User 2 Addr: %p\n\n", &u2)

	us := []User{u1, u2}

	for _, u := range us {
		fmt.Printf("Name: %s ID: %d\n", u.Name, u.ID)
		fmt.Printf("Addr: %p\n", &u)
	}
}
