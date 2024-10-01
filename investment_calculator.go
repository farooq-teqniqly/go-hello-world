package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount float64
	var years float64
	var expectedReturnRate float64

	fmt.Print("Enter investment amount: ")
	_, err := fmt.Scan(&investmentAmount)

	if err != nil {
		return
	}

	fmt.Print("Enter the number of years: ")
	_, err = fmt.Scan(&years)

	if err != nil {
		return
	}

	fmt.Print("Enter the expected rate of return: ")
	_, err = fmt.Scan(&expectedReturnRate)

	if err != nil {
		return
	}

	fv := futureValue(investmentAmount, years, expectedReturnRate)
	inflationAdjustedFV := futureValue(investmentAmount, years, inflationRate)

	fmt.Println("Future value: ", fv)
	fmt.Println("Inflation adjusted future value (", inflationRate, " inflation rate)", inflationAdjustedFV)
}

func futureValue(investmentAmount float64, years float64, rate float64) float64 {
	return investmentAmount * math.Pow(1+(rate/100), years)
}
