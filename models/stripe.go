package models

type StripeConnect struct {
	AccountToken *string `db:"stripe_account_id" json:"stripe_account_id"`
	ConsultantID *int64  `db:"consultant_id" json:"consultant_id"`

	Consultant *User `json:"-"`
}

type StripeAddress struct {
	Line1      string
	Line2      string
	City       string
	Country    string
	State      string
	PostalCode string
}

type StripeMetadata struct {
	//UserType string //client or consultant
}

type StripeCustomer struct {
	ID            *int64
	ClientID      *int64  `db:"client_id" json:"client_id"`
	ConsultantID  *int64  `db:"consultant_id" json:"consultant_id"`
	CustomerToken *string `db:"stripe_customer_id" json:"stripe_customer_id"`

	DefaultSource   *string       `db:"default_source" json:"default_source"`
	Currency        *string       `db:"currency" json:"currency"`
	BillingEmail    *string       `db:"billing_email" json:"billing_email"`
	BillingAddress  StripeAddress `json:"billing_address,omitempty"`
	ShippingAddress StripeAddress `json:"shipping_address,omitempty"`

	Client     *User `json:"-"`
	Consultant *User `json:"-"`
}

type StripeSource struct {
	ID          *int64  `db:"id" json:"id"`
	UserID      *int64  `db:"user_id" json:"user_id"`
	SourceToken *string `db:"stripe_source_id" json:"stripe_source_id"`

	Type            *string `db:"type" json:"type"`       //card or bank account
	Funding         *string `db:"funding" json:"funding"` //debit/credit
	Last4Digits     *string `db:"last4" json:"last4"`
	CardBrand       *string `db:"card_brand" json:"card_brand"`
	ExpirationMonth *int64  `db:"exp_month" json:"exp_month"`
	ExpirationYear  *int64  `db:"exp_year" json:"exp_year"`
}

type StripeCharge struct {
	ID           *int64  `db:"id" json:"id"`
	ConsultantID *int64  `db:"consultant_id" json:"consultant_id"`
	ClientID     *int64  `db:"client_id" json:"client_id"`
	SourceID     *int64  `db:"stripe_source_id" json:"stripe_source_id"`
	ChargeToken  *string `db:"stripe_charge_id" json:"stripe_charge_id"`
	Amount       *int64  `db:"amount" json:"amount"`
}
