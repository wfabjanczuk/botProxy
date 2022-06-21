package requests

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) TransferChat(chatId, targetType string, targetIds []string) error {
	errPrefix := "transferring chat failed: "

	p, err := c.getTransferChatPayload(chatId, targetType, targetIds)
	if err != nil {
		return fmt.Errorf("%s%w", errPrefix, err)
	}

	_, err = c.doRequest(http.MethodPost, "/agent/action/transfer_chat", p, nil, false)
	if err != nil {
		return fmt.Errorf("%s%w", errPrefix, err)
	}

	return nil
}

func (c *Client) getTransferChatPayload(chatId, targetType string, targetIds []string) ([]byte, error) {
	type target struct {
		Type string   `json:"type"`
		Ids  []string `json:"ids"`
	}
	type payload struct {
		Id                      string `json:"id"`
		Target                  target `json:"target"`
		IgnoreRequesterPresence bool   `json:"ignore_requester_presence"`
	}

	return json.Marshal(&payload{
		Id: chatId,
		Target: target{
			Type: targetType,
			Ids:  targetIds,
		},
		IgnoreRequesterPresence: true,
	})
}
