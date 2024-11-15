package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/GRACENOBLE/expense-tracker/pkg/helpers"
	"github.com/GRACENOBLE/expense-tracker/pkg/types"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listExpenses)

	listExpenses.Flags().IntVarP(&month, "month", "m", 0, "Specify the month to display expenses (1 for January, 2 for February, etc.)")
}

var listExpenses = &cobra.Command{
	Use:   "list",
	Short: "Long list all tasks",
	Long:  "This function lists all tasks in details.With the -m flag, you can list expenses from a specific month",
	Run: func(cmd *cobra.Command, args []string) {
		expenses, err := helpers.ReadExpensesFromFile("output/expenses.json")
		if err != nil {
			log.Fatal(err)
		}

		var filteredExpenses []types.Expense

		if month > 0 {
			for _, expense := range expenses {

				expenseDate, err := time.Parse("2006-01-02 15:04:05", expense.Date) // Assuming the date is in "YYYY-MM-DD" format
				if err != nil {
					log.Fatal("Invalid date format in expense data")
				}

				if int(expenseDate.Month()) == month {
					filteredExpenses = append(filteredExpenses, expense)
				}
			}

		} else {
			filteredExpenses = expenses
		}

		if len(filteredExpenses) > 0 {

			fmt.Println("ID\tDate\t\t\tDescription\tAmount")
			for _, expense := range filteredExpenses {
				fmt.Printf("%v\t%v\t%v\t\t%v\n", expense.ID, expense.Date, expense.Description, expense.Amount)
			}
		} else {
			fmt.Println("No expenses for the requested month")
		}

	},
}
