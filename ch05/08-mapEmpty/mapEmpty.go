// The aim of this program is to deal with data that its type is unknown using interface{}
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Record struct {
	flag bool
	sli1 []string
	map1 map[string]string
	str1 string
}

var JSONrecord = `{
	"Flag": true,
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

	exploreMap(JSONMap)
	typeSwitch(JSONMap)

}
