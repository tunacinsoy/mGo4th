// The aim of this program is to introduce custom sort behavior to statistics app
package main

import (
	"encoding/csv"
	"fmt"
	"math"
	"os"
	"slices"
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
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return nil, err
	}

	values := make([]float64, 0)
	for _, line := range lines {
		tmp, err := strconv.ParseFloat(line[0], 64)
		if err != nil {
			fmt.Println("Error reading:", line[0], err)
			continue
		}
		values = append(values, tmp)
	}

	return values, nil
}

func stdDev(x []float64) (float64, float64) {
	sum := float64(0)
	for _, val := range x {
		sum = sum + val
	}

	meanValue := sum / float64(len(x))
	fmt.Printf("Mean value: %.5f\n", meanValue)

	// Standard deviation
	var squared float64
	for i := 0; i < len(x); i++ {
		squared = squared + math.Pow((x[i]-meanValue), 2)
	}

	standardDeviation := math.Sqrt(squared / float64(len(x)))
	return meanValue, standardDeviation
}

type DataFile struct {
	Filename string
	Len      int
	Minimum  float64
	Maximum  float64
	Mean     float64
	StdDev   float64
}

type DFslice []DataFile

// Implement sort.Interface
// Why? Because we would like to sort our fields according to their mean values,
// which is a custom behavior.
func (a DFslice) Len() int {
	return len(a)
}
func (a DFslice) Less(i int, j int) bool {
	return a[i].Mean < a[j].Mean
}
func (a DFslice) Swap(i int, j int) {
	a[i], a[j] = a[j], a[i]
}

func main() {

	if len(os.Args) == 1 {
		fmt.Println("Please provide file paths.")
		return
	}

	files := DFslice{}
	// another definition
	// files1 := []DataFile{}

	os.Args = os.Args[1:]

	for _, filename := range os.Args {
		currentFile := DataFile{}
		currentFile.Filename = filename

		values, err := readFile(filename)

		if err != nil {
			fmt.Println("Error reading: ", filename, err)
			os.Exit(1)
		}

		currentFile.Len = len(values)
		currentFile.Len = len(values)
		currentFile.Minimum = slices.Min(values)
		currentFile.Maximum = slices.Max(values)
		meanValue, standardDeviation := stdDev(values)
		currentFile.Mean = meanValue
		currentFile.StdDev = standardDeviation

		files = append(files, currentFile)
	}

	sort.Sort(files)

	for _, val := range files {
		f := val.Filename
		fmt.Println(f, ":", val.Len, val.Mean, val.Maximum, val.Minimum)
	}

	// go run stats.go d1.txt d2.txt d3.txt
	// OUTPUTS
	// Mean value: 3.00000
	// Mean value: 18.20000
	// Mean value: 0.75000
	// d3.txt : 4 0.75 102 -300
	// d1.txt : 5 3 5 1
	// d2.txt : 5 18.2 100 -4

}
