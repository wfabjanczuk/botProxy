package requests

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) ListWebhooks() ([]string, error) {
	errPrefix := "listing webhooks failed: "

	type payload struct {
		OwnerClientId string `json:"owner_client_id"`
	}

	p, err := json.Marshal(&payload{OwnerClientId: c.conf.ClientId})
	if err != nil {
		return nil, fmt.Errorf("%s%w", errPrefix, err)
	}

	responseBody, err := c.doApiRequest(http.MethodPost, "/configuration/action/list_webhooks", p, nil, true)
	if err != nil {
		return nil, fmt.Errorf("%s%w", errPrefix, err)
	}

	type webhook struct {
		Id string `json:"id"`
	}

	var webhooks []webhook
	err = json.Unmarshal(responseBody, &webhooks)
	if err != nil {
		return nil, fmt.Errorf("%s%w", errPrefix, err)
	}

	var ids []string
	for _, w := range webhooks {
		ids = append(ids, w.Id)
	}

	return ids, nil
}

func (c *Client) UnregisterWebhook(id string) error {
	errPrefix := "unregistering webhook failed: "

	type payload struct {
		Id            string `json:"id"`
		OwnerClientId string `json:"owner_client_id"`
	}

	p, err := json.Marshal(&payload{
		Id:            id,
		OwnerClientId: c.conf.ClientId,
	})
	if err != nil {
		return fmt.Errorf("%s%w", errPrefix, err)
	}

	_, err = c.doApiRequest(http.MethodPost, "/configuration/action/unregister_webhook", p, nil, false)
	if err != nil {
		return fmt.Errorf("%s%w", errPrefix, err)
	}

	return nil
}

func (c *Client) RegisterWebhook(action, url, webhookType, authorType string) error {
	errPrefix := "registering webhook failed: "

	p, err := c.getRegisterWebhookPayload(action, url, webhookType, authorType)
	if err != nil {
		return fmt.Errorf("%s%w", errPrefix, err)
	}

	_, err = c.doApiRequest(http.MethodPost, "/configuration/action/register_webhook", p, nil, false)
	if err != nil {
		return fmt.Errorf("%s%w", errPrefix, err)
	}

	return nil
}

func (c *Client) getRegisterWebhookPayload(action, url, webhookType, authorType string) ([]byte, error) {
	type filters struct {
		AuthorType string `json:"author_type"`
	}

	type payload struct {
		Action        string  `json:"action"`
		Type          string  `json:"type"`
		Url           string  `json:"url"`
		OwnerClientId string  `json:"owner_client_id"`
		SecretKey     string  `json:"secret_key"`
		Filters       filters `json:"filters"`
	}

	return json.Marshal(&payload{
		Action:        action,
		Type:          webhookType,
		Url:           url,
		OwnerClientId: c.conf.ClientId,
		SecretKey:     c.conf.SecretKey,
		Filters: filters{
			AuthorType: authorType,
		},
	})
}

func (c *Client) EnableLicenseWebhooks() error {
	errPrefix := "enabling license webhooks failed: "

	_, err := c.doApiRequest(http.MethodPost, "/configuration/action/enable_license_webhooks", []byte("{}"), nil, false)
	if err != nil {
		return fmt.Errorf("%s%w", errPrefix, err)
	}

	return nil
}
