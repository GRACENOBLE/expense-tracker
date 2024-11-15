package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/GRACENOBLE/expense-tracker/pkg/helpers"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(deleteExpense)
}

var deleteExpense = &cobra.Command{
	Use:   "delete",
	Short: "Delete a task",
	Long:  "This function takes a command line argument as an integer and deletes corresponding task with the same ID.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			index, err := strconv.ParseInt(args[0], 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			expenses, err := helpers.ReadExpensesFromFile("output/expenses.json")
			if err != nil {
				log.Fatal(err)
			}
			for _, expense := range expenses {
				if expense.ID == index {
					expenses = append(expenses[:index-1], expenses[index:]...)
				}
			}
			for i := range expenses { //interesting how loops pass the iterated object as a copy but when you use indexing the actual slice is mutated
				expenses[i].ID = int64(i + 1)
			}
			for _, expense := range expenses {
				fmt.Println(expense)
			}
			helpers.WriteExpensesToFile("output/expenses.json",expenses)
			fmt.Println("Deleted expense")
		} else {
			log.Fatal("invalid Arguments")
		}
	},
}
