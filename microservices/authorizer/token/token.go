package token

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

const (
	CorruptToken = "Corrupt Token"
	InvalidToken = "Invalid Token"
	ExpiredToken = "Expired Token"
)

type ClaimsMap struct {
	Aud string
	Iss string
	Exp string
}

// GetSecret fetches the value for the JWT_SECRET from the environment variable
func GetSecret() string {
	return os.Getenv("JWT_SECRET")
}

// GenerateToken is generating the tokens.
func GenerateToken(header string, payload map[string]string, secret string) (string, error) {
	h := hmac.New(sha256.New, []byte(secret))
	header64 := base64.StdEncoding.EncodeToString([]byte(header))

	payloadstr, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error generating Token")
		return string(payloadstr), err
	}

	payload64 := base64.StdEncoding.EncodeToString(payloadstr)
	message := header64 + "." + payload64
	unsignedStr := header + string(payloadstr)

	h.Write([]byte(unsignedStr))
	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))

	tokenStr := message + "." + signature
	return tokenStr, nil
}

// ValidateToken is validating the token
func ValidateToken(token string, secret string) (bool, error) {
	splitToken := strings.Split(token, ".")
	if len(splitToken) != 3 {
		return false, errors.New(CorruptToken)
	}

	header, err := base64.StdEncoding.DecodeString(splitToken[0])
	if err != nil {
		return false, err
	}

	payload, err := base64.StdEncoding.DecodeString(splitToken[1])
	if err != nil {
		return false, err
	}

	unsignedStr := string(header) + string(payload)
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(unsignedStr))

	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))

	if signature != splitToken[2] {
		return false, errors.New(InvalidToken)
	}

	var payloadMap ClaimsMap

	err = json.Unmarshal(payload, &payloadMap)
	if err != nil {
		return false, err
	}

	if payloadMap.Exp < fmt.Sprint(time.Now().Unix()) {
		return false, errors.New(ExpiredToken)
	}

	return true, nil
}
