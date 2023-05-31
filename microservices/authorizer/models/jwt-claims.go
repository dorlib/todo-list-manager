package models

import (
	"fmt"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JwtClaims struct {
	ComapnyID string `json:"comapnyID,omitempty"`
	Username  string `json:"username,omitempty"`
	Roles     string `json:"roles,omitempty"`
	jwt.StandardClaims
}

func (claims JwtClaims) Valid() error {
	var now = time.Now().UTC().Unix()
	flags, _ := GetFlags()
	url, _ := flags.GetApplicationUrl()
	if claims.VerifyExpiresAt(now, true) && claims.VerifyIssuer(*url, true) {
		return nil
	}
	return fmt.Errorf("token is invalid")
}

func (claims JwtClaims) VerifyAudience(origin string) bool {
	return strings.Compare(claims.Audience, origin) == 0
}
