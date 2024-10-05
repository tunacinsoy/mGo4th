// The aim of this program is to read data from csv and store it into another file
package main

import (
	"encoding/csv"
	"log"
	"os"
)

type Record struct {
	Name       string
	Surname    string
	Number     string
	LastAccess string
}

var myData = []Record{}

func readCSVFile(filepath string) ([][]string, error) {
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
		return [][]string{}, err
	}
	//fmt.Printf("lines[0]: %s \n", lines[0])
	return lines, nil
}

func saveCSVFile(filepath string) error {

	// If the file already exists, it will be opened, but its contents will be cleared (truncated) to a size of 0 bytes.
	csvfile, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer csvfile.Close()
	csvwriter := csv.NewWriter(csvfile)
	csvwriter.Comma = '\t'

	for _, row := range myData {
		temp := []string{row.Name, row.Surname, row.Number, row.LastAccess}
		err = csvwriter.Write(temp)
		if err != nil {
			return err
		}
	}
	csvwriter.Flush()
	return nil
}

func main() {
	if len(os.Args) != 3 {
		log.Println("csvData input output!")
		os.Exit(1)
	}

	input := os.Args[1]
	output := os.Args[2]
	lines, err := readCSVFile(input)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	// lines is 2D string slice
	// line variable represents each line in lines 2d slice
	for _, line := range lines {
		//fmt.Printf("line[0]: %v\n", line[0]) // Prints out the names
		temp := Record{
			Name:       line[0],
			Surname:    line[1],
			Number:     line[2],
			LastAccess: line[3],
		}
		myData = append(myData, temp)
		log.Println(temp)
	}

	err = saveCSVFile(output)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
