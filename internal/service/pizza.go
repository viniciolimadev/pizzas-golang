package service

import (
	"errors"
	"pizzaria/internal/models"
)

func ValidatePizzaPrice(pizza *models.Pizza) error {
	if pizza.Preco < 0 {
		return errors.New("pizza price must be positive")
	}
	return nil
}
