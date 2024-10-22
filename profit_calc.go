package main

import "fmt"

func main() {
	revenue := getUserInput("Enter Revenue: ")
	expenses := getUserInput("Enter Expenses: ")
	taxRate := getUserInput("Enter Tax Rate: ")

	earningsBeforeTax, earningsAfterTax, ratio := calculateAllValues(revenue, expenses, taxRate)

	fmt.Printf("Earnings before tax: %.1f \n", earningsBeforeTax)
	fmt.Printf("Earnings after tax: %.1f \n", earningsAfterTax)
	fmt.Printf("Ratio: %.3f", ratio)
}

func getUserInput(infoText string) float64 {
	fmt.Print(infoText)
	var userInput float64
	fmt.Scan(&userInput)
	return userInput
}

func calculateAllValues(revenue, expenses, taxRate float64) (ebt float64, profit float64, ratio float64) {
	ebt = revenue - expenses
	profit = ebt - ((ebt / 100) * taxRate)
	ratio = ebt / profit
	return ebt, profit, ratio
}
