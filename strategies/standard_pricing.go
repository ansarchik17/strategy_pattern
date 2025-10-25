package strategies

import "fmt"

type StandardPricing struct{}

func (standard *StandardPricing) CalculateRentalCost(days int, basePrice float64) float64 {
	fmt.Println("Calculating cost using Standard Pricing")
	return basePrice * float64(days)
}
