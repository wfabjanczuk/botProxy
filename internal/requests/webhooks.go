package requests

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
)

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

func RegisterWebhook(action, url, webhookType string, filters WebhookFilters) error {
	secretKey := os.Getenv("SECRET_KEY")
	ownerClientId := os.Getenv("CLIENT_ID")

	if len(secretKey) == 0 {
		return errors.New("client secret key not set")
	}

	if len(ownerClientId) == 0 {
		return errors.New("client id not set")
	}

	payload, err := json.Marshal(&registerWebhookPayload{
		Action:        action,
		SecretKey:     secretKey,
		Url:           url,
		Filters:       filters,
		OwnerClientId: ownerClientId,
		WebhookType:   webhookType,
	})
	if err != nil {
		return err
	}

	request, err := newApiRequest(
		"POST",
		baseApiUrl+"/configuration/action/register_webhook",
		string(payload),
	)

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusOK {
		return errors.New("registering webhook failed")
	}

	return nil
}
