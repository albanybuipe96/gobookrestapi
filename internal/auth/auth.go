package auth

import (
	"errors"
	"net/http"
	"strings"
)

func ExtractAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("forbidden resource")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("malformed auth headers")
	}

	if vals[0] != "ApiKey" {
		return "", errors.New("malformed first part of headers")
	}

	return vals[1], nil
}
