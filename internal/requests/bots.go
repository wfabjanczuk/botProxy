package requests

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) CreateBot(name string) (string, error) {
	errPrefix := "creating bot failed: "

	type payload struct {
		Name string `json:"name"`
	}

	p, err := json.Marshal(&payload{Name: name})
	if err != nil {
		return "", fmt.Errorf("%s%w", errPrefix, err)
	}

	responseBody, err := c.doRequest(http.MethodPost, "/configuration/action/create_bot", p, nil, true)
	if err != nil {
		return "", fmt.Errorf("%s%w", errPrefix, err)
	}

	type result struct {
		Id string `json:"id"`
	}

	r := &result{}
	err = json.Unmarshal(responseBody, r)
	if err != nil {
		return "", fmt.Errorf("%s%w", errPrefix, err)
	}

	return r.Id, nil
}
