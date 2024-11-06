package main

import (
	"bufio"
	"fmt"
	"time"

	"os"
	"strings"

	"github.com/GRACENOBLE/expense-tracker/functions"
	"github.com/GRACENOBLE/expense-tracker/types"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	for {
		demo := types.Expense{
			ID:          1,
			Date:        time.Now().Format("2006-01-02 15:04:05"),
			Description: "lunch",
			Amount:      20,
		}

		fmt.Print("expense-tracker:")

		input, _ := reader.ReadString('\n')

		parts := strings.Fields(input)

		switch parts[0] {

		case "add":
			functions.AddExpense(demo)
		case "list":
			functions.ListExpenses()
		case "summary":
			functions.SummarizeExpenses()
		case "delete":
			functions.DeleteExpense(demo)
		default:
			fmt.Println("Command not recognized, try \"help\"")
		}
	}
}
