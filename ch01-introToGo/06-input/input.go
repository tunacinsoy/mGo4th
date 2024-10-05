// The aim of this go code to get familiar with the fmt.Scanln() function to get user input from the console.
package main

import "fmt"

func main() {
	// Get user input
	fmt.Printf("Please give me your name: ")
	var name string

	fmt.Scanln(&name)
	fmt.Printf("Your name is: %v\n", name)

}
