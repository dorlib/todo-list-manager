package middlewares

import (
	"fmt"
	"github.com/spf13/cobra"
	"net/http"
)

func AuthenticationMiddleware(cmd *cobra.Command, args []string) error {
	apiKey, _ := cmd.Flags().GetString("key")
	if apiKey == "" {
		return fmt.Errorf("API key is required")
	}

	// Perform authentication logic here
	// Send API call to authentication service to validate the token

	authURL := "https://auth-service.example.com/validate-token"
	req, err := http.NewRequest(http.MethodGet, authURL, nil)
	if err != nil {
		return fmt.Errorf("failed to create authentication request: %v", err)
	}
	req.Header.Set("Authorization", "Bearer "+apiKey)

	// Send the request to the authentication service
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to validate token: %v", err)
	}

	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("invalid token")
	}

	return nil
}
