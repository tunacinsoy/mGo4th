// INFO: Since we are using an external package in this code, this program
// should be placed under ~/go/src folder.
// In my case, it is located under `/home/tuna/go/src/github.com/tunacinsoy/mGo4th/ch06/connectSQLite3`
// After placing it there, the commands `go mod init` and `go mod tidy` should be run.

// The aim of this program is to understand basic operations (create, insert, delete) with sqlite3 database
package main

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var dbname string = "ch06.db"

func insertData(db *sql.DB, dsc string) error {

	// 	time.Now() gets the current time.
	// .Format(time.RFC1123) formats the current time as a string in RFC1123 format, which is a common format for timestamps, like Mon, 02 Jan 2006 15:04:05 MST.
	// cT now holds the current timestamp, which will be one of the values inserted into the database.
	cT := time.Now().Format(time.RFC1123)

	// NULL: Assumes the first column is an AUTO_INCREMENT or PRIMARY KEY field, so NULL tells the database to automatically generate a value for that field.
	// ?: The first placeholder will be replaced with the current timestamp (cT).
	// ?: The second placeholder will be replaced with the dsc (description) value passed to the function.
	stmt, err := db.Prepare("INSERT INTO book VALUES(NULL,?,?);")

	if err != nil {
		fmt.Println("Insert data table:", err)
		return err
	}

	// stmt.Exec(cT, dsc) executes the prepared SQL statement with the values cT (current timestamp) and dsc (description).
	// These values replace the ? placeholders in the query.
	// The function checks for errors during the execution. If an error occurs, it prints the error and returns it.
	_, err = stmt.Exec(cT, dsc)

	if err != nil {
		fmt.Println("Insert data table:", err)
		return err
	}
	return nil

}

func selectData(db *sql.DB, n int) error {
	// 	The db.Query() statement does not require Exec() to get executed. Therefore, we
	// 	replace ? with the actual values in the same statement.
	rows, err := db.Query("SELECT * from book WHERE id > ?", n)
	if err != nil {
		fmt.Println("Select:", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var dt string
		var description string

		// The order of the variables in Scan must match the order of columns returned by the query
		// (SELECT * selects all columns, so the order is based on the table schema).
		err = rows.Scan(&id, &dt, &description)

		if err != nil {
			fmt.Println("Row:", err)
			return err
		}

		// time.Parse() is used to convert the
		// string dt (which is expected to be in RFC1123 format) into a time.Time object.
		date, err := time.Parse(time.RFC1123, dt)

		if err != nil {
			fmt.Println("Date: ", err)
		}
		fmt.Printf("%d, %s, %s\n", id, date, description)
	}
	return nil
}

func updateData(db *sql.DB, idGreaterThan int) error {
	cT := time.Now().Format(time.RFC1123)
	_, err := db.Exec("UPDATE book SET time = ? WHERE id > ?", cT, idGreaterThan)
	if err != nil {
		fmt.Println("Error updating fields:", err)
		return err
	}
	return nil
}

func main() {

	os.Remove(dbname)

	db, err := sql.Open("sqlite3", dbname)

	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}
	defer db.Close()

	const create string = `
	CREATE TABLE IF NOT EXISTS book (
		id INTEGER NOT NULL PRIMARY KEY,
		time TEXT NOT NULL,
		description TEXT
	);
	`
	_, err = db.Exec(create)
	if err != nil {
		fmt.Println("Creating table: ", err)
		return
	}

	for i := 1; i < 11; i++ {
		dsc := "Description " + strconv.Itoa(i)
		err = insertData(db, dsc)
		if err != nil {
			fmt.Println("Insert data:", err)
		}
	}

	err = selectData(db, 5)
	if err != nil {
		fmt.Println("Select:", err)
	}

	time.Sleep(2 * time.Second)

	// update data
	err = updateData(db, 7)

	if err != nil {
		fmt.Println("updateData:", err)
	}

	// cT := time.Now().Format(time.RFC1123)
	// db.Exec("UPDATE book SET time = ? WHERE id > ?", cT, 7)

	// Select multiple rows
	err = selectData(db, 8)
	if err != nil {
		fmt.Println("Select:", err)
	}

	// Delete data
	stmt, _ := db.Prepare("DELETE from book where id = ?")
	_, err = stmt.Exec(8)
	if err != nil {
		fmt.Println("Delete:", err)
	}

	// Select multiple rows
	err = selectData(db, 7)
	if err != nil {
		fmt.Println("Select:", err)
	}

	// Count rows in the table
	query, err := db.Query("SELECT count(*) as count from book")
	if err != nil {
		fmt.Println("Select:", err)
		return
	}
	defer query.Close()

	var countOFRows int = 0

	for query.Next() {
		err = query.Scan(&countOFRows)
	}
	if err != nil {
		fmt.Println("Error while reading count of rows:", err)
	}
	fmt.Println("count:", countOFRows)

}
