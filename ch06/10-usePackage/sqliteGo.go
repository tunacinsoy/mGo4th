// User should run `go mod init` and `go mod tidy` to initialize go.mod and go.sum files.

// This program also assumes that tables `Users` and `Userdata` are already generated, so before running this program, user should run:
// `sqlite3 ch06.db`
// `.read createTables.sql`

// The aim of this program is to test previously created package `github.com/tunacinsoy/sqlite06`
package main

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/tunacinsoy/sqlite06"
)

func random(min int, max int) int {
	return rand.Intn(max-min) + min
}

// min is the lower range of characters interval
// max is the upper range of characters interval
// For instance -> parameters 5, 'A' and 'Z' would return
// random generated string that has length of 5 and all characters are within the ASCII interval of A-Z
func getString(length int, firstChar byte, secondChar byte) string {
	min := int(firstChar)
	max := int(secondChar)
	temp := ""
	var i int = 1
	for {

		myRand := random(min, max)
		newChar := byte(myRand)
		temp = temp + string(newChar)

		if i == length {
			break
		}

		i++
	}

	return temp
}

func main() {
	sqlite06.Filename = "ch06.db"

	random_username := strings.ToLower(getString(5, 'A', 'Z'))

	t := sqlite06.Userdata{
		Username:    random_username,
		Name:        "Tuna",
		Surname:     "Cinsoy",
		Description: "Author",
	}

	// Let's add this user
	fmt.Println("Adding user with username:", random_username)
	id := sqlite06.AddUser(t)

	if id == -1 {
		fmt.Println("An error occurred during addition of user", random_username)
	}

	fmt.Printf("User with username: %v has been added, and its id is: %v\n", random_username, id)

	// Print out all users
	data, _ := sqlite06.ListUsers()

	for _, v := range data {
		fmt.Println(v)
	}

	// Let's delete the user that we previously added
	fmt.Println("Deleting user with username:", random_username)
	err := sqlite06.DeleteUser(id)

	if err != nil {
		fmt.Println("An error occurred during deletion of user with id", id)
	}

	fmt.Printf("User with username: %v has been deleted, and its id was: %v\n", random_username, id)

	// Print out all users
	data, _ = sqlite06.ListUsers()

	for _, v := range data {
		fmt.Println(v)
	}

	// Let's add another user, and then update this user
	random_username = strings.ToLower(getString(5, 'A', 'Z'))
	random_name := strings.ToLower(getString(6, 'A', 'Z'))
	random_surname := strings.ToLower(getString(7, 'A', 'Z'))
	random_description := strings.ToLower(getString(8, 'A', 'Z'))

	t = sqlite06.Userdata{
		Username:    random_username,
		Name:        random_name,
		Surname:     random_surname,
		Description: random_description,
	}

	// Let's add this user
	fmt.Println("Adding user with username:", random_username)
	id = sqlite06.AddUser(t)
	fmt.Printf("User with username: %v has been added, and its id is: %v\n", random_username, id)

	// Print out all users
	data, _ = sqlite06.ListUsers()

	for _, v := range data {
		fmt.Println(v)
	}

	// Print out specified user
	data, _ = sqlite06.ListUsers()

	for _, v := range data {

		if v.Username == random_username {
			fmt.Println("Before update operation the values are:\n", v)
		}

	}

	// Let's update this user
	t.Name = "Sonny"
	t.Surname = "Blackbones"
	t.Description = "Father of D'Jok"

	fmt.Printf("Updating user with username: %v\n", t.Username)

	err = sqlite06.UpdateUser(t)
	if err != nil {
		fmt.Printf("An error occurred during the update of user with username: %v\n", random_username)
	}

	// Print out all users
	data, _ = sqlite06.ListUsers()

	for _, v := range data {
		fmt.Println(v)
	}

	fmt.Println("Printing values from tables after all operations:")
	// Print out all users
	data, _ = sqlite06.ListUsers()

	for _, v := range data {
		fmt.Println(v)
	}

	// Example Output
	// │Adding user with username: buqim                │
	// │User with username: buqim has been added, and it│
	// │s id is: 1                                      │
	// │{1 buqim Tuna Cinsoy Author}                    │
	// │Deleting user with username: buqim              │
	// │User with username: buqim has been deleted, and │
	// │its id was: 1                                   │
	// │Adding user with username: eykmq                │
	// │User with username: eykmq has been added, and it│
	// │s id is: 1                                      │
	// │{1 eykmq keexjo nevvrvw cucjhctx}               │
	// │Before update operation the values are:         │
	// │ {1 eykmq keexjo nevvrvw cucjhctx}              │
	// │Updating user with username: eykmq              │
	// │{1 eykmq Sonny Blackbones Father of D'Jok}      │
	// │Printing values from tables after all operations│
	// │:                                               │
	// │{1 eykmq Sonny Blackbones Father of D'Jok}      │

}
