package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 6.5
	var investmentAmount float64
	var years float64
	expectedReturnRate := 5.5

	// fmt.Print("Investment Amount: ")
	outputText("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Investment Amount: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Investment Amount: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Print(futureRealValue)

}

func outputText(text string) {
	fmt.Print(text)
}
