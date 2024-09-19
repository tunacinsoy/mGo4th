// The aim of this go code is to explain the importance of the location of curly brackets. (Line 9)
// The reason of the error is that go compiler adds semicolon to places where it thinks that is necessary, and if we put
// curly brackets in the line after declaring main(), compiler will put semicolon after main() declaration, which results in error.
package main

import (
	"fmt"
)

func main()
{
	fmt.Println("Go has strict rules for curly braces!")
}