package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/GRACENOBLE/expense-tracker/pkg/types"
)

func DeleteExpense(index int64){

	jsonData, err := os.ReadFile("output/expenses.json")
	if err != nil {
		log.Fatal(err)
	}

	var expenses []types.Expense
	
	err = json.Unmarshal(jsonData, &expenses)
	if err != nil{
		log.Fatal(err)
	}

	for _, expense := range expenses{
		
		fmt.Printf("%v", expense)
	}

}