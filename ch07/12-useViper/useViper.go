/*
	The aim of this program is to get familiar with defining command line flags using viper package.

	Usage:
	1. Run `go mod init <mod_name>`
	2. Run `go mod tidy`
	3. Run `go run useViper.go --help` to find out which flags ara available to use

	Some example runs and their outputs:
	`go run useViper.go --pass tunacinsoy`
	Mike tunacinsoy
	GOMAXPROCS is: 16

	`go run useViper.go -n tunacinsoy`
	tunacinsoy hardToGuess
	GOMAXPROCS is: 16

	`go run useViper.go -nm tunacinsoy`
	m hardToGuess # (Why? Because for flags which have more than one letter, we need to use double dash: `--`)
	GOMAXPROCS is: 16

	`go run useViper.go --name tuna --password supersecret`
	tuna supersecret
	GOMAXPROCS is: 16

*/

package main

import (
	"fmt"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func aliasNormalizeFunc(f *pflag.FlagSet, n string) pflag.NormalizedName {
	// These cases will also be acknowledged as flags (good for recovering user mistakes, or laziness)
	switch n {
	case "pass":
		n = "password"
	case "ps":
		n = "password"
	case "nm":
		n = "name"
	case "nme":
		n = "name"
	}
	return pflag.NormalizedName(n)
}

func main() {
	// Create a flag which has `name` as its name; `n` for its short name; `Name parameter` as its description
	// `StringP is like String, but accepts a shorthand letter that can be used after a single dash.`
	pflag.StringP("name", "n", "Mike", "Name parameter")

	// Create a flag which has `password` as its name; `p` for its short name; `Password` as its description
	pflag.StringP("password", "p", "hardToGuess", "Password")

	// it would be possible to create a flag named "getURL" and have it translated to "geturl"
	pflag.CommandLine.SetNormalizeFunc(aliasNormalizeFunc)

	// Parse all flags defined in command line, should be indicated after all flags are defined
	pflag.Parse()

	// Make the command line flags available to use for viper package
	viper.BindPFlags(pflag.CommandLine)

	// Getting command line flag values
	name := viper.GetString("name")
	password := viper.GetString("password")

	fmt.Println(name, password)

	// --------------------------------------------------------
	// Receiving env variable to be used in viper package
	viper.BindEnv("GOMAXPROCS")
	val := viper.Get("GOMAXPROCS")

	// It was defined before, so this if won't hold
	if val != nil {
		fmt.Println("GOMAXPROCS:", val)
	}

	// --------------------------------------------------------

	// Setting env variable
	viper.Set("GOMAXPROCS", 16)
	val = viper.Get("GOMAXPROCS")
	fmt.Println("GOMAXPROCS is:", val)
}
