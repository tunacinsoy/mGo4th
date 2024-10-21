// INFO: Since we are using an external package in this code, this program
// should be placed under ~/go/src folder.
// In my case, it is located under `/home/tuna/go/src/github.com/tunacinsoy/mGo4th/ch06/testSQLite`
// After placing it there, the commands `go mod init` and `go mod tidy` should be run.

// The aim of this program is to get familiar with the usage of external packages
package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	// Opens a connection to the SQLite3 database file test.db.
	// If the file does not exist, SQLite will create it.
	// This .db file will be created under the current directory.
	db, err := sql.Open("sqlite3", "test.db")

	if err != nil {
		fmt.Println("Error connecting:", err)
	}
	defer db.Close()

	var version string

	// QueryRow executes the query and expects a single row of data.
	// Scan(&version) extracts the result (SQLite version) into the version variable.
	err = db.QueryRow("SELECT SQLITE_VERSION()").Scan(&version)

	if err != nil {
		fmt.Println("Version:", err)
		return
	}

	fmt.Println("SQLite3 version:", version)

	// Remove the created db
	os.Remove("test.db")
}
