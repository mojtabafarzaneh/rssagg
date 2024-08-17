package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no authentication headers founded")
	}

	values := strings.Split(val, " ")
	if len(values) != 2 {
		return "", errors.New("please enter the authentication in the right form")

	}
	if values[0] != "ApiKey" {
		return "", errors.New("please enter the right header")

	}

	return values[1], nil
}
