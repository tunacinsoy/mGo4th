/*
The aim of this program is to get familiar with encoding and decoding for serialization/deserialization of data.
Serialization means converting the data into JSON Format.
Deserialization means that converting the JSON data into usable format (struct type) in program.
*/
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
)

type Data struct {
	Key string `json:"key"`
	Val int    `json:"value"`
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

func DeSerialize(e *json.Decoder, slice interface{}) error {
	return e.Decode(slice)
}

func Serialize(e *json.Encoder, slice interface{}) error {
	return e.Encode(slice)
}

var (
	DataRecords []Data
)

func main() {

	t := Data{}

	for i := 0; i < 3; i++ {
		t.Key = getString(5, 'A', 'Z')
		t.Val = getRandomInt(1, 100)

		DataRecords = append(DataRecords, t)

	}

	// In encoding part, I assume that I am preparing a buffer that will hold the JSON data (serialized data)
	// of the data that i have as struct type.

	buf := new(bytes.Buffer)

	encoder := json.NewEncoder(buf)

	// Shorter version
	// err = json.NewEncoder(buf).Encode(DataRecords)

	err := Serialize(encoder, DataRecords)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("After serialization: %v\n", buf)

	// In decoding part, I assume that I received/had a buffer that holds the JSON data (serialized data)
	// that I need to decode (convert it into struct type) to use in my program.
	decoder := json.NewDecoder(buf)
	temp := []Data{}

	// Shorter version
	// err = json.NewDecoder(buf).Decode(DataRecords)

	err = DeSerialize(decoder, &temp)
	fmt.Println("After DeSerialize:")
	for index, value := range temp {
		fmt.Printf("Index: %v, Value: %v\n", index, value)
	}
}
