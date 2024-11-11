package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(addExpense)
}

var addExpense = &cobra.Command{
	Use:   "add",
	Short: "Add an expense to the JSON",
	Long:  "This command takes in command line arguments and uses them to create an ",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Adding expense")
	},
}
