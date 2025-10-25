package context

import "assignment6_Strategy/strategies"

type RentalContext struct {
	pricingStrategy strategies.PricingStrategy
}

func NewRentalContext(strategy strategies.PricingStrategy) *RentalContext {
	return &RentalContext{pricingStrategy: strategy}
}

func (rentalcontext *RentalContext) SetPricingStrategy(strategy strategies.PricingStrategy) {
	rentalcontext.pricingStrategy = strategy
}

func (rentalcontext *RentalContext) CalculateRental(days int, basePrice float64) float64 {
	return rentalcontext.pricingStrategy.CalculateRentalCost(days, basePrice)
}
