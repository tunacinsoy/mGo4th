// This go program is aimed to generate custom log entries using different flags.
package main

import (
	"fmt"
	"log"
	"os"
	"path"
)

func main() {
	LOGFILE := path.Join(os.TempDir(), "mGo.log")

	fmt.Println(LOGFILE)
	f, err := os.OpenFile(LOGFILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644) // 110 100 100 -> rw-r--r--

	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	LstdFlags := log.Ldate | log.Lshortfile // When we OR them, they appear in this order (2024/09/20 customLogLineNumber.go:22)
	iLog := log.New(f, "LNum ", LstdFlags)
	iLog.Println("Mastering Go, 4th Edition!")
	iLog.SetFlags(log.Lshortfile | log.LstdFlags) // Actually, log.LstdFlags is equal to log.Ldate | log.Ltime
	iLog.Println("Another log entry!")

	// However, as we can see from the output, there is no control over the order they appear.
	// LNum 2024/09/20 customLogLineNumber.go:22: Mastering Go, 4th Edition!
	// LNum 2024/09/20 21:00:15 customLogLineNumber.go:24: Another log entry!

}
