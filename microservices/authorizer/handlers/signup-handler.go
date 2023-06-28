package handlers

import (
	"net/http"
)

// SignupHandler adds the user to the database of users.
func SignupHandler(rw http.ResponseWriter, r *http.Request) {
	if _, ok := r.Header["Email"]; !ok {
		rw.WriteHeader(http.StatusBadRequest)
		_, err := rw.Write([]byte("Email Missing"))
		if err != nil {
			return
		}

		return
	}

	if _, ok := r.Header["Username"]; !ok {
		rw.WriteHeader(http.StatusBadRequest)
		_, err := rw.Write([]byte("Username Missing"))
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

	if _, ok := r.Header["Fullname"]; !ok {
		rw.WriteHeader(http.StatusBadRequest)
		_, err := rw.Write([]byte("Fullname Missing"))
		if err != nil {
			return
		}

		return
	}

	// validate and then add the user
	_, err := rw.Write([]byte(r.Header["Passwordhash"][0]))
	if err != nil {
		return
	}

	rw.WriteHeader(http.StatusOK)
	_, err := rw.Write([]byte("User Created"))
	if err != nil {
		return
	}
}
