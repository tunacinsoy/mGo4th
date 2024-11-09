/*
The aim of this program is to make JSON outputs as more human-readable.
*/
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
)

type Data struct {
	Name string `json:"key"`
	Val  int    `json:"value"`
}

func getRandomInt(min int, max int) int {
	return rand.Intn(max-min) + min
}

// min is the lower range of characters interval
// max is the upper range of characters interval
// For instance -> parameters 5, 'A' and 'Z' would return:
// random generated string that has length of 5 and all characters are within the ASCII interval of A-Z
func getString(length int, firstChar byte, secondChar byte) string {
	min := int(firstChar)
	max := int(secondChar)
	temp := ""
	var i int = 1
	for {

		myRand := getRandomInt(min, max)
		newChar := byte(myRand)
		temp = temp + string(newChar)

		if i == length {
			break
		}

		i++
	}

	return temp
}

func PrettyPrint(v ...Data) (retByte []byte, err error) {
	b, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		return b, err
	} else {
		return b, nil
	}
}

func JSONstream(data interface{}) (prettyString string, retErr error) {
	buffer := new(bytes.Buffer)
	encoder := json.NewEncoder(buffer)
	encoder.SetIndent("", "\t")

	err := encoder.Encode(data)
	if err != nil {
		return "Error occurred", err
	} else {
		return buffer.String(), nil
	}
}

func main() {

	dataSlice := []Data{}

	t := Data{}

	for i := 0; i < 3; i++ {
		t.Name = getString(5, 'A', 'Z')
		t.Val = getRandomInt(5, 100)

		dataSlice = append(dataSlice, t)
	}

	fmt.Printf("Last record %v\n", t)

	prettyOutput, err := PrettyPrint(dataSlice...)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(prettyOutput))

	val, _ := JSONstream(dataSlice)

	fmt.Printf("Streamed JSON data: %v\n", val)

}
