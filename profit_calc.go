package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Print("Enter Revenue: ")
	fmt.Scan(&revenue)
	fmt.Print("Enter Expenses: ")
	fmt.Scan(&expenses)
	fmt.Print("Enter Tax Rate: ")
	fmt.Scan(&taxRate)

	ebt := revenue - expenses
	profit := ebt - ((ebt / 100) * taxRate)
	ratio := ebt / profit

	fmt.Printf("Earnings before tax: %v \n", ebt)
	fmt.Println("Earnings after tax:", profit)
	fmt.Println("Ratio:", ratio)
}