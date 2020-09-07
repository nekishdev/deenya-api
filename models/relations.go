package models

type ConsultantClient struct { //just return orders/bookings instead of relationship table? performance costs?
	ID           *int64 `json:"id,omitempy" db:"id"`
	ConsultantID *int64 `json:"consultant_id,omitempty" db:"consultant_id"`
	ClientID     *int64 `json:"client_id,omitempty" db:"client_id"`
	//CreatedAt    *time.Time `json:"created_at,omitempty" db:"created_at"`

	Orders   []*Order
	Bookings []*Booking

	Client     *User `json:"client"`
	Consultant *User `json:"consultant"`
}
