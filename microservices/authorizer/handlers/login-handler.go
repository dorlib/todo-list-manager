package handlers

import (
	"authorizer/token"
	"fmt"
	"net/http"
	"time"
)

func getSignedToken() (string, error) {
	claimsMap := map[string]string{
		"aud": "frontend.knowsearch.ml",
		"iss": "knowsearch.ml",
		"exp": fmt.Sprint(time.Now().Add(time.Minute * 1).Unix()),
	}

	// should be passed as a System Environment variable.
	secret := "S0m3_R4n90m_sss"
	header := "HS256"
	tokenString, err := token.GenerateToken(header, claimsMap, secret)
	if err != nil {
		return tokenString, err
	}

	return tokenString, nil
}

func SigninHandler(rw http.ResponseWriter, r *http.Request) {
	if _, ok := r.Header["Email"]; !ok {
		rw.WriteHeader(http.StatusBadRequest)
		_, err := rw.Write([]byte("Email Missing"))
		if err != nil {
			return
		}

		return
	}

	if _, ok := r.Header["Passwordhash"]; !ok {
		rw.WriteHeader(http.StatusBadRequest)
		_, err := rw.Write([]byte("Passwordhash Missing"))
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
