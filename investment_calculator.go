package main

import (
	"fmt"
	"log"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount float64
	var years float64
	var expectedReturnRate float64

	getInput("Enter investment amount: ", &investmentAmount)
	getInput("Enter the number of years: ", &years)
	getInput("Enter expected rate of return: ", &expectedReturnRate)

	fv := futureValue(investmentAmount, years, expectedReturnRate)
	inflationAdjustedFV := futureValue(investmentAmount, years, inflationRate)

	fmt.Println("Future value: ", fv)
	fmt.Println("Inflation adjusted future value (", inflationRate, " inflation rate)", inflationAdjustedFV)
}

func getInput(prompt string, value *float64) {
	fmt.Print(prompt)
	_, err := fmt.Scan(value)

	if err != nil {
		log.Fatal("Error: ", err)
		return
	}
}

func futureValue(investmentAmount float64, years float64, rate float64) float64 {
	return investmentAmount * math.Pow(1+(rate/100), years)
}
