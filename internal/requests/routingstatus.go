package requests

import (
	"encoding/json"
	"errors"
	"net/http"
)

type setRoutingStatusPayload struct {
	AgentId string `json:"agent_id"`
	Status  string `json:"status"`
}

func SetRoutingStatus(agentId, status string) error {
	payload, err := json.Marshal(&setRoutingStatusPayload{
		AgentId: agentId,
		Status:  status,
	})

	request, err := newApiRequest(
		"POST",
		baseApiUrl+"/agent/action/set_routing_status",
		string(payload),
	)
	if err != nil {
		return err
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusOK {
		return errors.New("setting routing status failed")
	}

	return nil
}
