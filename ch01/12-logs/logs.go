package main

import (
	"log"
	"os"
)

func main() {

	args := os.Args
	if len(args) == 1 {
		// In most cases, this custom error message can be Not enough arguments, Cannot access file, or similar.
		// log.Fatal() calls log.Print() and then os.Exit(1),
		log.Fatal("User should provide a command line argument.")
	}
	// We can use log.Panic() for events that are unexpected, such as not being able to find a file that was previously accessed or not having enough disk space
	log.Panic("Something unusual happened.")
}
