/*
The aim of this program is to use ReadDir function of os package to calculate total size of
a given directory as input path. If the directory contains directories, program will calculate them
in a recursive manner.
*/
package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func GetSize(path string) (size int64, err error) {

	contents, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("an error occurred while reading contents of the directory")
		return -1, err
	}

	var total int64

	for _, entry := range contents {
		if entry.IsDir() {
			// If we have a directory within the path, we need to call the func again recursively
			temp, err := GetSize(filepath.Join(path, entry.Name()))
			if err != nil {
				fmt.Println("an error occurred while calling GetSize recursively.")
				return -1, err
			}
			// Recursive func will also return a size, so we need to add that as well
			total += temp
		} else {
			info, err := entry.Info()
			if err != nil {
				return -1, err
			}

			total += info.Size()
		}

	}

	return total, nil
}

func main() {

	if len(os.Args) == 1 {
		fmt.Println("Provide a path to a directory.")
		return
	}

	// If the given path is a symlink, then this func returns the actual path
	// that is pointed by symlink.
	inputPath, err := filepath.EvalSymlinks(os.Args[1])

	if err != nil {
		fmt.Println(err)
		return
	}

	fileInfo, _ := os.Lstat(inputPath)

	if !fileInfo.IsDir() {
		fmt.Println(inputPath, "not a directory")
		return
	}

	/*
		fileInfo is a struct that has fields:
		type FileInfo interface {
			Name() string       // base name of the file
			Size() int64        // length in bytes for regular files; system-dependent for others
			Mode() FileMode     // file mode bits
			ModTime() time.Time // modification time
			IsDir() bool        // abbreviation for Mode().IsDir()
			Sys() any           // underlying data source (can return nil)
		}
	*/

	i, err := GetSize(inputPath)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Total size of the path '%v' is: %v Bytes.\n", inputPath, i) // Total size of the path '/home/tuna/Downloads' is: 1170125 Bytes.

}
