package helpers

import (
	"encoding/json"
	"fmt"
	"os"
	"github.com/GRACENOBLE/expense-tracker/pkg/types"
)

func CreateFileIfNotExists(filename string){

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		os.Create(filename)
	}

}

func IsFileEmpty(filename string) (bool, error) {

	fileInfo, err := os.Stat(filename)
	if err != nil {
		return false, err
	}

	return fileInfo.Size() == 0, nil

}

// Helper function to write expenses to file
func WriteExpensesToFile(filename string, expenses []types.Expense) error {
	jsonData, err := json.MarshalIndent(expenses, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}

	return os.WriteFile(filename, jsonData, 0644)
}

// Helper function to read expenses from file
func ReadExpensesFromFile(filename string) ([]types.Expense, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	var expenses []types.Expense
	if err := json.Unmarshal(data, &expenses); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	return expenses, nil
}