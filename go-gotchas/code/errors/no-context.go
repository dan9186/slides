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
		return nil, err
	}

	age, err := getAge()
	if err != nil {
		return nil, err
	}

	return &user{name, age}, nil
}

func getName() (string, error) {
	return "", fmt.Errorf("lookup error")
}

func getAge() (int, error) {
	return 0, fmt.Errorf("lookup error")
}
