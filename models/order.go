package models

import "time"

type OrderData struct {
	ID           *int64 `db:"id" json:"id,omitempty"`
	ClientID     *int64 `db:"client_id" json:"client_id,omitempty"`
	ConsultantID *int64 `db:"consultant_id" json:"consultant_id,omitempty"`
	//InvoiceID *int64     `db:"invoice_id" json:"invoice_id,omitempty"` //create invoice on order creation
	Discount  *int64     `db:"discount" json:"discount,omitempty"`
	Total     *int64     `db:"total" json:"total,omitempty"`
	Status    *string    `db:"status" json:"status,omitempty"`
	CreatedAt *time.Time `db:"created_at" json:"created_at,omitempty"`
}

type Order struct {
	OrderData
	OrderItems []OrderItem    `json:"order_items,omitempty"`
	Delivery   *OrderDelivery `json:"order_delivery,omitempty"`
	Invoice    *Invoice       `json:"invoice,omitempty"`
	Client     *User          `json:"client,omitempty"`
	Consultant *User          `json:"consultant,omitempty"`
}

//https://github.com/clarketm/json - If omitempty is applied to a struct and all the children of the struct are empty, then on marshalling it will be omitted from the encoded json.

type OrderItemData struct {
	ID        *int64 `db:"id" json:"id,omitempty"`
	OrderID   *int64 `db:"order_id" json:"order_id,omitempty"`
	ProductID *int64 `db:"product_id" json:"product_id,omitempty"`
	//InvoiceID *int64 `db:"invoice_id" json:"invoice_id,omitempty"` //update in postgres table? multiple invoices/consultants per order? Or only one consultant per order?
	//allows external embeddable flow?
	Quantity *int64 `db:"quantity" json:"quantity,omitempty"`
	//SellerID *int64 `json:"seller_id,omitempty"`
}

type OrderItem struct {
	OrderItemData
	Product *Product `json:"product,omitempty"`
	//Seller  *User    `json:"seller,omitempty"`
}

type OrderDeliveryData struct {
	ID           *int64     `db:"id" json:"id,omitempty"` //delivery_id, merge with Order as anonymous struct?
	OrderID      *int64     `db:"order_id" json:"order_id,omitempty"`
	Status       *string    `db:"status" json:"status,omitempty"`
	TrackingCode *string    `db:"tracking_code" json:"tracking_code,omitempty"`
	ShippedAt    *time.Time `db:"shipped_at" json:"shipped_at,omitempty"`
	// IsDelivered  *bool      `db:"is_delivered" json:"is_delivered,omitempty"`
}

type OrderDelivery struct {
	OrderDeliveryData

	Order *Order
}
