/*
This program is created with command `cobra-cli add two`; and can be called with command:
`go run main.go two -username <user_name>` OR
`go run main.go two -u <user_name>` OR
`go run main.go two`
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// twoCmd represents the two command
var twoCmd = &cobra.Command{
	Use:     "two",
	Aliases: []string{"cmd2"},
	Short:   "command two",
	Long:    `A longer description for command two.`,
	Run: func(cmd *cobra.Command, args []string) {
		username, err := cmd.Flags().GetString("username")
		if err != nil {
			fmt.Println("Error reading username", err)
		} else {
			fmt.Println("Username:", username)
		}
	},
}

func init() {
	rootCmd.AddCommand(twoCmd)
	twoCmd.Flags().StringP("username", "u", "Tuna", "Username")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// twoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// twoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
