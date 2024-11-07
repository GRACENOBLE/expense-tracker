package cmd

import (
	"fmt"
	"log"
	"github.com/GRACENOBLE/expense-tracker/pkg/helpers"
	"github.com/GRACENOBLE/expense-tracker/pkg/types"
)

func AddExpense(expenseBody string, expenseAmount int64) {


	// Ensure the file exists
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
		if err != nil{
			log.Fatal(err)
		}

		expenseCount := len(expenses)+1

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
}