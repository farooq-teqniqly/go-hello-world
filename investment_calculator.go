package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	investmentAmount := 1000.0
	years := 10.0
	expectedReturnRate := 5.5

	fv := futureValue(investmentAmount, years, expectedReturnRate)
	inflationAdjustedFV := futureValue(investmentAmount, years, inflationRate)

	fmt.Println(fv)
	fmt.Println(inflationAdjustedFV)
}

func futureValue(investmentAmount float64, years float64, rate float64) float64 {
	return investmentAmount * math.Pow(1+(rate/100), years)
}
