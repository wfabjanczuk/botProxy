package requests

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type createBotPayload struct {
	Name string `json:"name"`
}

type createBotResponse struct {
	Id string `json:"id"`
}

func (c *Client) CreateBot(name string) (botId string, err error) {
	errPrefix := "creating bot failed: "

	payload, err := json.Marshal(&createBotPayload{
		Name: name,
	})
	if err != nil {
		return botId, errors.New(errPrefix + err.Error())
	}

	request, err := c.newApiRequest(
		"POST",
		"/configuration/action/create_bot",
		string(payload),
	)
	if err != nil {
		return botId, errors.New(errPrefix + err.Error())
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return botId, errors.New(errPrefix + err.Error())
	}

	if response.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(response.Body)
		return botId, errors.New(errPrefix + "api responded: " + string(body))
	}

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return botId, errors.New(errPrefix + err.Error())
	}

	result := &createBotResponse{}
	err = json.Unmarshal(responseBody, result)
	if err != nil {
		return botId, errors.New(errPrefix + err.Error())
	}

	return result.Id, nil
}
