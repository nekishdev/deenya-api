package models

import "time"

// type User struct {
// 	ID         *int64
// 	Created    *time.Time
// 	Type       *string
// 	Email      *string
// 	Username   *string `json:"username,omitempty" db:"username"`
// 	Password   *string
// 	FirstName  *string
// 	LastName   *string
// 	MiddleName *string
// 	Title      *string
// 	Country    *string
// 	City       *string
// }

type UserData struct {
	UserBase
	UserDetails
	UserConsultant
	UserFinance
}

type User struct {
	UserData
	Analytics      *UserAnalytics   `json:"analytics,omitempty"`
	Bookings       []*Booking       `json:"booking,omitempty"`
	Products       []*Product       `json:"product,omitempty"`
	Portfolios     []*Portfolio     `json:"portfolio,omitempty"`
	Posts          []*Post          `json:"post,omitempty"`
	Questionnaires []*Questionnaire `json:"questionnaire,omitempty"`
	Orders         []*Order         `json:"order,omitempty"`
	Clinic         *Clinic          `json:"clinic,omitempty"`
}

type UserAnalytics struct { //query results
	NextBooking *Booking `json:"next_booking,omitempty"`
	TotalSpent  *int64   `json:"total_spent,omitempty"`
}

type UserBase struct {
	ID        *int64     `json:"id,omitempty" db:"id" readonly:"true"`
	Email     *string    `json:"email,omitempty" db:"email"`
	Username  *string    `json:"username,omitempty" db:"username"`
	Password  *string    `json:"password" db:"password"`
	CreatedAt *time.Time `json:"created_at,omitempty" db:"created_at" readonly:"true"`
	Type      *string    `json:"type,omitempty" db:"type"`
}

type UserDetails struct {
	UserID           *int64    `json:"-" db:"user_id"`
	ProfilePictureID *int64    `json:"profile_picture_id,omitempty" db:"profile_picture_id"`
	FirstName        *string   `json:"first_name,omitempty" db:"first_name"`
	LastName         *string   `json:"last_name,omitempty" db:"last_name"`
	Title            *string   `json:"title,omitempty" db:"title"`
	Gender           *string   `json:"gender,omitempty" db:"gender"`
	Country          *string   `json:"country" db:"country"`
	City             *string   `json:"city,omitempty" db:"city"`
	Tags             []*string `json:"tags,omitempty" db:"tags"`         //for identifying user skin types etc for matching with correct products and questionnaires etc
	Timezone         *string   `json:"timezone,omitempty" db:"timezone"` //or store in user consultant
	ProfilePicture   *Media    `json:"profile_picture,omitempty" swaggerignore:"true"`
}

type UserContact struct {
	Email   *string
	Mobile  *string
	Address *string
}

type UserFinance struct {
	StripeAccountID string `json:"stripe_account_id,omitempty"`
}

//add sign up questionnaire to gather details about user in order to provide relevant content/product suggestions
type UserSkin struct {
	Type []*string
	Tags []*string
}

type UserConsultant struct {
	ConsultantID  *int64  `db:"consultant_id" json:"-"`
	AvailableFrom *int    `db:"available_from" json:"available_from,omitempty"`
	AvailableTo   *int    `db:"available_to" json:"available_to,omitempty"`
	ClinicID      *int64  `db:"clinic_id" json:"clinic_id,omitempty"`
	Clinic        *Clinic `json:"clinic,omitempty" swaggerignore:"true"`
}

//get timezone from client
//send
//get available bookings
