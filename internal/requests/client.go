package requests

import (
	"bytes"
	"fmt"
	"github.com/wfabjanczuk/botProxy/internal/config"
	"io/ioutil"
	"net/http"
)

type Client struct {
	conf         config.Config
	accessToken  string
	refreshToken string
}

func NewClient(conf config.Config) *Client {
	return &Client{conf: conf}
}

func (c *Client) doApiRequest(method, path string, body []byte, extraHeaders map[string][]string, readResponseBody bool) ([]byte, error) {
	request, err := http.NewRequest(method, c.conf.BaseApiUrl+path, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	request.Header = map[string][]string{
		"Authorization": {"Bearer " + c.accessToken},
		"Content-Type":  {"application/json"},
	}
	for key, value := range extraHeaders {
		request.Header[key] = value
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		responseBody, _ := ioutil.ReadAll(response.Body)
		return nil, fmt.Errorf("api responded: %s", responseBody)
	}

	if readResponseBody == false {
		return nil, nil
	}

	return ioutil.ReadAll(response.Body)
}
