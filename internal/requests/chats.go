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
		Id                       string `json:"id"`
		IgnoreAgentsAvailability bool   `json:"ignore_agents_availability"`
		IgnoreRequesterPresence  bool   `json:"ignore_requester_presence"`
		Target                   target `json:"target"`
	}

	return json.Marshal(&payload{
		Id:                       chatId,
		IgnoreAgentsAvailability: true,
		IgnoreRequesterPresence:  true,
		Target: target{
			Type: targetType,
			Ids:  targetIds,
		},
	})
}
