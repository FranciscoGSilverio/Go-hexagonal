package application_test

import (
	"testing"

	"github.com/FranciscoGSilverio/Go-hexagonal/application"
	"github.com/stretchr/testify/require"
)

func TestProduct_Enable(t *testing.T){
	product := application.Product{}

	product.Name = "Keyboard"
	product.Status = application.DISABLED
	product.Price = 10

	err := product.Enable()
	require.Nil(t, err)

	product.Price = 0
	err = product.Enable()
	require.Equal(t, application.PriceMustBeGreaterThanZero, err.Error())
}

func TestProduct_Disable(t *testing.T){
	product := application.Product{}

	product.Name = "Keyboard"
	product.Status = application.DISABLED
	product.Price = 0

	err := product.Disable()
	require.Nil(t, err)

	product.Price = 10
	err = product.Disable()
	require.Equal(t, application.PriceMustBeZeroToDisable, err.Error())
}