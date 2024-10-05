// The aim of this program is to get familiar with type package, and time conversion / time parsing
package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Please supply an input line argument.")
		log.Fatal()
	}

	input := os.Args[1]

	now, err := time.Parse("02 January 2006 15:04 MST", input)

	if err != nil {
		fmt.Println("Err occurred during time parse")
		log.Fatal()
	}

	//fmt.Println("now", now)

	//loc, err := time.LoadLocation("Local")

	//	fmt.Printf("Curre")

	const longForm = "Jan 2, 2006 at 3:04pm (MST)"
	// A string gets converted/parsed into time.Time value according to the
	// form that it gets compared to.
	t, _ := time.Parse(longForm, "Feb 3, 2013 at 7:54pm (PST)")
	fmt.Printf("%T, %s\n", t, t) // time.Time, 2013-02-03 19:54:00 +0000 PST

	// Input is: "14 December 2023 19:20 EET"
	// Local time
	loc, _ := time.LoadLocation("Local")
	fmt.Printf("Current Location: %s\n", now.In(loc))

	// NY
	loc, _ = time.LoadLocation("America/New_York")
	fmt.Printf("New York Time: %s\n", now.In(loc))

	// London
	loc, _ = time.LoadLocation("Europe/London")
	fmt.Printf("London Time: %s\n", now.In(loc))

	// Tokyo
	loc, _ = time.LoadLocation("Asia/Tokyo")
	fmt.Printf("Tokyo Time: %s\n", now.In(loc))
}
