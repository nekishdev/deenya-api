package models

import "time"

type ClinicData struct {
	ID   *int64  `db:"id" json:"id,omitempty"`
	Name *string `db:"name" json:"name,omitempty"`
	//places ID
	Lat      *float64 `db:"lat" json:"lat,omitempty"`
	Lng      *float64 `db:"lon" json:"lng,omitempty"`
	City     *string  `db:"city" json:"city,omitempty"`
	Country  *string  `db:"country" json:"country,omitempty"`
	Street   *string  `db:"street" json:"street,omitempty"`
	Postcode *string  `db:"postcode" json:"postcode,omitempty"`
	AdminIds []*int64 `db:"admin_ids" json:"admin_ids,omitempty"`
}

type Clinic struct {
	ClinicData
	Admins  []*User `json:"admins,omitempty"`
	Members []*User `json:"members,omitempty"`
}

type ClinicMemberData struct {
	ID           *int64     `db:"id" json:"id,omitempty"`
	ConsultantID *int64     `db:"consultant_id" json:"consultant_id,omitempty"`
	ClinicID     *int64     `db:"clinic_id" json:"clinic_id,omitempty"`
	IsAccepted   *bool      `db:"is_accepted" json:"is_accepted,omitempty"` //request or member
	CreatedAt    *time.Time `db:"created_at" json:"created_at,omitempty"`
}

type ClinicMember struct {
	ClinicMemberData
	Consultant *User `json:"consultant,omitempty"`
}

type ClinicList struct {
	Requests    []*ClinicMember
	Consultants []*User
}
