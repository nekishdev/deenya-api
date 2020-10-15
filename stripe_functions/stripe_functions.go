package stripe_functions

import (
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/customer"
	"github.com/stripe/stripe-go/oauth"
)

const (
	STRIPE_TEST_SECRET            = `sk_test_JqR28oyXlK7IDYKvMMi4Be5i00ytnVrnmE`
	STRIPE_LIVE_SECRET            = ``
	STRIPE_DEFAULT_CURRENCY       = "GBP"
	STRIPE_CONNECT_CLIENT_ID_TEST = `ca_Hgt2WXrYld14yrb6zEvCE5x5m3k57eUv`
)

func CreateCustomer() (string, error) {

	stripe.Key = STRIPE_TEST_SECRET

	params := &stripe.CustomerParams{}
	c, err := customer.New(params)

	if err != nil {
		return "", err
	}

	return c.ID, nil

}

func GetAccountAuthToken(code string) (string, error) {
	stripe.Key = STRIPE_TEST_SECRET
	params := &stripe.OAuthTokenParams{
		GrantType: stripe.String("authorization_code"),
		Code:      stripe.String(code),
	}
	token, err := oauth.New(params)
	if err != nil {
		return "", err
	}

	return token.StripeUserID, nil
}
