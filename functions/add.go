package functions

import (
	
	"fmt"
	"log"
	"github.com/GRACENOBLE/expense-tracker/types"

)

func AddExpense(expense types.Expense) {
	// Ensure the file exists
	CreateFileIfNotExists("expenses.json")


	// Check if the file is empty
	empty, err := IsFileEmpty("expenses.json")
	if err != nil {
		log.Fatalf("Failed to check if file is empty: %v", err)
	}

	// If the file is empty, initialize with the new expense
	if empty {
		if err := writeExpensesToFile("expenses.json", []types.Expense{expense}); err != nil {
			log.Fatalf("Failed to write new expense: %v", err)
		}
	} else {
		// Otherwise, append the new expense to existing data
		existingData, err := readExpensesFromFile("expenses.json")
		if err != nil {
			log.Fatalf("Failed to read existing expenses: %v", err)
		}

		// Append the new expense
		newData := append(existingData, expense)

		// Write the updated data back to the file
		if err := writeExpensesToFile("expenses.json", newData); err != nil {
			log.Fatalf("Failed to write updated expenses: %v", err)
		}
	}

	fmt.Println("Added Successfully.")
}
