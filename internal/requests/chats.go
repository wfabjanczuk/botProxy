package requests

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type transferChatPayload struct {
	Id                       string `json:"id"`
	Target                   target `json:"target"`
	IgnoreAgentsAvailability bool   `json:"ignore_agents_availability"`
	IgnoreRequesterPresence  bool   `json:"ignore_requester_presence"`
}

type target struct {
	Type string   `json:"type"`
	Ids  []string `json:"ids"`
}

func (c *Client) TransferChat(chatId, targetType string, targetIds []string) error {
	errPrefix := "transferring chat failed: "

	payload, err := json.Marshal(&transferChatPayload{
		Id: chatId,
		Target: target{
			Type: targetType,
			Ids:  targetIds,
		},
		IgnoreAgentsAvailability: true,
		IgnoreRequesterPresence:  true,
	})

	request, err := c.newApiRequest(
		"POST",
		"/agent/action/transfer_chat",
		string(payload),
	)
	if err != nil {
		return errors.New(errPrefix + err.Error())
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return errors.New(errPrefix + err.Error())
	}

	if response.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(response.Body)
		return errors.New(errPrefix + "api responded: " + string(body))
	}

	return nil
}
