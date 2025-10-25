package strategies

type PricingStrategy interface {
	CalculateRentalCost(days int, basePrice float64) float64
}
