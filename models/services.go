package models

type Service struct {
	ID          int64
	OwnerID     int64
	Name        string
	Cost        int64 //cents
	Length      int64 //minutes
	Description string
}
