// The aim of this program to understand how versioning can be done using git commit hashes.

/*
	 User should run:
	 # To retrieve the latest commit hash
		export VERSION=$(git rev-list -1 HEAD)
		# Following command is used to build a Go program with custom values set for variables at compile time
		go build -ldflags "-X main.VERSION=$VERSION" gitVersion.go

		In one line:
		export VERSION=$(git rev-list -1 HEAD) && go build -ldflags "-X main.VERSION=$VERSION" gitVersion.go                                                                                                                  │

		After that, program can be run using:
		./gitVersion version
		And the output will be the version hash for the new program.
*/
package main

import (
	"fmt"
	"os"
)

var VERSION string

func main() {
	if len(os.Args) == 2 {
		if os.Args[1] == "version" {
			fmt.Println("Version: ", VERSION)
		}
	}
}
