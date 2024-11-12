package helpers

import (
	"fmt"
	"time"

	"github.com/GRACENOBLE/expense-tracker/pkg/types"
)

// Factory to create a new expense
func CreateNewExpense(id int, description string, ammount int64) types.Expense {
	return types.Expense{
		ID:          int64(id),
		Date:        time.Now().Format("2006-01-02 15:04:05"),
		Description: description,
		Amount:      fmt.Sprintf("$%v", ammount),
	}
}
