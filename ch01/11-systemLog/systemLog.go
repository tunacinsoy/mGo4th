package main

import (
	"fmt"
	"log"
	"log/syslog"
)

func main() {
	// Establishes a new connection to the system log daemon. Each write to the returned writer sends a log message with the given priority.
	// LOG_SYSLOG is one of the logging facilities, which are specific categories for loggging purposes.
	// second input argument is just a tag, which is expressed along with the log.
	sysLog, err := syslog.New(syslog.LOG_SYSLOG, "systemLog.go")

	if err != nil {
		fmt.Println(err)
		return
	} else {
		log.SetOutput(sysLog) // Sets the output destination for logs
		log.Print("Everything is fine")
	}

	// After running `go run systemLog.go`, run `journalctl -xe` to check the log.
	// The `journalctl -xe` command is used to view system logs in real time on Linux systems that use `systemd` as their init system.
	// -x for extra details, -e for jump to end of the file, where recent logs occurred.
}
