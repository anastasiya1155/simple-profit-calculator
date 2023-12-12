package main

import "fmt"

func main() {
	revenue := getUserInput("Input your revenue: ")
	expenses := getUserInput("Input your expenses amount: ")
	tax_rate := getUserInput("Input your tax rate (%): ")

	ebt, profit, ratio := calculateResult(revenue, expenses, tax_rate)

	// fmt.Println("Your earning before tax ", ebt, ", profit ", profit, ", ratio ", ratio)
	fmt.Printf("Your earning before tax %.2f, profit %.2f, ratio %.2f\n", ebt, profit, ratio)
}

func getUserInput(label string) float64 {
	fmt.Print(label)
	var myVar float64
	fmt.Scan(&myVar)

	return myVar
}

func calculateResult(revenue, expenses, tax_rate float64) (ebt float64, profit float64, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - tax_rate/100)
	ratio = ebt / profit

	return ebt, profit, ratio
}
