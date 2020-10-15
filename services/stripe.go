package services

import (
	"deenya-api/database"
	"deenya-api/models"
	"deenya-api/stripe_functions"
	"fmt"
)

func NewStripeCustomerForClient(clientID int64) (models.StripeCustomer, error) {
	var customer models.StripeCustomer
	stripeCustomerID, err := stripe_functions.CreateCustomer()
	if err != nil {
		return customer, fmt.Errorf("failed to create customer on Stripe server, %s", err.Error())
	}

	customer = models.StripeCustomer{
		ClientID:      &clientID,
		CustomerToken: &stripeCustomerID,
	}
	err = database.NewCustomer(&customer)
	if err != nil {
		return customer, err
	}

	return customer, nil
}
