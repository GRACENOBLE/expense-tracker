package main

// import (
// 	"bufio"
// 	"fmt"
// 	"github.com/GRACENOBLE/expense-tracker/cmd"
// 	"log"
// 	"os"
// 	"strconv"
// 	"strings"
// )

import "github.com/spf13/cobra"

func main() {

	// reader := bufio.NewReader(os.Stdin)
	// for {

	// 	fmt.Print("expense-tracker:")

	// 	input, _ := reader.ReadString('\n')

	// 	parts := strings.Fields(input)

	// 	switch parts[0] {

	// 	case "add":
	// 		if len(parts) == 3 {

	// 			ammount, err := strconv.ParseInt(parts[2], 10, 64)
	// 			if err != nil {
	// 				log.Fatal(err)
	// 			}
	// 			cmd.AddExpense(parts[1], ammount)

	// 			return

	// 		} 

	// 		fmt.Println("Invalid command-line Arguments")
			
	// 	case "list":

	// 		cmd.ListExpenses()

	// 	case "summary":

	// 		cmd.SummarizeExpenses()

	// 	case "delete":
	// 		// cmd.DeleteExpense()
	// 	default:
	// 		fmt.Println("Command not recognized, try \"help\"")
	// 	}
	// }
}
