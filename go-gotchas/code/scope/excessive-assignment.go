package main

import (
	"fmt"
)

const (
	adminUserQuery = "get admin user"
	userQuery      = "get regular user"
)

type User struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func main() {
	admin := true

	var u *User
	if admin {
		u, err := DBQuery(adminUserQuery)
		if err != nil {
			fmt.Printf("Err: %v\n", err.Error())
			return
		}

		fmt.Printf("User: %v\n", u)
	}

	fmt.Printf("User: %v", u)
}

func DBQuery(q string) (*User, error) {
	return &User{"Admin", "admin"}, nil
}
