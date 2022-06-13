package handlers

import (
	"errors"
	"github.com/wfabjanczuk/botProxy/internal/requests"
	"log"
	"net/http"
	"os"
)

var botId string

func Install(w http.ResponseWriter, r *http.Request) {
	installationSteps := []func(w http.ResponseWriter) bool{
		createBot, setRoutingStatus, registerWebhook, enableLicenseWebhooks,
	}

	for _, stepFunction := range installationSteps {
		if !stepFunction(w) {
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Installation successful."))
}

func createBot(w http.ResponseWriter) bool {
	var err error
	botId, err = requests.CreateBot("Onboarding bot")

	if err != nil {
		log.Println(err)

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Installation failed - could not create the bot."))

		return false
	}

	return true
}

func setRoutingStatus(w http.ResponseWriter) bool {
	err := requests.SetRoutingStatus(botId, "accepting_chats")

	if err != nil {
		log.Println(err)

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Installation failed - could not set routing status of the bot."))

		return false
	}

	return true
}

func registerWebhook(w http.ResponseWriter) bool {
	var err error

	baseAppUrl := os.Getenv("BASE_APP_URL")
	if len(baseAppUrl) == 0 {
		err = errors.New("base app url is not set")
	}

	if err == nil {
		err = requests.RegisterWebhook("incoming_event", baseAppUrl+"/webhook", "bot", requests.WebhookFilters{
			AuthorType: "customer",
		})
	}

	if err != nil {
		log.Println(err)

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Installation failed - could not register webhook."))

		return false
	}

	return true
}

func enableLicenseWebhooks(w http.ResponseWriter) bool {
	var err error

	baseAppUrl := os.Getenv("BASE_APP_URL")
	if len(baseAppUrl) == 0 {
		err = errors.New("base app url is not set")
	}

	if err == nil {
		err = requests.EnableLicenseWebhooks()
	}

	if err != nil {
		log.Println(err)

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Installation failed - could not enable license webhooks."))

		return false
	}

	return true
}
