package main

import (
	"fmt"
	"os"
)

func main() {
	revenue := getUserInput("Enter Revenue: ")
	expenses := getUserInput("Enter Expenses: ")
	taxRate := getUserInput("Enter Tax Rate: ")

	earningsBeforeTax, earningsAfterTax, ratio := calculateAllValues(revenue, expenses, taxRate)

	result := fmt.Sprintf("Earnings before tax: %.1f \nEarnings after tax: %.1f \nRatio: %.3f", earningsBeforeTax, earningsAfterTax, ratio)
	fmt.Println(result)
	writeResultToFile(result)
}

func getUserInput(infoText string) float64 {
	fmt.Print(infoText)
	var userInput float64
	fmt.Scan(&userInput)

	if userInput <= 0 {
		panic("You have entered a value greater than or equal to 0.")
	}

	return userInput
}

func calculateAllValues(revenue, expenses, taxRate float64) (ebt float64, profit float64, ratio float64) {
	ebt = revenue - expenses
	profit = ebt - ((ebt / 100) * taxRate)
	ratio = ebt / profit
	return ebt, profit, ratio
}

func writeResultToFile(resultText string) {
	os.WriteFile("Result.txt", []byte(resultText), 0644)
}
