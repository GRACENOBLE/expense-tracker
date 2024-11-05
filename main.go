package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	for{

		fmt.Print("expense-tracker:")
		
		input, _ := reader.ReadString('\n')
 
		parts := strings.Fields(input)
		
		switch parts[0]{

		case "add":
			
		case "list":
		case "summary":
		case "delete":
			
		}
	}
}