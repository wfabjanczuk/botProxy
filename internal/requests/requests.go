package requests

import (
	"encoding/base64"
	"errors"
	"net/http"
	"os"
	"strings"
)

const baseApiUrl = "https://api.labs.livechatinc.com/v3.4"

func newApiRequest(method, url, body string) (*http.Request, error) {
	r, err := http.NewRequest(method, url, strings.NewReader(body))

	if err != nil {
		return nil, err
	}

	r.Header, err = getApiHeaders()

	return r, err
}

func getApiHeaders() (map[string][]string, error) {
	accountId := os.Getenv("ACCOUNT_ID")
	pat := os.Getenv("PAT")

	if len(accountId) == 0 || len(pat) == 0 {
		return nil, errors.New("basic auth credentials not set")
	}

	basicAuthToken := base64.StdEncoding.EncodeToString([]byte(accountId + ":" + pat))
	headers := map[string][]string{
		"Authorization": {"Basic " + basicAuthToken},
		"Content-Type":  {"application/json"},
	}

	return headers, nil
}
