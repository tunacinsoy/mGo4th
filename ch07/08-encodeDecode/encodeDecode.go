/*
The aim of this program is to understand json marshaling and unmarshaling
Usage: `go run encodeDecode.go`
*/
package main

import (
	"encoding/json"
	"fmt"
)

// What the following metadata tells us is that the Name field of the UseAll structure is translated to
// username in the JSON record, and vice versa; the Surname field is translated to surname, and vice
// versa; and the Year structure field is translated to created in the JSON record, and vice versa
type UseAll struct {
	Name    string `json:"username"`
	Surname string `json:"surname"`
	Year    int    `json:"created"`
}

func main() {
	useall := UseAll{Name: "Mike", Surname: "Wazowski", Year: 2023}

	t, err := json.Marshal(&useall)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Value %s\n", t) // Value {"username":"Mike","surname":"Wazowski","created":2023}
	}

	str := `{"username": "T.", "surname": "Cns", "created": 2024}`
	jsonRecord := []byte(str)

	temp := UseAll{}

	err = json.Unmarshal(jsonRecord, &temp)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Data type: %T, with value %v\n", temp, temp) // Data type: main.UseAll, with value {T. Cns 2024}
	}

}
