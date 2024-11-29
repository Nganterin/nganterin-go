package helpers

import (
	"fmt"
	"net/url"
)

func ValidateURL(input string) error {
	parsedURL, err := url.ParseRequestURI(input)
	if err != nil {
		return fmt.Errorf("invalid URL: %w", err)
	}

	if parsedURL.Scheme == "" || parsedURL.Host == "" {
		return fmt.Errorf("invalid URL: missing scheme or host")
	}

	return nil
}