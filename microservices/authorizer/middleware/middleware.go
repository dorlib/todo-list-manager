package middleware

import (
	"fmt"
	"net/http"

	"go.uber.org/zap"
)

// TokenMiddleware is the token validation route handler
type TokenMiddleware struct {
	logger *zap.Logger
}

// NewTokenMiddleware returns a frsh Token controller
func NewTokenMiddleware(logger *zap.Logger) *TokenMiddleware {
	return &TokenMiddleware{
		logger: logger,
	}
}

// TokenValidationMiddleware Middleware itself returns a function that is a Handler. it is executed for each request.
// We want all our routes for REST to be authenticated. So, we validate the token
func (ctrl *TokenMiddleware) TokenValidationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		// check if token is present
		if _, ok := r.Header["Token"]; !ok {
			ctrl.logger.Warn("Token was not found in the header")
			rw.WriteHeader(http.StatusUnauthorized)
			rw.Write([]byte("Token Missing"))
			return
		}
		token := r.Header["Token"][0]

		secret := token.GetSecret()
		if secret == "" {
			ctrl.logger.Error("Empty JWT secret")
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write([]byte("Internal Server Error"))
			return
		}

		err := token.ValidateToken(token, secret)
		if err != nil {
			errInString := fmt.Sprint(err)
			ctrl.logger.Error(errInString, zap.String("token", token))
			if errInString == token.CORRUPT_TOKEN || errInString == token.INVALID_TOKEN || errInString == token.EXPIRED_TOKEN {
				rw.WriteHeader(http.StatusUnauthorized)
			} else {
				rw.WriteHeader(http.StatusInternalServerError)
			}
			rw.Write([]byte(errInString))
			return
		}
		// rw.WriteHeader(http.StatusOK)
		// rw.Write([]byte("Authorized Token"))

		// this calls the next function. If not included, the router wont entertain any requests
		next.ServeHTTP(rw, r)
	})
}
