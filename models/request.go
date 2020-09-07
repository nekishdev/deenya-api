package models

type RegisterReq struct {
	Title      string
	FirstName  string
	MiddleName string
	LastName   string
	Username   string
	Email      string
	Password   string
}

type LoginReq struct {
	Username string
	Email    string
	Password string
}
