package requests

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) ListAgentsForTransfer(chatId string) ([]string, error) {
	errPrefix := "listing agents for transfer failed: "

	type payload struct {
		ChatId string `json:"chat_id"`
	}

	p, err := json.Marshal(&payload{ChatId: chatId})
	if err != nil {
		return nil, fmt.Errorf("%s%w", errPrefix, err)
	}

	responseBody, err := c.doRequest(http.MethodPost, "/agent/action/list_agents_for_transfer", p, nil, true)
	if err != nil {
		return nil, fmt.Errorf("%s%w", errPrefix, err)
	}

	type agent struct {
		AgentId          string `json:"agent_id"`
		TotalActiveChats int    `json:"total_active_chats"`
	}

	var agents []agent
	err = json.Unmarshal(responseBody, &agents)
	if err != nil {
		return nil, fmt.Errorf("%s%w", errPrefix, err)
	}

	var ids []string
	for _, w := range agents {
		ids = append(ids, w.AgentId)
	}

	return ids, nil
}
