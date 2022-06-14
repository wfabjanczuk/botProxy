package requests

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

type tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (c *Client) Authorize(code string) error {
	errPrefix := "authorizing failed: "
	receivedTokens := &tokens{}

	values := url.Values{}
	values.Set("grant_type", "authorization_code")
	values.Set("code", code)
	values.Set("client_id", c.conf.ClientId)
	values.Set("client_secret", c.conf.SecretKey)
	values.Set("redirect_uri", c.conf.BaseAppUrl+"/install")

	codeExchangeUrl := c.conf.BaseAccountsUrl + "/token?" + values.Encode()
	request, err := http.NewRequest("POST", codeExchangeUrl, nil)
	if err != nil {
		return errors.New(errPrefix + err.Error())
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return errors.New(errPrefix + err.Error())
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return errors.New(errPrefix + err.Error())
	}

	err = json.Unmarshal(body, receivedTokens)
	if err != nil {
		return errors.New(errPrefix + err.Error())
	}

	c.accessToken = receivedTokens.AccessToken
	c.refreshToken = receivedTokens.RefreshToken

	return nil
}
