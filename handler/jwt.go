package handler

import (
	"fmt"
	"net/http"

	"github.com/go-chi/jwtauth"
)

func GetAuthUsername(r *http.Request) string {
	_, claims, _ := jwtauth.FromContext(r.Context())
	data := fmt.Sprintf("%v", claims["username"])

	return data
}

func GetAuthType(r *http.Request) string {
	_, claims, _ := jwtauth.FromContext(r.Context())
	data := fmt.Sprintf("%v", claims["type"])

	return data
}

func GetAuthID(r *http.Request) int64 {
	_, claims, err := jwtauth.FromContext(r.Context())
	if err != nil {
		fmt.Println(err)
	}

	id := claims["id"]
	if id == nil {

	}
	//add error handling
	data := int64(id.(float64))

	return data
}
