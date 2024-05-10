package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetAPIKey(headers http.Header) (string, error) {
	key := headers.Get("Authorization")
	if key != "" {
		parts := strings.Split(key, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return "", errors.New("no api key found")
		}
		return parts[1], nil
	}
	return "", errors.New("no api key found")
}
