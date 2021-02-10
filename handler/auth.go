package handler

import (
	"deenya-api/database"
	"deenya-api/models"
	"deenya-api/services"
	"errors"
	"fmt"
	"net/http"

	"github.com/clarketm/json"

	"github.com/dgrijalva/jwt-go"

	"github.com/go-chi/jwtauth"
)

type JwtToken struct {
	Token string `json:"token"`
}

var TokenAuth = jwtauth.New("HS256", []byte("secret123"), nil)

func Verify(w http.ResponseWriter, r *http.Request) {
	var auth models.AuthCode
	panic(auth)
}

// Register
// @Summary Register
// @Description Register user
// @Tags Auth
// @ID auth-register
// @Accept  json
// @Produce  json
// @Param body body models.UserData true "User Data"
// @Success 200 {object} models.UserData
// @Failure 400 {object} RegisterResponse
// @Router /register [post]

type RegisterRequest struct {
}

type RegisterResponse struct {
	User  models.User `json:"user"`
	Token string      `json:"token"`
}

func Register(w http.ResponseWriter, r *http.Request) {

	var user models.User

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if user.UserBase == (models.UserBase{}) {
		http.Error(w, errors.New("Unable to read input data").Error(), http.StatusBadRequest)
		return
	}

	err := database.NewUser(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = services.NewStripeCustomerForClient(*user.ID)
	if err != nil {
		http.Error(w, fmt.Sprintf("error create a Stripe customer, %s", err.Error()), http.StatusBadRequest)
		return
	}

	_, token, err := TokenAuth.Encode(jwt.MapClaims{
		"id":       *user.ID,
		"username": *user.Username,
		"type":     *user.Type,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := struct {
		User  models.User `json:"user"`
		Token string      `json:"token"`
	}{
		User:  user,
		Token: token,
	}

	js, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//fmt.Println(string(js))
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// Login
// @Summary Login
// @Description Login user
// @Tags Auth
// @ID login-register
// @Accept  json
// @Produce  json
// @Param body body models.LoginData true "Login data"
// @Success 200 {object} JwtToken
// @Failure 400 {object} interface{}
// @Router /login [post]
func Login(w http.ResponseWriter, r *http.Request) {
	input := models.LoginData{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println(input)

	var user models.User
	var err error

	email := false //add email validator

	if email {
		user, err = database.UserByEmail(input.Username)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	if !email {
		user, err = database.UserByUsername(input.Username)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	if *user.Password != input.Password {
		http.Error(w, errors.New("Wrong password").Error(), http.StatusBadRequest)
		return
	}

	_, token, err := TokenAuth.Encode(jwt.MapClaims{
		"id":       *user.ID,
		"username": *user.Username,
		"type":     *user.Type,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := JwtToken{
		Token: token,
	}

	js, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// func Login(w http.ResponseWriter, r *http.Request) {
// 	var credentials models.UserCredentials
// 	TokenAuth = jwtauth.New("HS256", []byte("secret123"), nil)
// 	decoder := json.NewDecoder(r.Body)

// 	if err := decoder.Decode(&credentials); err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	user, err := database.UserByUsername(credentials.Username.String)
// 	if err := decoder.Decode(&credentials); err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}
// 	if user.Password.String == "" || user.Password.String != credentials.Password.String {
// 		http.Error(w, "Incorrect username or password", http.StatusUnauthorized)
// 		return
// 	}

// 	//userdata := database.GetUser(account.Username)
// 	//here is the user data, how should it be stored/returned?

// 	_, token, _ := TokenAuth.Encode(jwt.MapClaims{
// 		"username": *credentials.Username,
// 		"password": *credentials.Password,
// 	})

// 	response := JwtToken{Token: token}

// 	js, err := json.Marshal(response)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// w.Write(js)
// }
