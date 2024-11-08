/*
	The aim of this program is to learn json tags 'omitempty' and '-'
	These tags are used for omitting field if it is not declared,
	and ignoring sensitive fields for json.marshal operation.
*/

package main

import (
	"encoding/json"
	"fmt"
)

// Ignoring empty fields in JSON
type NoEmpty struct {
	Name    string `json:"username"`
	Surname string `json:"surname"`
	Year    int    `json:"creationyear,omitempty"` // If this field is not declared, json.marshal won't include it
}

// Removing private fields and ignoring empty fields
type Password struct {
	Name    string `json:"username"`
	Surname string `json:"surname,omitempty"`
	Year    int    `json:"creationyear,omitempty"`
	Pass    string `json:"-"` //This will be ignored when converting a Password structure into a JSON record using json.Marshal
}

func main() {

	fullInput := NoEmpty{Name: "Tuna", Surname: "Cinsoy", Year: 1999}
	emptyInput := NoEmpty{Name: "Tuna"}
	marshaledFullInput, err := json.Marshal(&fullInput)

	if err != nil {
		fmt.Println(err)
		return
	}

	marshaledEmptyInput, err := json.Marshal(&emptyInput)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("MarshaledFullInput is: %v\n", string(marshaledFullInput))   // MarshaledFullInput is: {"username":"Tuna","surname":"Cinsoy","creationyear":1999}
	fmt.Printf("MarshaledEmptyInput is: %v\n", string(marshaledEmptyInput)) // MarshaledEmptyInput is: {"username":"Tuna","surname":""}

	// ---

	pwdInput := Password{Name: "Tuna", Surname: "Cinsoy", Year: 1999, Pass: "fullsecret"}

	marshaledPwdInput, err := json.Marshal(&pwdInput)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("MarshaledPwdInput is: %s\n", marshaledPwdInput) // MarshaledPwdInput is: {"username":"Tuna","surname":"Cinsoy","creationyear":1999}

}
