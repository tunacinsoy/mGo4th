// The aim of this program is to introduce normalization for the statistics application
package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"time"
)

func randomFloat(min float64, max float64, r *rand.Rand) float64 {
	return r.Float64()*(max-min) + min
}

func normalize(data []float64, meanVal float64, standardDeviation float64) []float64 {
	if standardDeviation == 0 {
		return data
	}
	normalized := make([]float64, len(data))
	for i, val := range data {
		normalized[i] = math.Floor((val-meanVal)/standardDeviation*10000) / 10000
	}

	return normalized
}

func main() {
	arguments := os.Args

	if len(arguments) == 1 {
		fmt.Println("Need one or more arguments")
		log.Fatal()
	}

	values := []float64{}

	// omitting file name
	arguments = arguments[1:]

	for _, val := range arguments {
		n, err := strconv.ParseFloat(val, 64)
		if err != nil {

			continue
		}
		values = append(values, n)
	}

	if len(values) == 0 {
		fmt.Println("Generating 10 random values...")
		r := rand.New(rand.NewSource(time.Now().Unix()))

		for i := 0; i < 10; i++ {
			val := randomFloat(-10, 10, r)
			values = append(values, val)
		}
	}

	sort.Float64s(values)

	fmt.Println("Count of values:", len(values))
	fmt.Println("Min:", values[0])
	fmt.Println("Max:", values[len(values)-1])

	var sum float64 = 0

	for _, val := range values {
		sum += val
	}

	meanValue := sum / float64(len(values))
	fmt.Printf("Mean value is: %.5f:\n", meanValue)

	var squared float64 = 0

	for _, val := range values {
		squared += math.Pow((val - meanValue), 2)
	}

	standardDeviation := math.Sqrt((squared / float64(len(values))))
	fmt.Printf("Standard Deviation is: %.5f\n", standardDeviation)

	// normalization
	normalized := normalize(values, meanValue, standardDeviation)
	fmt.Println("Normalized:", normalized)
}
