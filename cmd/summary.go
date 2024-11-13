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
	rootCmd.AddCommand(sumariseExpenses)
}

var sumariseExpenses = &cobra.Command{
	Use:   "sumarize",
	Short: "Sumarize all tasks",
	Long:  "This lists all tasks in a more human readable way.",
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

		var total int64

		for _, expense := range expenses {

			total += expense.Amount

		}

		fmt.Printf("Summary : %v\n", total)
	},
}
