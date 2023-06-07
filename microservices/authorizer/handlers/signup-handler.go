package handlers

import (
	"net/http"
)

// SignupHandler adds the user to the database of users.
func SignupHandler(rw http.ResponseWriter, r *http.Request) {
	if _, ok := r.Header["Email"]; !ok {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Email Missing"))

		return
	}

	if _, ok := r.Header["Username"]; !ok {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Username Missing"))

		return
	}

	if _, ok := r.Header["Passwordhash"]; !ok {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Passwordhash Missing"))

		return
	}

	if _, ok := r.Header["Fullname"]; !ok {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Fullname Missing"))

		return
	}

	// validate and then add the user
	rw.Write([]byte(r.Header["Passwordhash"][0]))

	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte("User Created"))
}
