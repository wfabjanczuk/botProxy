package handlers

import (
	"github.com/wfabjanczuk/botProxy/internal/requests"
	"log"
	"net/http"
)

func Install(w http.ResponseWriter, r *http.Request) {
	installationSteps := []func(w http.ResponseWriter) bool{
		createBot, setBotRoutingStatus,
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
	BotId, err = requests.CreateBot("Onboarding bot")

	if err != nil {
		log.Println(err)

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Installation failed - could not create the bot."))

		return false
	}

	return true
}

func setBotRoutingStatus(w http.ResponseWriter) bool {
	err := requests.SetRoutingStatus(BotId, "accepting_chats")

	if err != nil {
		log.Println(err)

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Installation failed - could not set routing status of the bot."))

		return false
	}

	return true
}
