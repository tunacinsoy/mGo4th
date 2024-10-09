// The aim of this program is to get familiar with reflect package
// This package is used in order to determine the data types of elements of structs,
// retrieve their values. It is basically about understanding fields of structs.
package main

import (
	"fmt"
	"reflect"
)

type Secret struct {
	Username string
	Password string
}

type Record struct {
	Field1 string
	Field2 float64
	Field3 Secret
}

func main() {

	A := Record{Field1: "String value", Field2: -12.123, Field3: Secret{"Mihalis", "Tsoukalos"}}
	fmt.Printf("Type of and val of A: %T, %v\n", A, A) // 	Type of and val of A: main.Record, {String value -12.123 {Mihalis Tsoukalos}}

	r := reflect.ValueOf(A)

	fmt.Println("String value of r:", r.String()) // String value of r: <main.Record Value>
	// 	main.Record is the full unique name of the structure as defined by Go—main is the package name
	// and Record is the struct name. This happens so that Go can differentiate between the elements
	// of different packages.

	iType := r.Type()
	fmt.Printf("i Type: %s\n", iType)                            // i Type: main.Record
	fmt.Printf("The %d fields of %s are\n", r.NumField(), iType) // The 3 fields of main.Record are

	for i := 0; i < r.NumField(); i++ {
		fmt.Printf("\t%s ", iType.Field(i).Name)
		fmt.Printf("\twith type: %s ", r.Field(i).Type())
		fmt.Printf("\tand value: %v\n", r.Field(i).Interface())

		// Field1 	with type: string 	and value: String value
		// Field2 	with type: float64 	and value: -12.123
		// Field3 	with type: main.Secret 	and value: {Mihalis Tsoukalos}

		// Check whether there are other structures embedded in Record
		k := reflect.TypeOf(r.Field(i).Interface()).Kind()

		// Need to convert it to string in order to compare it
		if k.String() == "struct" {
			fmt.Println(r.Field(i).Type()) // main.Secret
		}

		// Same as before if
		if k == reflect.Struct {
			fmt.Println(r.Field(i).Type()) // main.Secret
		}
	}
}
