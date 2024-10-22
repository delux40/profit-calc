package main

import "fmt"

func main() {
	revenue := getUserInput("Enter Revenue: ")
	expenses := getUserInput("Enter Expenses: ")
	taxRate := getUserInput("Enter Tax Rate: ")

	earningsBeforeTax, earningsAfterTax, ratio := getAllValues(revenue, expenses, taxRate)

	fmt.Printf("Earnings before tax: %v \n", earningsBeforeTax)
	fmt.Println("Earnings after tax:", earningsAfterTax)
	fmt.Println("Ratio:", ratio)
}

func getUserInput(infoText string) float64 {
	fmt.Print(infoText)
	var userInput float64
	fmt.Scan(&userInput)
	return userInput
}

func getAllValues(revenue, expenses, taxRate float64) (ebt float64, profit float64, ratio float64) {
	ebt = revenue - expenses
	profit = ebt - ((ebt / 100) * taxRate)
	ratio = ebt / profit
	return ebt, profit, ratio
}
