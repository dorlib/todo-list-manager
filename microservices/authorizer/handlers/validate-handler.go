package handlers

import (
	"fmt"
	"net/http"
)

func ValidateHandler(rw http.ResponseWriter, r *http.Request) {
	if _, ok := r.Header["api-key"]; !ok {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("API key Missing"))

		return
	}

	valid, err := validateUser(r.Header["Email"][0], r.Header["Passwordhash"][0])
	if err != nil {
		// this means either the user does not exist
		rw.WriteHeader(http.StatusUnauthorized)
		rw.Write([]byte("User Does not Exist"))

		return
	}

	if !valid {
		rw.WriteHeader(http.StatusUnauthorized)
		rw.Write([]byte("Incorrect Password"))

		return
	}

	tokenString, err := getSignedToken()
	if err != nil {
		fmt.Println(err)
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte("Internal Server Error"))

		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte(tokenString))
}
