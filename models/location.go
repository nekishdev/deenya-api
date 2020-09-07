package models

type Location struct {
	PlacesID int64 //google places id

	Name          string
	Lat           float64
	Lon           float64
	StreetNumber  *int64
	StreetAddress string
	Area          string
	City          string
	State         string
	Country       string
	Postcode      string
}
