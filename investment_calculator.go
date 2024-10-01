package main

import (
	"fmt"
	"log"
	"teqniqly.com/go-investment-calculator/finance"
)

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	getInput("Enter revenue: ", &revenue)
	getInput("Enter expenses: ", &expenses)
	getInput("Enter tax rate: ", &taxRate)

	ensurePositive(revenue, "Revenue")
	ensurePositive(expenses, "Expenses")
	ensurePositive(taxRate, "Tax rate")

	ebt, profit, ratio, warning := finance.GetEarningsMetrics(revenue, expenses, taxRate)

	if warning {
		fmt.Println("Warning: Profit is zero, division by zero avoided in ratio calculation.")
	}

	fmt.Println("EBT: ", ebt)
	fmt.Println("Profit: ", profit)
	fmt.Println("Ratio: ", ratio)
}

func getInput(prompt string, value *float64) {
	fmt.Print(prompt)
	_, err := fmt.Scan(value)

	if err != nil {
		log.Fatal("Error: ", err)
		return
	}
}

func ensurePositive(value float64, param string) {
	if value < 0 {
		log.Fatal("Error:", param, " value must be positive")
	}
}
