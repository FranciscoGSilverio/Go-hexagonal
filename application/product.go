package application

import (
	"errors"
)

type ProductInterface interface {
	IsValid() (bool, error)
	Enable() error
	Disable() error
	GetID() string
	GetName() string
	GetStatus() string
	GetPrice() float64
}

type Product struct {
	ID string
	Name string
	Price float64
	Status string
}

const (
	DISABLED = "disabled"
	ENABLED = "enabled"
)


func IsValid(p *Product) (bool, error) {
	return true, nil

}

func (p *Product) Enable() error{
	if p.Price > 0 {
		p.Status = ENABLED
		return nil
	}

	return errors.New(PriceMustBeGreaterThanZero)
}

func (p *Product) Disable() error{
	if p.Price == 0 {
		p.Status = DISABLED
		return nil
	}

	return errors.New(PriceMustBeZeroToDisable)
}

