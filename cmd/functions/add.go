package functions

import (
	"fmt"
	"log"

	"github.com/GRACENOBLE/expense-tracker/cmd/helpers"
	"github.com/GRACENOBLE/expense-tracker/cmd/types"
)

func AddExpense(expense types.Expense) {
	// Ensure the file exists
	helpers.CreateFileIfNotExists("expenses.json")


	// Check if the file is empty
	empty, err := helpers.IsFileEmpty("expenses.json")
	if err != nil {
		log.Fatalf("Failed to check if file is empty: %v", err)
	}

	// If the file is empty, initialize with the new expense
	if empty {
		if err := helpers.WriteExpensesToFile("expenses.json", []types.Expense{}); err != nil {
			log.Fatalf("Failed to write new expense: %v", err)
		}
	} else {
		// Otherwise, append the new expense to existing data
		existingData, err := helpers.ReadExpensesFromFile("/output/expenses.json")
		if err != nil {
			log.Fatalf("Failed to read existing expenses: %v", err)
		}

		// Append the new expense
		newData := append(existingData, expense)

		// Write the updated data back to the file
		if err := helpers.WriteExpensesToFile("/output/expenses.json", newData); err != nil {
			log.Fatalf("Failed to write updated expenses: %v", err)
		}
	}

	fmt.Println("Added Successfully.")
}