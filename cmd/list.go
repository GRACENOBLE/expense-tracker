package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/GRACENOBLE/expense-tracker/pkg/types"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listExpenses)
}

var listExpenses = &cobra.Command{
	Use:   "list",
	Short: "Long list all tasks",
	Long:  "This function lists all tasks in details.",
	Run: func(cmd *cobra.Command, args []string) {
		jsonData, err := os.ReadFile("output/expenses.json")
		if err != nil {
			log.Fatal(err)
		}

		var expenses []types.Expense

		err = json.Unmarshal(jsonData, &expenses)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("ID\tDate\t\t\tDescription\tAmount")
		for _, expense := range expenses {
			fmt.Printf("%v\t%v\t%v\t%v\n", expense.ID, expense.Date, expense.Description, expense.Amount)
		}

	},
}
