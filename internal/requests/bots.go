package requests

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
)

type createBotPayload struct {
	Name          string `json:"name"`
	OwnerClientId string `json:"owner_client_id"`
}

type createBotResponse struct {
	Id string `json:"id"`
}

func CreateBot(name string) (botId string, err error) {
	ownerClientId := os.Getenv("CLIENT_ID")
	if len(ownerClientId) == 0 {
		return botId, errors.New("client id not set")
	}

	payload, err := json.Marshal(&createBotPayload{
		Name:          name,
		OwnerClientId: ownerClientId,
	})
	if err != nil {
		return botId, err
	}

	request, err := newApiRequest(
		"POST",
		baseApiUrl+"/configuration/action/create_bot",
		string(payload),
	)
	if err != nil {
		return botId, err
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return botId, err
	}

	if response.StatusCode != http.StatusOK {
		return botId, errors.New("creating bot failed")
	}

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return botId, err
	}

	result := &createBotResponse{}
	err = json.Unmarshal(responseBody, result)
	if err != nil {
		return botId, err
	}

	return result.Id, nil
}
