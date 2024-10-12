// The aim of this program is to deal with data that its type is unknown using interface{}
package main

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
)

var JSONrecord = `{
	"Flag": true,
	"Integer1": 2,
	"Float1": 3.4,
	"Array": ["a","b","c"],
	"Entity": {
		"a1": "b1",
		"a2": "b2",
		"Value": "-456",
		"Null": null
	},
	"Message": "Hello Go!"
}`

func typeSwitch(m map[string]interface{}) {
	for k, v := range m {
		switch c := v.(type) {
		case string:
			fmt.Println("Is a string!", k, c)
		case float64:
			fmt.Println("Is a float64!", k, c)
		case bool:
			fmt.Println("Is a Boolean!", k, c)
		case map[string]interface{}:
			fmt.Println("Is a map!", k, c)
			typeSwitch(v.(map[string]interface{}))
		default:
			fmt.Printf("...Is %v: %T!\n", k, c)
		}
	}
	return
}

// the input parameter looks confusing, but it means that i know that i will receive a map that contains string type of keys, however I am not sure about the data types of values.W
func exploreMap(m map[string]interface{}) {

	for k, v := range m {
		embMap, ok := v.(map[string]interface{})
		// If it is a map, explore deeper
		if ok {
			fmt.Printf("{\"%v\": \n", k)
			exploreMap(embMap)
			fmt.Printf("}\n")
		} else {
			fmt.Printf("%v: %v\n", k, v)
		}
	}
}

func hasDecimal(f interface{}) bool {
	v, ok := f.(float64)
	if ok {
		_, frac := math.Modf(v)
		if frac == 0 {
			return false
		} else {
			return true
		}
	}
	return false
}

// Analyzes and returns the count of each data type of a json record
func analyzeMap(m map[string]interface{}) {

	var counterMap int = 0
	var counterInt int = 0
	var counterFloat64 int = 0
	var counterString int = 0
	var counterBool int = 0
	var counterSlice int = 0
	for _, v := range m {

		switch T := v.(type) {
		case map[string]interface{}:
			counterMap++
		case int:
			// Somehow it does not fall into this case, this part will be handled in float64 case
			counterInt++
		case float64:
			if hasDecimal(v) {
				counterFloat64++
			} else {
				counterInt++
			}
		case string:
			counterString++
		case bool:
			counterBool++
		case []interface{}:
			counterSlice++
		default:
			fmt.Printf("Not supported type: %T\n", T)
		}
	}

	fmt.Printf("Analysis Results:\n")
	fmt.Printf("\t Count of maps: %d\n", counterMap)
	fmt.Printf("\t Count of ints: %d\n", counterInt)
	fmt.Printf("\t Count of floats: %d\n", counterFloat64)
	fmt.Printf("\t Count of strings: %d\n", counterString)
	fmt.Printf("\t Count of bools: %d\n", counterBool)
	fmt.Printf("\t Count of slice: %d\n", counterSlice)
}

func main() {

	if len(os.Args) == 1 {
		fmt.Println("Using def json record.")
	} else {
		JSONrecord = os.Args[1]
	}

	// map[string]interface{} is extremely handy for storing JSON records when you do not know
	// their schema in advance. In other words, map[string]interface{} is good at storing arbitrary
	// JSON data of unknown schema.
	JSONMap := make(map[string]interface{})

	err := json.Unmarshal([]byte(JSONrecord), &JSONMap)

	if err != nil {
		fmt.Println(err)
		return
	}

	analyzeMap(JSONMap)
	// exploreMap(JSONMap)
	// typeSwitch(JSONMap)

}
