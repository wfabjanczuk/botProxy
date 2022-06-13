package requests

import (
	"encoding/base64"
	"github.com/wfabjanczuk/botProxy/internal/config"
	"net/http"
	"strings"
)

type Client struct {
	conf config.Config
}

func NewClient(conf config.Config) *Client {
	return &Client{
		conf: conf,
	}
}

func (c *Client) newApiRequest(method, path, body string) (*http.Request, error) {
	r, err := http.NewRequest(method, c.conf.BaseApiUrl+path, strings.NewReader(body))

	if err != nil {
		return nil, err
	}

	basicAuthToken := base64.StdEncoding.EncodeToString([]byte(c.conf.AccountId + ":" + c.conf.PAT))
	r.Header = map[string][]string{
		"Authorization": {"Basic " + basicAuthToken},
		"Content-Type":  {"application/json"},
	}

	return r, err
}
