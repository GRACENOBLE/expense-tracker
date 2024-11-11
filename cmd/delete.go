package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(deleteExpense)
}

var deleteExpense = &cobra.Command{
	Use:   "delete",
	Short: "Delete a task",
	Long:  "This function takes a command line argument and deletes a corresponding task.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Deleting expense")
	},
}

