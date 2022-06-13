package requests

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type listWebhooksPayload struct {
	OwnerClientId string `json:"owner_client_id"`
}

type webhookListData struct {
	Id string `json:"id"`
}

func (c *Client) ListWebhooks() ([]string, error) {
	var ids []string
	errPrefix := "listing webhooks failed: "

	payload, err := json.Marshal(&listWebhooksPayload{
		OwnerClientId: c.conf.ClientId,
	})
	if err != nil {
		return ids, errors.New(errPrefix + err.Error())
	}

	request, err := c.newApiRequest(
		"POST",
		"/configuration/action/list_webhooks",
		string(payload),
	)

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return ids, errors.New(errPrefix + err.Error())
	}

	body, err := ioutil.ReadAll(response.Body)
	if response.StatusCode != http.StatusOK || err != nil {
		return ids, errors.New(errPrefix + "api responded: " + string(body))
	}

	var webhooks []webhookListData
	err = json.Unmarshal(body, &webhooks)
	if err != nil {
		return ids, errors.New(errPrefix + err.Error())
	}

	for _, webhook := range webhooks {
		ids = append(ids, webhook.Id)
	}

	return ids, nil
}

type unregisterWebhookPayload struct {
	Id            string `json:"id"`
	OwnerClientId string `json:"owner_client_id"`
}

func (c *Client) UnregisterWebhook(id string) error {
	errPrefix := "unregistering webhook failed: "

	payload, err := json.Marshal(&unregisterWebhookPayload{
		Id:            id,
		OwnerClientId: c.conf.ClientId,
	})
	if err != nil {
		return errors.New(errPrefix + err.Error())
	}

	request, err := c.newApiRequest(
		"POST",
		"/configuration/action/unregister_webhook",
		string(payload),
	)

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

type registerWebhookPayload struct {
	Action        string         `json:"action"`
	SecretKey     string         `json:"secret_key"`
	Url           string         `json:"url"`
	Filters       WebhookFilters `json:"filters"`
	OwnerClientId string         `json:"owner_client_id"`
	WebhookType   string         `json:"type"`
}

type WebhookFilters struct {
	AuthorType string `json:"author_type"`
}

func (c *Client) RegisterWebhook(action, url, webhookType string, filters WebhookFilters) error {
	errPrefix := "registering webhook failed: "

	payload, err := json.Marshal(&registerWebhookPayload{
		Action:        action,
		SecretKey:     c.conf.SecretKey,
		Url:           url,
		Filters:       filters,
		OwnerClientId: c.conf.ClientId,
		WebhookType:   webhookType,
	})
	if err != nil {
		return errors.New(errPrefix + err.Error())
	}

	request, err := c.newApiRequest(
		"POST",
		"/configuration/action/register_webhook",
		string(payload),
	)

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

type toggleLicenseWebhooksPayload struct {
	OwnerClientId string `json:"owner_client_id"`
}

func (c *Client) EnableLicenseWebhooks() error {
	errPrefix := "enabling license webhooks failed: "

	payload, err := json.Marshal(&toggleLicenseWebhooksPayload{
		OwnerClientId: c.conf.ClientId,
	})
	if err != nil {
		return errors.New(errPrefix + err.Error())
	}

	request, err := c.newApiRequest(
		"POST",
		"/configuration/action/enable_license_webhooks",
		string(payload),
	)

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
