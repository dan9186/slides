package main

import (
	"fmt"
	"log"
)

type user struct {
	Name string
	Age  int
}

func main() {
	u, err := getUser()
	if err != nil {
		log.Fatalf("Failed to get user: %v", err.Error())
	}

	log.Printf("User: %v", u)
}

func getUser() (*user, error) {
	name, err := getName()
	if err != nil {
		return nil, fmt.Errorf("failed to get name: %v", err.Error())
	}

	age, err := getAge()
	if err != nil {
		return nil, fmt.Errorf("failed to get age: %v", err.Error())
	}

	return &user{name, age}, nil
}

func getName() (string, error) {
	return "Rob", nil
}

func getAge() (int, error) {
	return 0, fmt.Errorf("lookup error")
}
