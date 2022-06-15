package requests

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) SendEvent(chatId, authorId, text string) error {
	errPrefix := "sending event failed: "

	p, err := c.getSendEventPayload(chatId, text)
	if err != nil {
		return fmt.Errorf("%s%w", errPrefix, err)
	}

	extraHeaders := map[string][]string{"X-Author-Id": {authorId}}
	_, err = c.doRequest(http.MethodPost, "/agent/action/send_event", p, extraHeaders, false)
	if err != nil {
		return fmt.Errorf("%s%w", errPrefix, err)
	}

	return nil
}

func (c *Client) getSendEventPayload(chatId, text string) ([]byte, error) {
	type button struct {
		Text       string   `json:"text"`
		Type       string   `json:"type"`
		Value      string   `json:"value"`
		PostbackId string   `json:"postback_id"`
		UserIds    []string `json:"user_ids"`
	}

	type element struct {
		Title   string   `json:"title"`
		Buttons []button `json:"buttons"`
	}

	type event struct {
		Type       string    `json:"type"`
		TemplateId string    `json:"template_id"`
		Elements   []element `json:"elements"`
	}

	type payload struct {
		ChatId string `json:"chat_id"`
		Event  event  `json:"event"`
	}

	return json.Marshal(&payload{
		ChatId: chatId,
		Event: event{
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
		},
	})
}
