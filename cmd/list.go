package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listExpenses)
}

var listExpenses = &cobra.Command{
	Use:   "list",
	Short: "Long list all tasks",
	Long:  "This function lists all tasks in details.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Listing expenses")
	},
}
