package context

import (
	"assignment6_Strategy/pricing_strategy"
	"fmt"
)

type RentalContext struct {
	pricingStrategy pricing_strategy.PricingStrategy
}

func NewRentalContext(strategy pricing_strategy.PricingStrategy) *RentalContext {
	return &RentalContext{pricingStrategy: strategy}
}

func (rentalcontext *RentalContext) SetPricingStrategy(strategy pricing_strategy.PricingStrategy) {
	rentalcontext.pricingStrategy = strategy
}

func (rentalcontext *RentalContext) CalculateRental(days int, basePrice float64) float64 {
	if rentalcontext.pricingStrategy == nil {
		fmt.Println("No pricing strategy selected!")
		return 0
	}
	return rentalcontext.pricingStrategy.CalculateRentalCost(days, basePrice)
}
