package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/GRACENOBLE/expense-tracker/pkg/helpers"
	"github.com/GRACENOBLE/expense-tracker/pkg/types"
	"github.com/spf13/cobra"
)

var month int

func init() {
	rootCmd.AddCommand(sumariseExpenses)

	sumariseExpenses.Flags().IntVarP(&month, "month", "m", 0, "Specify the month to display expenses (1 for January, 2 for February, etc.)")
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
			log.Fatal("You don't have any Expenses")
		}

		expenses, err := helpers.ReadExpensesFromFile("output/expenses.json")
		if err != nil {
			log.Fatal(err)
		}

		var total int64

		var filteredExpenses []types.Expense
		if month > 0 {
			for _, expense := range expenses {

				expenseDate, err := time.Parse("2006-01-02 15:04:05", expense.Date) // Assuming the date is in "YYYY-MM-DD" format
				if err != nil {
					log.Fatal("Invalid date format in expense data")
				}

				if int(expenseDate.Month()) == month {
					filteredExpenses = append(filteredExpenses, expense)
					total += expense.Amount
				}
			}
			if len(filteredExpenses) > 0 {

				fmt.Printf("Summary for month %d: %v\n", month, total)
			}
		} else {

			filteredExpenses = expenses
			for _, expense := range expenses {
				total += expense.Amount
			}

			fmt.Printf("Summary for all months: %v\n", total)
		}

		if len(filteredExpenses) == 0 {
			fmt.Println("No expenses found for the specified month.")
		}
	},
}
