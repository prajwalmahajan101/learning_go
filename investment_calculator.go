package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5

	var investmentAmount, years, expectedReturnRate float64

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Expected ReturnRate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years, inflationRate)

	formattedFV := fmt.Sprintf("Future value: %.2f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future real value: %.2f \n", futureRealValue)
	fmt.Print(formattedFV, formattedRFV)
}

func calculateFutureValues(investmentAmount float64, expectedReturnRate, years, inflationRate float64) (futureValue float64, futureRealValue float64) {
	futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue = futureValue / math.Pow(1+inflationRate/100, years)
	return
}
