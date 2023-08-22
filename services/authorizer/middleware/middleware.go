package middleware

import (
	token2 "authorizer/token"
	"net/http"
)

// tokenValidationMiddleware Middleware itself returns a function that is a Handler. it is executed for each request.
// We want all our routes for REST to be authenticated. So, we validate the token.
func tokenValidationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		// check if token is present.
		if _, ok := r.Header["Token"]; !ok {
			rw.WriteHeader(http.StatusUnauthorized)
			_, err := rw.Write([]byte("Token Missing"))
			if err != nil {
				return
			}

			return
		}

		token := r.Header["Token"][0]

		check, err := token2.ValidateToken(token, "S0m3_R4n90m_sss")
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			_, err := rw.Write([]byte("Token Validation Failed"))
			if err != nil {
				return
			}

			return
		}

		if !check {
			rw.WriteHeader(http.StatusUnauthorized)
			_, err := rw.Write([]byte("Token Invalid"))
			if err != nil {
				return
			}

			return
		}

		rw.WriteHeader(http.StatusOK)
		_, err = rw.Write([]byte("Authorized Token"))
		if err != nil {
			return
		}
	})
}
