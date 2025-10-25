package strategies

import "fmt"

type PremiumPricing struct{}

func (premium *PremiumPricing) CalculateRentalCost(days int, basePrice float64) float64 {
	fmt.Println("Calculating cost using Premium Pricing")
	total := basePrice * float64(days)
	if days > 7 {
		total *= 0.9
	}
	return total
}
