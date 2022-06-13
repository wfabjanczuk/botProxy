package requests

import (
	"github.com/wfabjanczuk/botProxy/internal/config"
	"net/http"
	"strings"
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

func (c *Client) newApiRequest(method, path, body string) (*http.Request, error) {
	r, err := http.NewRequest(method, c.conf.BaseApiUrl+path, strings.NewReader(body))

	if err != nil {
		return nil, err
	}

	r.Header = map[string][]string{
		"Authorization": {"Bearer " + c.accessToken},
		"Content-Type":  {"application/json"},
	}

	return r, err
}
