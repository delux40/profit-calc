package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	inputText("Enter Revenue: ")
	fmt.Scan(&revenue)
	inputText("Enter Expenses: ")
	fmt.Scan(&expenses)
	inputText("Enter Tax Rate: ")
	fmt.Scan(&taxRate)

	earningsBeforeTax, earningsAfterTax, ratio := getAllValues(revenue, expenses, taxRate)

	fmt.Printf("Earnings before tax: %v \n", earningsBeforeTax)
	fmt.Println("Earnings after tax:", earningsAfterTax)
	fmt.Println("Ratio:", ratio)
}

func inputText(text string) {
	fmt.Print(text)
}

func getAllValues(revenue, expenses, taxRate float64) (ebt float64, profit float64, ratio float64) {
	ebt = revenue - expenses
	profit = ebt - ((ebt / 100) * taxRate)
	ratio = ebt / profit
	return ebt, profit, ratio
}
