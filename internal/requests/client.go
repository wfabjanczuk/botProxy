package requests

import (
	"bytes"
	"github.com/wfabjanczuk/botProxy/internal/config"
	"net/http"
)

type Client struct {
	conf         config.Config
	accessToken  string
	refreshToken string
}

func NewClient(conf config.Config) *Client {
	return &Client{
		conf: conf,
	}
}

func (c *Client) newApiRequest(method, path string, body []byte) (*http.Request, error) {
	r, err := http.NewRequest(method, c.conf.BaseApiUrl+path, bytes.NewReader(body))

	if err != nil {
		return nil, err
	}

	r.Header = map[string][]string{
		"Authorization": {"Bearer " + c.accessToken},
		"Content-Type":  {"application/json"},
	}

	return r, err
}
