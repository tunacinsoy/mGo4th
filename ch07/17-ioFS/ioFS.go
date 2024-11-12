/*
The aim of this program is to use io/fs package to:
- search file in a given directory (which is embed.FS type)
- list files in a given directory (which is embed.FS type)
- extract the contents of a file in a given directory (which is embed.FS type)
- writing to a file

To achieve this, we use embed package to embed the contents of 'static' folder.
*/
package main

import (
	"embed" // Since we have a variable called f that is type of embed.FS, we do not need to put _ here.
	"fmt"
	"io/fs"
	"os"
)

//go:embed static
var f embed.FS

var searchString string

func walkFunction(path string, d fs.DirEntry, err error) error {
	if err != nil {
		return err
	}
	fmt.Printf("Path:%q, isDir:%v\n", path, d.IsDir())
	return nil
}

func walkSearch(path string, d fs.DirEntry, err error) error {
	if err != nil {
		return err
	}
	if d.Name() == searchString {
		fileInfo, err := fs.Stat(f, path)
		fileInfo2, err := os.Lstat(path) // Alternative use case
		if err != nil {
			return err
		}
		fmt.Println("Found", path, "with size", fileInfo.Size(), "using function fs.Stat()")
		fmt.Println("Found", path, "with size", fileInfo2.Size(), "using function os.Lstat()") // Alternative call with os package
	}
	return nil
}
func list(f embed.FS) error {
	// WalkDir walks the file tree rooted at root (second input parameter),
	// calling walkFunction for each file or directory in the tree, including root (second input parameter).
	return fs.WalkDir(f, ".", walkFunction)
}

func search(f embed.FS) error {
	// WalkDir walks the file tree rooted at root (second input parameter),
	// calling walkFunction for each file or directory in the tree, including root (second input parameter).
	return fs.WalkDir(f, ".", walkSearch)
}

func writeToFile(inputBuffer []byte, path string) error {
	// Create if does not exist, if exists, append the content of file into the contents
	// at destination file, and open as write only.
	// set permissions to rw- for owner; r-- for group; r-- for others
	fd, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("an error occurred while opening file:", path)
		return err
	}
	n, err := fd.Write(inputBuffer)
	if err != nil {
		fmt.Println("an error occurred during write operation")
		return err
	}
	fmt.Printf("Wrote %v bytes to file: %v\n", n, path)
	return nil
}

func extract(f embed.FS, path string) ([]byte, error) {
	// ReadFile reads the named file from the file system fs and returns its contents
	s, err := fs.ReadFile(f, path)
	if err != nil {
		return nil, err
	}
	return s, nil
}
func main() {

	// List all files
	err := list(f)

	if err != nil {
		fmt.Println(err)
		return
	}

	// Search a file
	searchString = "file.txt"
	err = search(f)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Extract into a byte slice
	buffer, err := extract(f, "static/file.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	// At this point, buffer ,which is a []byte, contains the contents of "static/file.txt"

	err = writeToFile(buffer, "static/file3.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
}
