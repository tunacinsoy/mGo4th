// The aim of this program is to generate simple standard deviation calculator.
package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	arguments := os.Args

	if len(arguments) == 1 {
		log.Fatal("Please provide numeric values to calculate standard deviation.")
	}

	var totalCount uint
	var minVal float64
	var maxVal float64
	var initialized uint8
	var sum float64

	for _, val := range arguments {
		val, err := strconv.ParseFloat(val, 64)

		if err != nil {
			continue
		}

		if initialized == 0 {
			minVal = val
			maxVal = val
			totalCount++
			initialized++
			sum += val
			continue
		}

		if val < minVal {
			minVal = val
		} else if val > maxVal {
			maxVal = val
		}

		totalCount++
		sum += val
	}

	fmt.Printf("#total: %d, #max: %.2f, #min: %.2f, #sum: %.2f \n", totalCount, maxVal, minVal, sum)

	meanValue := sum / float64(totalCount)
	fmt.Printf("MeanVal: %.3f\n", meanValue)

	var squared float64

	for _, val := range arguments {
		val, err := strconv.ParseFloat(val, 64)

		if err != nil {
			continue
		}

		squared = squared + math.Pow((val-meanValue), 2)

	}

	standardDeviation := math.Sqrt((squared / float64(totalCount)))

	fmt.Printf("Std. devation: %.5f\n", standardDeviation)

}
