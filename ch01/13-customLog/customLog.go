// The aim of this program is to get familiar with logging into custom files.
package main

import (
	"fmt"
	"log"
	"os"
	"path"
)

func main() {
	LOGFILE := path.Join(os.TempDir(), "mastering-go-ch01-13.log")
	fmt.Println(LOGFILE) // /tmp/mastering-go-ch01-13.log

	// If the file doesn't exist, create it, or append to the file
	f, err := os.OpenFile(LOGFILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644) // 110 100 100 -> rw-r--r--

	if err != nil {
		log.Fatal(err)
	}

	// defer keyword tells Go to execute the statement just before the current function returns.
	//This means that f.Close() is going to be executed just before main() returns.
	defer f.Close()

	// First input argument is the destination to which log data will be written;
	// Second input argument is the prefix before each generated log line;
	// Third input argument is the flag, which are the initial values for standard logger.
	iLog := log.New(f, "iLog ", log.LstdFlags)
	iLog.Println("Hello there!")
	iLog.Println("Mastering Go 4th Edition!")

	// if we run command `cat /tmp/mastering-go-ch01-13.log` the output would be:
	// iLog 2024/09/20 19:12:31 Hello there!
	// iLog 2024/09/20 19:12:31 Mastering Go 4th Edition!

}
