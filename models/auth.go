package models

type AuthCode struct {
	User *int64
	Code *int64
}

type LoginData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
