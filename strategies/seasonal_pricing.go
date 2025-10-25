package strategies

import "fmt"

type SeasonalPricing struct{}

func (seasonal *SeasonalPricing) CalculateRentalCost(days int, basePrice float64) float64 {
	fmt.Println("Calculating cost using Seasonal Pricing")
	return basePrice * float64(days) * 1.2
}
