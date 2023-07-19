package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIfItGetsAnErrorIfIDIsBlank(t *testing.T) {
	order := Order{}

	assert.Error(t, order.Validate(), "Id is required")

}

func TestIfItGetsAnErrorIfPriceIsBlank(t *testing.T) {
	order := Order{ID: "123"}

	assert.Error(t, order.Validate(), "Price must be greater than 0")
}

func TestIfItGetsAnErrorIfTaxIsBlank(t *testing.T) {
	order := Order{ID: "123", Price: 100}

	assert.Error(t, order.Validate(), "Tax must be greater than 0")
}

func TestFinalPrice(t *testing.T) {
	order := Order{ID: "123", Price: 100.0, Tax: 10.0}

	assert.NoError(t, order.Validate())
	assert.Equal(t, "123", order.ID)
	assert.Equal(t, 100.0, order.Price)
	assert.Equal(t, 10.0, order.Tax)

	order.CalculateFinalPrice()
	assert.Equal(t, 110.0, order.FinalPrice)

}
