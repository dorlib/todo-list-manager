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
	CORRUPT_TOKEN = "Corrupt Token"
	INVALID_TOKEN = "Invalid Token"
	EXPIRED_TOKEN = "Expired Token"
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

// GenerateToken is used for generating the tokens.
func GenerateToken(header string, payload ClaimsMap, secret string) (string, error) {
	h := hmac.New(sha256.New, []byte(secret))
	header64 := base64.StdEncoding.EncodeToString([]byte(header))

	payloadstr, err := json.Marshal(payload)
	if err != nil {
		return string(payloadstr), fmt.Errorf("Error generating token when encoding payload to string: %w", err)
	}

	payload64 := base64.StdEncoding.EncodeToString(payloadstr)
	message := header64 + "." + payload64
	unsignedStr := header + string(payloadstr)

	h.Write([]byte(unsignedStr))
	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))

	tokenStr := message + "." + signature
	return tokenStr, nil
}

// ValidateToken helps in validating the token
func ValidateToken(token string, secret string) error {
	splitToken := strings.Split(token, ".")
	if len(splitToken) != 3 {
		return errors.New(CORRUPT_TOKEN)
	}

	header, err := base64.StdEncoding.DecodeString(splitToken[0])
	if err != nil {
		return err
	}

	payload, err := base64.StdEncoding.DecodeString(splitToken[1])
	if err != nil {
		return err
	}

	unsignedStr := string(header) + string(payload)
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(unsignedStr))

	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))

	if signature != splitToken[2] {
		return errors.New(INVALID_TOKEN)
	}

	var payloadMap ClaimsMap
	json.Unmarshal(payload, &payloadMap)

	if payloadMap.Exp < fmt.Sprint(time.Now().Unix()) {
		return errors.New(EXPIRED_TOKEN)
	}

	return nil
}
