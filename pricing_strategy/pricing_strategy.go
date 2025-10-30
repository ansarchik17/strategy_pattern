package pricing_strategy

type PricingStrategy interface {
	CalculateRentalCost(days int, basePrice float64) float64
}
