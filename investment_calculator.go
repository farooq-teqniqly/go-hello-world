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

	fmt.Println("EBT: ", formatCurrency(ebt))
	fmt.Println("Profit: ", formatCurrency(profit))
	fmt.Println("Ratio: ", formatFloat(ratio))
}

func formatFloat(value float64) string {
	return fmt.Sprintf("%.2f", value)
}

func formatCurrency(value float64) string {
	return fmt.Sprintf("$%.2f", value)
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
