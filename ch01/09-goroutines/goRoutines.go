// This program aims to shortly show the execution of goRoutines.
package main

import (
	"fmt"
	"time"
)

func myPrint(start int, finish int) {
	for i := start; i <= finish; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
	time.Sleep(100 * time.Microsecond)
}

func main() {

	for i := 0; i < 4; i++ {
		// After go myPrint() is called, the function does not block the next line of code in the main() function.
		// The main program keeps running, and myPrint() executes independently.
		go myPrint(i, 5)
	}

	// Importantly, goroutines don’t block the termination of the main function.
	// This means if the main program ends before the goroutine finishes its task, the goroutine will also be terminated.
	// This is a common issue that leads to unexpected behavior in simple examples.
	// That's why we are adding this sleep function here, so that we'll wait until myPrint() goRoutine finishes.
	time.Sleep(time.Second)

}
