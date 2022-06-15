package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func (c *Client) Authorize(code string) error {
	errPrefix := "authorizing failed: "

	values := url.Values{}
	values.Set("grant_type", "authorization_code")
	values.Set("code", code)
	values.Set("client_id", c.conf.ClientId)
	values.Set("client_secret", c.conf.SecretKey)
	values.Set("redirect_uri", c.conf.BaseAppUrl+"/install")

	codeExchangeUrl := c.conf.BaseAccountsUrl + "/token?" + values.Encode()
	request, err := http.NewRequest(http.MethodPost, codeExchangeUrl, nil)
	if err != nil {
		return fmt.Errorf("%s%w", errPrefix, err)
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return fmt.Errorf("%s%w", errPrefix, err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("%s%w", errPrefix, err)
	}

	type tokens struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	}

	receivedTokens := &tokens{}
	err = json.Unmarshal(body, receivedTokens)
	if err != nil {
		return fmt.Errorf("%s%w", errPrefix, err)
	}

	c.accessToken = receivedTokens.AccessToken
	c.refreshToken = receivedTokens.RefreshToken

	return nil
}
