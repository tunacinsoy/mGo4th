/*
The aim of this program is to use `/dev/random` system device	to generate seed for
random number generator.
The program also compares various methods for random number generation.
*/
package main

import (
	"encoding/binary"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	f, err := os.Open("/dev/random")
	defer f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	// ----------------------------------------------------------------------------------------
	var seed1 int64
	var seed2 int64
	// IF data is 0x123456, then:
	// Little Endian: 0x 65 42 31
	// Big Endian: 0x 12 34 56
	binary.Read(f, binary.LittleEndian, &seed1) // 	-2852676401199697018
	binary.Read(f, binary.BigEndian, &seed2)    // -7070070872815073262
	fmt.Println(seed1)
	fmt.Println(seed2)
	r0 := rand.New(rand.NewSource(seed1))
	fmt.Printf("Generated val is: %.5f\n", r0.Float64()) // Generated val is: 0.50333
	// ----------------------------------------------------------------------------------------
	// Alternatively, we can use following code block for random float64 generation
	seed := time.Now().Unix()
	r := rand.New(rand.NewSource(seed))
	fmt.Printf("Alternatively, generated val is: %.5f\n", r.Float64()) // Alternatively, generated val is: 0.97321
	// ----------------------------------------------------------------------------------------
	// This block will generate same number every time, since we use a constant seed value
	const seed_2 int64 = 15
	r1 := rand.New(rand.NewSource(seed_2))
	fmt.Printf("Constant generated val is: %.5f\n", r1.Float64()) // Constant generated val is: 0.40755
	// ----------------------------------------------------------------------------------------
}
