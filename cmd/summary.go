package cmd

import (
	"fmt"
	"log"

	"github.com/GRACENOBLE/expense-tracker/pkg/helpers"
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

		helpers.CreateFileIfNotExists("output/expenses.json")

		empty, err := helpers.IsFileEmpty("output/expenses.json")
		if err != nil {
			log.Fatal(err)
		}

		if empty {
			log.Fatal("You dont have any Expenses")
		}

		expenses, err := helpers.ReadExpensesFromFile("output/expenses.json")
		if err != nil {
			log.Fatal(err)
		}

		var total int64

		for _, expense := range expenses {

			total += expense.Amount

		}

		fmt.Printf("Summary : %v\n", total)
	},
}
