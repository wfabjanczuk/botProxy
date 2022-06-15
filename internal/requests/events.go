package requests

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type sendEventPayload struct {
	ChatId string `json:"chat_id"`
	Event  event  `json:"event"`
}

type event struct {
	Type       string    `json:"type"`
	TemplateId string    `json:"template_id"`
	Elements   []element `json:"elements"`
}

type element struct {
	Title   string   `json:"title"`
	Buttons []button `json:"buttons"`
}

type button struct {
	Text       string   `json:"text"`
	Type       string   `json:"type"`
	Value      string   `json:"value"`
	PostbackId string   `json:"postback_id"`
	UserIds    []string `json:"user_ids"`
}

func (c *Client) SendEvent(chatId, authorId, text string) error {
	errPrefix := "sending event failed: "

	payload, err := json.Marshal(&sendEventPayload{
		ChatId: chatId,
		Event:  getRichMessageEvent(text),
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

func getRichMessageEvent(text string) event {
	return event{
		Type:       "rich_message",
		TemplateId: "quick_replies",
		Elements: []element{
			{
				Title: text,
				Buttons: []button{
					{
						Text:       "Yes, I want human",
						Type:       "message",
						Value:      "yes",
						PostbackId: "transfer",
						UserIds:    []string{},
					},
					{
						Text:       "No, bot is fine",
						Type:       "message",
						Value:      "no",
						PostbackId: "transfer",
						UserIds:    []string{},
					},
				},
			},
		},
	}
}
