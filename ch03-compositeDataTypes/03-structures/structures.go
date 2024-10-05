// The aim of this program is to get familiar with struct date type
package main

import "fmt"

type Entry struct {
	Name    string
	Surname string
	Year    int
}

func zeroS() Entry {
	var p Entry = Entry{}
	return p
}

func zeroPtoS() *Entry {
	var p *Entry = &Entry{}
	return p
}

func initS(name string, surname string, year int) Entry {
	if year < 1940 {
		return Entry{name, surname, 1941}
	}
	return Entry{name, surname, year}
}

func initPtoS(name string, surname string, year int) *Entry {
	if len(surname) == 0 {
		return &Entry{Name: name, Surname: "Unknown", Year: year}
	}
	var p *Entry = &Entry{name, surname, year}
	return p
}
func main() {
	s1 := zeroS()
	fmt.Println("s1:", s1)

	p1 := zeroPtoS()

	// Why did we put & before p1? It is already a pointer.
	fmt.Printf("p1: %v\n", &p1)

	s2 := initS("Mihalis", "Tsoukalos", 2024)
	fmt.Printf("s2: %v\n", s2)

	p2 := initPtoS("Mihalis", "Tsoukalos", 2024)
	fmt.Printf("p2: %v\n", *p2)

	fmt.Println("Year:", s1.Year, s2.Year, p1.Year, p2.Year)
	pS := new(Entry)
	fmt.Println("pS:", *pS)

	// var x int = 5
	// var xPtr *int = &x
	// fmt.Printf("x is: %v, xPtr is: %v\n", x, xPtr)

}
