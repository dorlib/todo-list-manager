package handlers

import (
	"fmt"
	"net/http"
)

func ValidateHandler(rw http.ResponseWriter, r *http.Request) {
	if _, ok := r.Header["api-key"]; !ok {
		rw.WriteHeader(http.StatusBadRequest)
		_, err := rw.Write([]byte("API key Missing"))
		if err != nil {
			return
		}

		return
	}

	valid, err := validateUser(r.Header["Email"][0], r.Header["Passwordhash"][0])
	if err != nil {
		// this means either the user does not exist
		rw.WriteHeader(http.StatusUnauthorized)
		_, err := rw.Write([]byte("User Does not Exist"))
		if err != nil {
			return
		}

		return
	}

	if !valid {
		rw.WriteHeader(http.StatusUnauthorized)
		_, err := rw.Write([]byte("Incorrect Password"))
		if err != nil {
			return
		}

		return
	}

	tokenString, err := getSignedToken()
	if err != nil {
		fmt.Println(err)
		rw.WriteHeader(http.StatusInternalServerError)
		_, err := rw.Write([]byte("Internal Server Error"))
		if err != nil {
			return
		}

		return
	}

	rw.WriteHeader(http.StatusOK)
	_, err = rw.Write([]byte(tokenString))
	if err != nil {
		return
	}
}
