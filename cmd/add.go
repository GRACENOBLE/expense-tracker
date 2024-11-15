package cmd

import (
	"fmt"
	"log"
	"strconv"
	"github.com/GRACENOBLE/expense-tracker/pkg/helpers"
	"github.com/GRACENOBLE/expense-tracker/pkg/types"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(addExpense)
}

var addExpense = &cobra.Command{
	Use:   "add",
	Short: "Add an expense to the expense-tracker",
	Long:  "This command takes in command line arguments and uses them to create an ",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			log.Fatal("Invalid number of arguents")
		}

		expenseBody := args[0]

		expenseAmount, err := strconv.ParseInt(args[1], 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		//Ensure the file exists
		helpers.CreateFileIfNotExists("output/expenses.json")

		// Check if the file is empty
		empty, err := helpers.IsFileEmpty("output/expenses.json")
		if err != nil {
			log.Fatalf("Failed to check if file is empty: %v", err)
		}

		// If the file is empty, initialize with the new expense
		if empty {
			//create a new expense
			expense := helpers.CreateNewExpense(1, expenseBody, expenseAmount)
			if err := helpers.WriteExpensesToFile("output/expenses.json", []types.Expense{expense}); err != nil {
				log.Fatalf("Failed to write new expense: %v", err)
			}
		} else {
			//count expenses
			expenses, err := helpers.ReadExpensesFromFile("output/expenses.json")
			if err != nil {
				log.Fatal(err)
			}

			expenseCount := len(expenses) + 1

			//create new expeanse
			expense := helpers.CreateNewExpense(expenseCount, expenseBody, expenseAmount)

			// Append the new expense
			newData := append(expenses, expense)

			// Write the updated data back to the file
			if err := helpers.WriteExpensesToFile("output/expenses.json", newData); err != nil {
				log.Fatalf("Failed to write updated expenses: %v", err)
			}
		}

		fmt.Println("Added Successfully.")
	},
}
