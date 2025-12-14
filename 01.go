package main

import "fmt"

func main() {
	fmt.Println("Hello and welcome to Go programming")
	fmt.Println("Hello and welcome to", "Go Programming")
	message := []byte("Hello and Welcome to Go Programming")
	fmt.Printf("%s\n", message)
	fmt.Printf("%s !! I love %s\n", message, "Go Programming")
}
