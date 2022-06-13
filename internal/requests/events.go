package requests

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type sendEventPayload struct {
	ChatId string `json:"chat_id"`
	Event  Event  `json:"event"`
}

type Event struct {
	Text string `json:"text"`
	Type string `json:"type"`
}

func (c *Client) SendEvent(chatId, authorId, text string) error {
	errPrefix := "sending event failed: "

	payload, err := json.Marshal(&sendEventPayload{
		ChatId: chatId,
		Event: Event{
			Text: text,
			Type: "message",
		},
	})

	request, err := c.newApiRequest(
		"POST",
		"/agent/action/send_event",
		string(payload),
	)
	request.Header["X-Author-Id"] = []string{authorId}
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
