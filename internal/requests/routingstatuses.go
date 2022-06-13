package requests

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type setRoutingStatusPayload struct {
	AgentId string `json:"agent_id"`
	Status  string `json:"status"`
}

func (c *Client) SetRoutingStatus(agentId, status string) error {
	errPrefix := "setting routing status failed: "

	payload, err := json.Marshal(&setRoutingStatusPayload{
		AgentId: agentId,
		Status:  status,
	})

	request, err := c.newApiRequest(
		"POST",
		"/agent/action/set_routing_status",
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
