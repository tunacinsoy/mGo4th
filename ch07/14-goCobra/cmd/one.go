/*
This program is created with command `cobra-cli add one`; and can be called with command:
`go run main.go one`
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// oneCmd represents the one command
var oneCmd = &cobra.Command{
	Use:     "one",
	Aliases: []string{"cmd1"},
	Short:   "command one",
	Long:    `A longer description for command one.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("one called")
	},
}

func init() {
	rootCmd.AddCommand(oneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// oneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// oneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
