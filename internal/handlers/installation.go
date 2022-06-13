package handlers

import (
	"github.com/wfabjanczuk/botProxy/internal/requests"
	"net/http"
)

func (a *app) Install(w http.ResponseWriter, r *http.Request) {
	for _, stepFunction := range a.installationSteps {
		if !stepFunction(w) {
			return
		}
	}

	a.writeSuccess(w, "Installation successful.")
}

func (a *app) createBot(w http.ResponseWriter) bool {
	var err error
	a.botId, err = a.client.CreateBot("Onboarding bot")

	if err != nil {
		return a.writeServerError(w, err, "Installation failed - could not create the bot.")
	}

	return true
}

func (a *app) setRoutingStatus(w http.ResponseWriter) bool {
	err := a.client.SetRoutingStatus(a.botId, "accepting_chats")

	if err != nil {
		return a.writeServerError(w, err, "Installation failed - could not set routing status of the bot.")
	}

	return true
}

func (a *app) registerWebhook(w http.ResponseWriter) bool {
	err := a.client.RegisterWebhook("incoming_event", a.conf.BaseAppUrl+"/webhook", "bot", requests.WebhookFilters{
		AuthorType: "customer",
	})

	if err != nil {
		return a.writeServerError(w, err, "Installation failed - could not register webhook.")
	}

	return true
}

func (a *app) enableLicenseWebhooks(w http.ResponseWriter) bool {
	err := a.client.EnableLicenseWebhooks()

	if err != nil {
		return a.writeServerError(w, err, "Installation failed - could not enable license webhooks.")
	}

	return true
}
