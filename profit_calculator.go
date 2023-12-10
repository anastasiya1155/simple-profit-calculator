package main

import "fmt"

func main() {
	var revenue, expenses float64
	var tax_rate int
	fmt.Print("Input your revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Input your expenses amount: ")
	fmt.Scan(&expenses)

	fmt.Print("Input your tax rate (%): ")
	fmt.Scan(&tax_rate)

	var ebt = revenue - expenses
	var profit = ebt * (1 - float64(tax_rate)/100)
	var ratio = ebt / profit

	// fmt.Println("Your earning before tax ", ebt, ", profit ", profit, ", ratio ", ratio)
	fmt.Printf("Your earning before tax %.2f, profit %.2f, ratio %.2f\n", ebt, profit, ratio)
}
