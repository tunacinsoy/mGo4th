/*
The aim of this program is to understand reading json formatted data from a config file
using package viper.

Usage:

	// If go.mod and go.sum files do not exist where this program resides, then run these first:
	`go mod init <module_name>`
	`go mod tidy`

	// To execute program:
	`go run jsonViper.go <file_path>`

	// Sample execution:
	`go run jsonViper.go /home/tuna/github-repos/mGo4th/ch07/13-jsonViper/myConfig.json`

	// Example output:

	Using config: /home/tuna/github-repos/mGo4th/ch07/13-jsonViper/myConfig.json
	macos: pass_macos
	P: machine1 MySQL: machine2 mongo: machine3
	DoesNotExist is not set.

[

	{
		"MacPass": "pass_macos",
		"LinuxPass": "pass_linux",
		"WindowsPass": "pass_windows",
		"ActiveFlag": true,
		"PostHost": "machine1",
		"MySQLHost": "machine2",
		"MongoDBHost": "machine3"
	}

]
*/
package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/viper"
)

/*
Use mapstructure for configuration libraries like Viper, as they rely on it for decoding structured configuration data into structs.
Use json when you're serializing/deserializing JSON using Go's encoding/json package.
*/
type Config struct {
	MacPass     string `mapstructure:"macos"`
	LinuxPass   string `mapstructure:"linux"`
	WindowsPass string `mapstructure:"windows"`
	ActiveFlag  bool   `mapstructure:"active"`
	PostHost    string `mapstructure:"postgres"`
	MySQLHost   string `mapstructure:"mysql"`
	MongoDBHost string `mapstructure:"mongodb"`
}

func PrettyPrint(v ...Config) (err error) {
	b, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		return err
	} else {
		fmt.Println(string(b))
		fmt.Println()
		return nil
	}
}

func main() {

	config := "myConfig.json"

	if len(os.Args) == 1 {
		fmt.Println("Using default file")
	} else {
		config = os.Args[1]
	}

	viper.SetConfigType("json")
	viper.SetConfigFile(config)

	fmt.Printf("Using config: %s\n", viper.ConfigFileUsed())

	// Read values in json file
	viper.ReadInConfig()

	// check whether a value is set or not
	if viper.IsSet("macos") {
		fmt.Println("macos:", viper.Get("macos"))
	} else {
		fmt.Println("macos is not set.")
	}

	if viper.IsSet("active") {
		value := viper.Get("active")
		if value.(bool) { // type assertion
			postgres := viper.Get("postgres")
			mysql := viper.Get("mysql")
			mongo := viper.Get("mongodb")
			fmt.Println("P:", postgres, "MySQL:", mysql, "mongo:", mongo)
		} else {
			fmt.Println("active value is false, omitting print statements.")
		}
	} else {
		fmt.Println("active is not set.")
	}

	if viper.IsSet("DoesNotExist") {
		fmt.Printf("Val of DoesNotExist is: %v\n", viper.Get("DoesNotExist"))
	} else {
		fmt.Println("DoesNotExist is not set.")
	}
	// Unmarshal operation
	var t Config
	err := viper.Unmarshal(&t)
	if err != nil {
		fmt.Println(err)
		return
	}
	PrettyPrint(t)
}
