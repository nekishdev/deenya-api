package models

import "time"

type InvoiceData struct {
	ID           *int64     `json:"id,omitempty" db:"id"`
	IsPaid       *bool      `json:"is_paid,omitempty" db:"is_paid"`
	Total        *int64     `json:"total,omitempty" db:"total"`
	CreatedAt    *time.Time `json:"created_at,omitempty" db:"created_at"`
	ConsultantID *int64     `json:"consultant_id,omiempty" db:"consultant_id"`
	ClientID     *int64     `json:"client_id,omitempty" db:"client_id"`
	IsBooking    *bool      `json:"is_booking,omitempty" db:"is_booking"`
	IsOrder      *bool      `json:"is_order,omitempty" db:"is_order"`
}

type Invoice struct {
	InvoiceData
	OrderItems []*OrderItem `json:"order_items,omitempty"`
	Consultant *User        `json:"consultant,omitempty"`
	Client     *User        `json:"client,omitempty"`
}
