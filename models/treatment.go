package models

type TreatmentData struct {
	ID           *int64    `db:"id" json:"id,omitempty"`
	ConsultantID *int64    `db:"consultant_id" json:"consultant_id,omitempty"`
	ClientID     *int64    `db:"client_id" json:"client_id,omitempty"`
	OrderID      *int64    `db:"order_id" json:"order_id,omitempty"`
	PortfolioID  *int64    `db:"portfolio_id" json:"portfolio_id,omitempty"`
	Notes        *string   `db:"notes" json:"notes,omitempty"`
	Tags         []*string `db:"tags" json:"tags,omitempty"`
	CreatedAt    *int64    `db:"created_at" json:"created_at"`
}

type Treatment struct {
	TreatmentData

	Order Order `json:"order,omitempty"`
	//Portfolio  Portfolio `json:"portfolio,omitempty"`
	Consultant *User `json:"consultant,omitempty"`
	Client     *User `json:"client,omitempty"`
}
