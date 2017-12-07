package main

func main() {
	count := 0
	fmt.Printf("Count: %v\n", count)

	// Works
	count++
	fmt.Printf("Count: %v\n", count)

	// Won't work
	fmt.Printf("Count: %v\n", count++)
}
