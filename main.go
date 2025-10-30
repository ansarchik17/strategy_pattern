package main

import (
	"assignment6_Strategy/context"
	"assignment6_Strategy/strategies"
	"fmt"
)

func main() {
	var choice int
	basePrice := 50.0
	days := 0

	fmt.Println("Choose a Pricing Strategy:")
	fmt.Println("1. Standard Pricing")
	fmt.Println("2. Premium Pricing")
	fmt.Println("3. Seasonal Pricing")
	fmt.Print("Enter choice (1-3): ")
	fmt.Scan(&choice)

	fmt.Print("Enter number of rental days: ")
	fmt.Scan(&days)

	var rentalContext *context.RentalContext

	switch choice {
	case 1:
		rentalContext = context.NewRentalContext(&strategies.StandardPricing{})
	case 2:
		rentalContext = context.NewRentalContext(&strategies.PremiumPricing{})
	case 3:
		rentalContext = context.NewRentalContext(&strategies.SeasonalPricing{})
	default:
		rentalContext = context.NewRentalContext(nil)
	}

	total := rentalContext.CalculateRental(days, basePrice)
	fmt.Printf("\nTotal rental cost: $%.2f\n", total)
}
