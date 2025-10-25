package main

import (
	"assignment6_Strategy/context"
	"assignment6_Strategy/strategies"
	"fmt"
)

func main() {
	context := context.NewRentalContext(&strategies.StandardPricing{})
	basePrice := 50.0

	cost1 := context.CalculateRental(5, basePrice)
	fmt.Printf("Cost for 5 days (Standard): $%.2f\n", cost1)

	context.SetPricingStrategy(&strategies.PremiumPricing{})

	cost2 := context.CalculateRental(10, basePrice)
	fmt.Printf("Cost for 10 days (Premium): $%.2f\n", cost2)

	context.SetPricingStrategy(&strategies.SeasonalPricing{})

	cost3 := context.CalculateRental(3, basePrice)
	fmt.Printf("Cost for 3 days (Seasonal): $%.2f\n", cost3)
}
