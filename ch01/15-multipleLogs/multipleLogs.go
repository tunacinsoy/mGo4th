// The aim of this program to show how to write logs to multiple files
package main

import (
	"io"
	"log"
	"os"
)

func main() {
	flag := os.O_APPEND | os.O_CREATE | os.O_WRONLY

	file, err := os.OpenFile("myLog.log", flag, 0644)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	w := io.MultiWriter(file, os.Stderr) // Write both to opened file, and to os.Stderr (on terminal)

	logger := log.New(w, "myApp: ", log.LstdFlags)
	logger.Printf("Book %d\n", os.Getpid()) // os.Getpid() returns the process id of the caller
}
