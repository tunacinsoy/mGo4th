/*
	The aim of this program is to use embed package to embed the contents of files (image, text)
	into respective variables, and then writing the contents of these variables to new files.

	Data that will be embedded can be found in inputs folder.
	Outputs can be seen in outputs folder.
*/

package main

import (
	// As the embed package is not used directly, we need to put underscore in front of it
	_ "embed"
	"fmt"
	"os"
)

// Following `go:embed`` directives embeds the content of the files into respective variables

//go:embed inputs/image.png
var f1 []byte

//go:embed inputs/textfile
var f2 string

func writeToFile(dataToWrite interface{}, destinationFilePath string) (err error) {

	switch dataToWrite.(type) { // Type switch to check the variable type of input `dataToWrite`
	case []byte:
		fd, err := os.OpenFile(destinationFilePath, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println("Error while opening file")
			return err
		}
		defer fd.Close()

		n, err := fd.Write(dataToWrite.([]byte)) // Type assertion, since we know that input x is `[]byte` type
		if err != nil {
			return err
		}
		fmt.Printf("Wrote %d bytes.\n", n)
	case string:
		fd, err := os.OpenFile(destinationFilePath, os.O_CREATE|os.O_WRONLY, 0644)

		if err != nil {
			fmt.Println("Error while opening file")
			return err
		}
		defer fd.Close()

		n, err := fd.WriteString(dataToWrite.(string)) // Type assertion, since we know that input x is `string` type
		if err != nil {
			return err
		}
		fmt.Printf("Wrote %d bytes.\n", n)

	}
	return nil
}
func main() {

	if len(os.Args) == 1 {
		fmt.Println("Please input 1 or 2.")
		return
	}

	fmt.Println("f1:", len(f1), "f2", len(f2))

	switch os.Args[1] {
	case "1":
		filename := "./outputs/newImage.png"
		err := writeToFile(f1, filename)
		if err != nil {
			fmt.Println(err)
			return
		}

	case "2":
		fileName := "./outputs/newTextfile"
		err := writeToFile(f2, fileName)
		if err != nil {
			fmt.Println(err)
			return
		}
	default:
		fmt.Println("Please specify 1 or 2 as input.")
	}
}
