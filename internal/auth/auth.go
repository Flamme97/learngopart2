package auth

import (
	"errors"
	"net/http"
	"strings"
)

// get APIKEy extact an API from
// the header of an http request
// example:
// authorization: apiKey {insert apikey here}

func GetAPIKey(headers http.Header)(string, error){
	val := headers.Get("authorization")
	if val == "" {
		return "", errors.New("no authentication info found")
	}

	vals := strings.Split(val, " ")

	if len(vals) != 2 {
		return "", errors.New("malFormed Auth header")
	}
	if vals[0] != "ApiKey" {
		return "", errors.New("malFormed first part Auth header")
	}

	return vals[1], nil
}