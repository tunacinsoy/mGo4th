// The aim of this program is to enhance the existing stats program by incorporating readFile utility from a CSV file

package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
)

func readFile(filepath string) ([]float64, error) {
	_, err := os.Stat(filepath)
	if err != nil {
		return nil, err
	}
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// This function might throw an error if each line in data file does not contain
	// same number of fields. for instance,
	// line 1 -> 1.1,2
	// line 2 -> 2
	// would cause an error due the difference between line 1 and line 2 in terms of value counts.
	lines, err := csv.NewReader(f).ReadAll()
	fmt.Printf("Type of lines: %T\n", lines)
	fmt.Printf("lines: %v\n", lines)
	if err != nil {
		return nil, err
	}
	//Create a float64 slice with size 0 and capacity 0
	values := make([]float64, 0, 0)

	for _, line := range lines {
		fmt.Println("len(line)", len(line))
		tmp, err := strconv.ParseFloat(line[0], 64)
		if err != nil {
			log.Println("Error reading:", line[0], err)
			continue
		}
		values = append(values, tmp)
	}

	return values, nil
}

func stdDev(x []float64) (float64, float64) {
	sum := 0.0
	for _, val := range x {
		sum += val
	}
	// Mean value
	meanValue := sum / float64(len(x))
	fmt.Printf("Mean value is: %.3f\n", meanValue)
	// Std devation
	var squaredSum float64

	for _, val := range x {
		squaredSum += math.Pow(meanValue-val, 2)
	}

	var standardDeviation float64 = math.Sqrt(squaredSum / float64(len(x)))

	return meanValue, standardDeviation
}

func normalize(data []float64, mean float64, stdDev float64) []float64 {
	if stdDev == 0 {
		return data
	}

	normalized := make([]float64, len(data))
	for i, val := range data {
		normalized[i] = math.Floor((val-mean)/stdDev*10000) / 10000
	}

	return normalized
}

func main() {

	if len(os.Args) == 1 {
		log.Println("Need one argument!")
		return
	}

	file := os.Args[1]
	values, err := readFile(file)
	if err != nil {
		log.Println("Error reading:", file, err)
		os.Exit(0)
	}
	sort.Float64s(values)

	fmt.Println("Number of values:", len(values))
	fmt.Println("Min:", values[0])
	fmt.Println("Max:", values[len(values)-1])

	meanValue, standardDeviation := stdDev(values)
	fmt.Printf("Standard deviation: %.5f\n", standardDeviation)

	// Now normalize the list of values
	normalized := normalize(values, meanValue, standardDeviation)
	fmt.Println("Normalized:", normalized)
}
