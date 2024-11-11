package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(sumariseExpenses)
}

var sumariseExpenses = &cobra.Command{
	Use:   "sumarize",
	Short: "Sumarize all tasks",
	Long:  "This lists all tasks in a more human readable way.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Summarising Expenses.")
	},
}
