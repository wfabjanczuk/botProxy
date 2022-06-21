package handlers

import (
	"net/http"
)

func (a *app) Install(w http.ResponseWriter, r *http.Request) {
	for _, stepFunction := range a.installationSteps {
		if !stepFunction(w, r) {
			return
		}
	}

	a.writeSuccess(w, "Installation successful.")
}

func (a *app) authorize(w http.ResponseWriter, r *http.Request) bool {
	code := r.URL.Query().Get("code")
	err := a.client.Authorize(code)

	if err != nil {
		return a.writeClientError(w, err, "Installation failed - invalid authorization code.")
	}

	return true
}

func (a *app) createBot(w http.ResponseWriter, r *http.Request) bool {
	var err error
	a.botId, err = a.client.CreateBot("Onboarding bot")

	if err != nil {
		return a.writeServerError(w, err, "Installation failed - could not create the bot.")
	}

	return true
}

func (a *app) setRoutingStatus(w http.ResponseWriter, r *http.Request) bool {
	err := a.client.SetRoutingStatus(a.botId, "accepting_chats")

	if err != nil {
		return a.writeServerError(w, err, "Installation failed - could not set routing status of the bot.")
	}

	return true
}
