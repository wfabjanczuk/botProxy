package requests

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) SetRoutingStatus(agentId, status string) error {
	errPrefix := "setting routing status failed: "

	type payload struct {
		AgentId string `json:"agent_id"`
		Status  string `json:"status"`
	}

	p, err := json.Marshal(&payload{
		AgentId: agentId,
		Status:  status,
	})
	if err != nil {
		return fmt.Errorf("%s%w", errPrefix, err)
	}

	_, err = c.doRequest(http.MethodPost, "/agent/action/set_routing_status", p, nil, false)
	if err != nil {
		return fmt.Errorf("%s%w", errPrefix, err)
	}

	return nil
}
